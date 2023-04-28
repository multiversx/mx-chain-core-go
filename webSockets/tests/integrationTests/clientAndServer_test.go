package integrationTests

import (
	"strings"
	"sync"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/stretchr/testify/require"
)

func TestStartServerAdd2ClientsAndSendData(t *testing.T) {
	url := "localhost:8833"
	wsServer, err := createServer(url, &mock.LoggerMock{})
	require.Nil(t, err)

	wg := &sync.WaitGroup{}

	wg.Add(3)

	wsServer.RegisterPayloadHandler(&testscommon.PayloadHandlerStub{
		HandlePayloadCalled: func(payload []byte) (*data.PayloadData, error) {
			payloadData := extractPayload(payload)
			require.Equal(t, uint64(1), payloadData.Counter)
			require.True(t, payloadData.WithAcknowledge)

			wg.Done()
			return payloadData, nil
		},
	})

	go func() {
		wsServer.Listen()
		wg.Done()
	}()

	wsClient1, err := createClient(url)
	require.Nil(t, err)

	err = wsClient1.Send(data.WsSendArgs{
		Payload: []byte(""),
	})
	require.Nil(t, err)

	wsClient2, err := createClient(url)
	require.Nil(t, err)

	err = wsClient2.Send(data.WsSendArgs{
		Payload: []byte(""),
	})
	require.Nil(t, err)

	_ = wsClient1.Close()

	_ = wsServer.Close()
	wg.Wait()
}

func TestStartServerAddClientAndCloseClientAndServerShouldReceiveClose(t *testing.T) {
	url := "localhost:8833"

	wg1, wg2 := &sync.WaitGroup{}, &sync.WaitGroup{}
	wg1.Add(2)
	wg2.Add(1)
	serverReceivedCloseMessage := false
	log := &mock.LoggerMock{
		InfoCalled: func(message string, args ...interface{}) {
			if strings.Contains(message, "connection closed") {
				serverReceivedCloseMessage = true
				wg2.Done()
			}
		},
	}

	wsServer, err := createServer(url, log)
	require.Nil(t, err)

	wsServer.RegisterPayloadHandler(&testscommon.PayloadHandlerStub{
		HandlePayloadCalled: func(payload []byte) (*data.PayloadData, error) {
			payloadData := extractPayload(payload)
			require.Equal(t, uint64(1), payloadData.Counter)
			require.True(t, payloadData.WithAcknowledge)

			wg1.Done()
			return payloadData, nil
		},
	})

	go func() {
		wsServer.Listen()
		wg1.Done()
	}()

	wsClient1, err := createClient(url)
	require.Nil(t, err)
	err = wsClient1.Send(data.WsSendArgs{
		Payload: []byte(""),
	})
	err = wsClient1.Close()
	require.Nil(t, err)
	wg2.Wait()
	_ = wsServer.Close()
	wg1.Wait()
	require.True(t, serverReceivedCloseMessage)
}
