package closing

import (
	"errors"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/data/endProcess"
	"github.com/stretchr/testify/assert"
)

func TestNewShuffleOutCloser_InvalidMinWaitShouldErr(t *testing.T) {
	t.Parallel()

	soc, err := NewShuffleOutCloser(
		minDuration-1,
		minDuration,
		make(chan endProcess.ArgEndProcess),
		&mock.LoggerMock{},
	)

	assert.True(t, check.IfNil(soc))
	assert.True(t, errors.Is(err, core.ErrInvalidValue))
}

func TestNewShuffleOutCloser_InvalidMaxWaitShouldErr(t *testing.T) {
	t.Parallel()

	soc, err := NewShuffleOutCloser(
		minDuration,
		minDuration-1,
		make(chan endProcess.ArgEndProcess),
		&mock.LoggerMock{},
	)

	assert.True(t, check.IfNil(soc))
	assert.True(t, errors.Is(err, core.ErrInvalidValue))
}

func TestNewShuffleOutCloser_NilChannelShouldErr(t *testing.T) {
	t.Parallel()

	soc, err := NewShuffleOutCloser(
		minDuration,
		minDuration,
		nil,
		&mock.LoggerMock{},
	)

	assert.True(t, check.IfNil(soc))
	assert.True(t, errors.Is(err, core.ErrNilSignalChan))
}

func TestNewShuffleOutCloser_NilLoggerShouldErr(t *testing.T) {
	t.Parallel()

	soc, err := NewShuffleOutCloser(
		minDuration,
		minDuration,
		make(chan endProcess.ArgEndProcess),
		nil,
	)

	assert.True(t, check.IfNil(soc))
	assert.True(t, errors.Is(err, core.ErrNilLogger))
}

func TestNewShuffleOutCloser_MinWaitDurationLargerThanMaxShouldErr(t *testing.T) {
	t.Parallel()

	soc, err := NewShuffleOutCloser(
		minDuration+1,
		minDuration,
		make(chan endProcess.ArgEndProcess),
		&mock.LoggerMock{},
	)

	assert.True(t, check.IfNil(soc))
	assert.True(t, errors.Is(err, core.ErrInvalidValue))
}

func TestNewShuffleOutCloser_ShouldWork(t *testing.T) {
	t.Parallel()

	soc, err := NewShuffleOutCloser(
		minDuration,
		minDuration,
		make(chan endProcess.ArgEndProcess),
		&mock.LoggerMock{},
	)

	assert.False(t, check.IfNil(soc))
	assert.Nil(t, err)
}

func TestShuffleOutCloser_EndOfProcessingHandlerShouldWork(t *testing.T) {
	t.Parallel()

	ch := make(chan endProcess.ArgEndProcess)
	soc, _ := NewShuffleOutCloser(
		minDuration,
		minDuration,
		ch,
		&mock.LoggerMock{},
	)

	event := endProcess.ArgEndProcess{
		Reason:      "reason",
		Description: "description",
	}
	err := soc.EndOfProcessingHandler(event)
	assert.Nil(t, err)

	time.Sleep(minDuration * 2)

	var recoveredEvent endProcess.ArgEndProcess
	select {
	case recoveredEvent = <-ch:
		assert.Equal(t, event, recoveredEvent)
	default:
		assert.Fail(t, "should have written on channel")
	}
}

func TestShuffleOutCloser_CloseAfterStartShouldWork(t *testing.T) {
	t.Parallel()

	ch := make(chan endProcess.ArgEndProcess)
	soc, _ := NewShuffleOutCloser(
		minDuration,
		minDuration,
		ch,
		&mock.LoggerMock{},
	)

	event := endProcess.ArgEndProcess{
		Reason:      "reason",
		Description: "description",
	}
	_ = soc.EndOfProcessingHandler(event)

	time.Sleep(time.Millisecond * 100)

	err := soc.Close()
	assert.Nil(t, err)

	time.Sleep(minDuration * 2)

	select {
	case <-ch:
		assert.Fail(t, "should have not written on channel")
	default:
	}
}

func TestShuffleOutCloser_CloseBeforeStartShouldWork(t *testing.T) {
	t.Parallel()

	ch := make(chan endProcess.ArgEndProcess)
	soc, _ := NewShuffleOutCloser(
		minDuration,
		minDuration,
		ch,
		&mock.LoggerMock{},
	)

	err := soc.Close()
	assert.Nil(t, err)

	time.Sleep(time.Millisecond * 100)

	event := endProcess.ArgEndProcess{
		Reason:      "reason",
		Description: "description",
	}
	_ = soc.EndOfProcessingHandler(event)

	time.Sleep(minDuration * 2)

	select {
	case <-ch:
		assert.Fail(t, "should have not written on channel")
	default:
	}
}
