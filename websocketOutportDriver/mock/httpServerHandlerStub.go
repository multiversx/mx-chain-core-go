package mock

import "context"

// HttpServerStub -
type HttpServerStub struct {
	ListenAndServeCalled func() error
	ShutdownCalled       func(ctx context.Context) error
}

// ListenAndServe -
func (h *HttpServerStub) ListenAndServe() error {
	if h.ListenAndServeCalled != nil {
		return h.ListenAndServeCalled()
	}

	return nil
}

//Shutdown -
func (h *HttpServerStub) Shutdown(ctx context.Context) error {
	if h.ShutdownCalled != nil {
		return h.ShutdownCalled(ctx)
	}

	return nil
}
