package mock

// WebsocketConnectionStub -
type WebsocketConnectionStub struct {
	ReadMessageCalled  func() (messageType int, payload []byte, err error)
	WriteMessageCalled func(messageType int, data []byte) error
	CloseCalled        func() error
}

// ReadMessage -
func (w *WebsocketConnectionStub) ReadMessage() (messageType int, payload []byte, err error) {
	if w.ReadMessageCalled != nil {
		return w.ReadMessageCalled()
	}

	return 0, nil, err
}

// WriteMessage -
func (w *WebsocketConnectionStub) WriteMessage(messageType int, data []byte) error {
	if w.WriteMessageCalled != nil {
		return w.WriteMessageCalled(messageType, data)
	}

	return nil
}

// Close -
func (w *WebsocketConnectionStub) Close() error {
	if w.CloseCalled != nil {
		return w.CloseCalled()
	}

	return nil
}
