package closing

import (
	"sync"
)

type safeChanCloser struct {
	closeMut sync.RWMutex
	chClose  chan struct{}
}

// NewSafeChanCloser returns a safe chan closer instance
func NewSafeChanCloser() *safeChanCloser {
	return &safeChanCloser{
		chClose: make(chan struct{}),
	}
}

// Close will close the channel in a safe concurrent manner
func (closer *safeChanCloser) Close() {
	closer.closeMut.Lock()
	defer closer.closeMut.Unlock()

	closer.closeUnprotected()
}

func (closer *safeChanCloser) closeUnprotected() {
	select {
	case <-closer.chClose:
		return
	default:
		close(closer.chClose)
	}
}

// ChanClose returns the closing channel
func (closer *safeChanCloser) ChanClose() <-chan struct{} {
	closer.closeMut.RLock()
	defer closer.closeMut.RUnlock()

	return closer.chClose
}

// Reset will close and then open again the channel
func (closer *safeChanCloser) Reset() {
	closer.closeMut.Lock()
	defer closer.closeMut.Unlock()

	closer.closeUnprotected()
	closer.chClose = make(chan struct{})
}

// IsInterfaceNil returns true if there is no value under the interface
func (closer *safeChanCloser) IsInterfaceNil() bool {
	return closer == nil
}
