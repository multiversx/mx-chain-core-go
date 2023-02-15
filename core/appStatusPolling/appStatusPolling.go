package appStatusPolling

import (
	"context"
	"sync"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/check"
)

const minPollingDuration = time.Second

// AppStatusPolling will update an AppStatusHandler by polling components at a predefined interval
type AppStatusPolling struct {
	pollingDuration     time.Duration
	mutRegisteredFunc   sync.RWMutex
	registeredFunctions []func(appStatusHandler core.AppStatusHandler)
	appStatusHandler    core.AppStatusHandler
	log                 core.Logger
}

// NewAppStatusPolling will return an instance of AppStatusPolling
func NewAppStatusPolling(appStatusHandler core.AppStatusHandler, pollingDuration time.Duration, logger core.Logger) (*AppStatusPolling, error) {
	if check.IfNil(appStatusHandler) {
		return nil, ErrNilAppStatusHandler
	}
	if pollingDuration < minPollingDuration {
		return nil, ErrPollingDurationToSmall
	}
	if check.IfNil(logger) {
		return nil, core.ErrNilLogger
	}
	return &AppStatusPolling{
		pollingDuration:  pollingDuration,
		appStatusHandler: appStatusHandler,
		log:              logger,
	}, nil
}

// RegisterPollingFunc will register a new handler function
func (asp *AppStatusPolling) RegisterPollingFunc(handler func(appStatusHandler core.AppStatusHandler)) error {
	if handler == nil {
		return ErrNilHandlerFunc
	}
	asp.mutRegisteredFunc.Lock()
	asp.registeredFunctions = append(asp.registeredFunctions, handler)
	asp.mutRegisteredFunc.Unlock()
	return nil
}

// Poll will notify the AppStatusHandler at a given time
func (asp *AppStatusPolling) Poll(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				asp.log.Debug("closing AppStatusPolling.Poll go routine")
				return
			case <-time.After(asp.pollingDuration):
			}

			asp.mutRegisteredFunc.RLock()
			for _, handler := range asp.registeredFunctions {
				handler(asp.appStatusHandler)
			}
			asp.mutRegisteredFunc.RUnlock()
		}
	}()
}
