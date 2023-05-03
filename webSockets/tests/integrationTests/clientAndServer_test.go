package integrationTests

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSockets"
	"github.com/multiversx/mx-chain-core-go/webSockets/data"
	"github.com/stretchr/testify/require"
)

func TestStartServerAdd2ClientsAndSendData(t *testing.T) {
	url := "localhost:8833"
	wsServer, err := createServer(url, &mock.LoggerMock{})
	require.Nil(t, err)

	payloadConverter, _ := webSockets.NewWebSocketPayloadParser(uint64Converter)

	wg := &sync.WaitGroup{}

	wg.Add(4)

	_ = wsServer.SetPayloadHandler(&testscommon.PayloadHandlerStub{
		ProcessPayloadCalled: func(payload []byte) error {
			require.Equal(t, []byte("test"), payload)
			wg.Done()
			return nil
		},
	})

	go func() {
		wsServer.Start()
		wg.Done()
	}()

	go func() {
		wsServer.Listen()
		wg.Done()
	}()

	wsClient1, err := createClient(url)
	require.Nil(t, err)

	wsClient1.Start()
	time.Sleep(time.Second)

	payload := []byte("test")
	newPayload := payloadConverter.ExtendPayloadWithOperationType(payload, data.OperationSaveBlock)

	err = wsClient1.Send(data.WsSendArgs{
		Payload: newPayload,
	})
	require.Nil(t, err)

	wsClient2, err := createClient(url)
	require.Nil(t, err)

	wsClient2.Start()
	time.Sleep(time.Second)

	err = wsClient2.Send(data.WsSendArgs{
		Payload: newPayload,
	})
	require.Nil(t, err)

	_ = wsClient1.Close()

	_ = wsServer.Close()
	wg.Wait()
}

func TestStartServerAddClientAndCloseClientAndServerShouldReceiveClose(t *testing.T) {
	url := "localhost:8833"

	payloadConverter, _ := webSockets.NewWebSocketPayloadParser(uint64Converter)
	wg1, wg2 := &sync.WaitGroup{}, &sync.WaitGroup{}
	wg1.Add(3)
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

	_ = wsServer.SetPayloadHandler(&testscommon.PayloadHandlerStub{
		ProcessPayloadCalled: func(payload []byte) error {
			require.Equal(t, []byte("text"), payload)
			wg1.Done()
			return nil
		},
	})

	go func() {
		wsServer.Start()
		wg1.Done()
	}()
	go func() {
		wsServer.Listen()
		wg1.Done()
	}()

	wsClient1, err := createClient(url)
	require.Nil(t, err)
	wsClient1.Start()
	time.Sleep(time.Second)

	payload := payloadConverter.ExtendPayloadWithOperationType([]byte("text"), data.OperationSaveBlock)
	err = wsClient1.Send(data.WsSendArgs{
		Payload: payload,
	})
	err = wsClient1.Close()
	require.Nil(t, err)
	wg2.Wait()
	_ = wsServer.Close()
	wg1.Wait()
	require.True(t, serverReceivedCloseMessage)
}
