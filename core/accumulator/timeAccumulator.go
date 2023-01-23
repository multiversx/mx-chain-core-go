package accumulator

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
)

var _ core.Accumulator = (*timeAccumulator)(nil)

const minimumAllowedTime = time.Millisecond * 10

// timeAccumulator is a structure that is able to accumulate data and will try to write on the output channel
//once per provided interval
type timeAccumulator struct {
	cancel         func()
	maxAllowedTime time.Duration
	maxOffset      time.Duration
	mut            sync.Mutex
	data           []interface{}
	output         chan []interface{}
	log            core.Logger
}

// NewTimeAccumulator returns a new accumulator instance
func NewTimeAccumulator(maxAllowedTime time.Duration, maxOffset time.Duration, logger core.Logger) (*timeAccumulator, error) {
	if maxAllowedTime < minimumAllowedTime {
		return nil, fmt.Errorf("%w for maxAllowedTime as minimum allowed time is %v",
			core.ErrInvalidValue,
			minimumAllowedTime,
		)
	}
	if maxOffset < 0 {
		return nil, fmt.Errorf("%w for maxOffset: should not be negative", core.ErrInvalidValue)
	}
	if check.IfNil(logger) {
		return nil, core.ErrNilLogger
	}

	ctx, cancel := context.WithCancel(context.Background())

	ta := &timeAccumulator{
		cancel:         cancel,
		maxAllowedTime: maxAllowedTime,
		output:         make(chan []interface{}),
		maxOffset:      maxOffset,
		log:            logger,
	}

	go ta.continuousEviction(ctx)

	return ta, nil
}

// AddData will append a new data on the queue
func (ta *timeAccumulator) AddData(data interface{}) {
	ta.mut.Lock()
	ta.data = append(ta.data, data)
	ta.mut.Unlock()
}

// OutputChannel returns the output channel on which accumulated data will be sent periodically
func (ta *timeAccumulator) OutputChannel() <-chan []interface{} {
	return ta.output
}

// will call do eviction periodically until the context is done
func (ta *timeAccumulator) continuousEviction(ctx context.Context) {
	defer func() {
		close(ta.output)
	}()

	for {
		select {
		case <-time.After(ta.computeWaitTime()):
			isDone := ta.doEviction(ctx)
			if isDone {
				return
			}
		case <-ctx.Done():
			ta.log.Debug("closing timeAccumulator.continuousEviction go routine")
			return
		}
	}
}

func (ta *timeAccumulator) computeWaitTime() time.Duration {
	if ta.maxOffset == 0 {
		return ta.maxAllowedTime
	}

	randBuff := make([]byte, 4)
	_, _ = rand.Read(randBuff)
	randUint64 := binary.BigEndian.Uint32(randBuff)
	offset := time.Duration(randUint64) % ta.maxOffset

	return ta.maxAllowedTime - offset
}

// doEviction will do the eviction of all accumulated data
// if context.Done is triggered during the eviction, the whole operation will be aborted
func (ta *timeAccumulator) doEviction(ctx context.Context) bool {
	ta.mut.Lock()
	tempData := make([]interface{}, len(ta.data))
	copy(tempData, ta.data)
	ta.data = nil
	ta.mut.Unlock()

	select {
	case ta.output <- tempData:
		return false
	case <-ctx.Done():
		ta.log.Debug("closing timeAccumulator.doEviction go routine")
		return true
	}
}

// Close stops the time accumulator's eviction loop and closes the output chan
func (ta *timeAccumulator) Close() error {
	ta.cancel()
	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (ta *timeAccumulator) IsInterfaceNil() bool {
	return ta == nil
}
