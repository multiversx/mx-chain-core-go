package client

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/multiversx/mx-chain-core-go/websocketOutportDriver/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func filterAddress(originalURL string) string {
	if strings.Contains(originalURL, "://") {
		originalURL = strings.Split(originalURL, "://")[1]
	}

	return originalURL
}

func createURLForTestServer(server *httptest.Server) *url.URL {
	return &url.URL{
		Scheme: "ws",
		Host:   filterAddress(server.URL),
		Path:   "/echo",
	}
}

func TestWsConnClient_OpenCloseConnectionShouldWork(t *testing.T) {
	t.Parallel()

	testServer := mock.NewHttpTestEchoHandler()
	defer testServer.Close()

	u := createURLForTestServer(testServer)

	conClient := NewWSConnClient()
	err := conClient.OpenConnection(u.String())
	require.Nil(t, err)

	err = conClient.Close()
	require.Nil(t, err)
}

func TestWsConnClient_WriteAndReadMessageShouldWork(t *testing.T) {
	t.Parallel()

	testServer := mock.NewHttpTestEchoHandler()
	defer testServer.Close()

	u := createURLForTestServer(testServer)

	conClient := NewWSConnClient()
	_ = conClient.OpenConnection(u.String())
	defer func() {
		_ = conClient.Close()
	}()

	message := "TEST"
	err := conClient.WriteMessage(websocket.TextMessage, []byte(message))
	require.Nil(t, err)

	messageType, receivedMessage, err := conClient.ReadMessage()
	require.Nil(t, err)
	assert.Equal(t, websocket.TextMessage, messageType)
	assert.Equal(t, "ECHO: "+message, string(receivedMessage))
}

func TestWsConnClient_WorkingWithANonOpenedConnectionShouldNotPanic(t *testing.T) {
	t.Parallel()

	conClient := NewWSConnClient()
	assert.NotPanics(t, func() {
		err := conClient.Close()
		assert.Equal(t, errConnectionNotOpened, err)
	})
	assert.NotPanics(t, func() {
		err := conClient.WriteMessage(websocket.TextMessage, []byte("TEST"))
		assert.Equal(t, errConnectionNotOpened, err)
	})
	assert.NotPanics(t, func() {
		messageType, message, err := conClient.ReadMessage()
		assert.Equal(t, errConnectionNotOpened, err)
		assert.Equal(t, 0, messageType)
		assert.Nil(t, message)
	})
}

func TestWsConnClient_WorkingWithAClosedConnectionShouldNotPanic(t *testing.T) {
	t.Parallel()

	testServer := mock.NewHttpTestEchoHandler()
	defer testServer.Close()

	u := createURLForTestServer(testServer)

	conClient := NewWSConnClient()
	_ = conClient.OpenConnection(u.String())
	_ = conClient.Close()

	assert.NotPanics(t, func() {
		err := conClient.Close()
		assert.Equal(t, errConnectionNotOpened, err)
	})
	assert.NotPanics(t, func() {
		err := conClient.WriteMessage(websocket.TextMessage, []byte("TEST"))
		assert.Equal(t, errConnectionNotOpened, err)
	})
	assert.NotPanics(t, func() {
		messageType, message, err := conClient.ReadMessage()
		assert.Equal(t, errConnectionNotOpened, err)
		assert.Equal(t, 0, messageType)
		assert.Nil(t, message)
	})
}

func TestWsConnClient_ReOpenConnectionAfterCloseShouldWork(t *testing.T) {
	t.Parallel()

	testServer := mock.NewHttpTestEchoHandler()
	defer testServer.Close()

	u := createURLForTestServer(testServer)

	conClient := NewWSConnClient()
	err := conClient.OpenConnection(u.String())
	require.Nil(t, err)
	err = conClient.Close()
	require.Nil(t, err)

	err = conClient.OpenConnection(u.String())
	require.Nil(t, err)

	message := "TEST"
	err = conClient.WriteMessage(websocket.TextMessage, []byte(message))
	require.Nil(t, err)

	messageType, receivedMessage, err := conClient.ReadMessage()
	require.Nil(t, err)
	assert.Equal(t, websocket.TextMessage, messageType)
	assert.Equal(t, "ECHO: "+message, string(receivedMessage))

	err = conClient.Close()
	require.Nil(t, err)
}

func TestWsConnClient_ReOpenAlreadyOpenedConnectionShouldError(t *testing.T) {
	t.Parallel()

	testServer := mock.NewHttpTestEchoHandler()
	defer testServer.Close()

	u := createURLForTestServer(testServer)

	conClient := NewWSConnClient()
	err := conClient.OpenConnection(u.String())
	require.Nil(t, err)

	err = conClient.OpenConnection(u.String())
	assert.Equal(t, errConnectionAlreadyOpened, err)

	_ = conClient.Close()
}
