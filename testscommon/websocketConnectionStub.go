package testscommon

// WebsocketConnectionStub -
type WebsocketConnectionStub struct {
	OpenConnectionCalled  func(url string) error
	ReadMessageCalled     func() (messageType int, payload []byte, err error)
	WriteMessageCalled    func(messageType int, data []byte) error
	SetCloseHandlerCalled func(closeHandler func(code int, text string) error)
	GetIDCalled           func() string
	CloseCalled           func() error
}

// OpenConnection -
func (w *WebsocketConnectionStub) OpenConnection(url string) error {
	if w.OpenConnectionCalled != nil {
		return w.OpenConnectionCalled(url)
	}

	return nil
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

// GetID -
func (w *WebsocketConnectionStub) GetID() string {
	if w.GetIDCalled != nil {
		return w.GetIDCalled()
	}
	return ""
}

// Close -
func (w *WebsocketConnectionStub) Close() error {
	if w.CloseCalled != nil {
		return w.CloseCalled()
	}

	return nil
}

// SetCloseHandler -
func (w *WebsocketConnectionStub) SetCloseHandler(closeHandler func(code int, text string) error) {
	if w.SetCloseHandlerCalled != nil {
		w.SetCloseHandlerCalled(closeHandler)
	}
	return
}

// IsInterfaceNil -
func (w *WebsocketConnectionStub) IsInterfaceNil() bool {
	return w == nil
}
