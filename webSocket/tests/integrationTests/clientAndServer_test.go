package integrationTests

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/multiversx/mx-chain-core-go/testscommon"
	"github.com/multiversx/mx-chain-core-go/webSocket/data"
	"github.com/stretchr/testify/require"
)

func TestStartServerAddClientAndSendData(t *testing.T) {
	url := "localhost:8833"
	wsServer, err := createServer(url, &mock.LoggerMock{})
	require.Nil(t, err)

	wg := &sync.WaitGroup{}

	wg.Add(1)

	_ = wsServer.SetPayloadHandler(&testscommon.PayloadHandlerStub{
		ProcessPayloadCalled: func(payloadData *data.PayloadData) error {
			require.Equal(t, []byte("test"), payloadData.Payload)
			wg.Done()
			return nil
		},
	})

	wsServer.Start()

	wsClient, err := createClient(url, &mock.LoggerMock{})
	require.Nil(t, err)

	wsClient.Start()

	for {
		err = wsClient.Send(data.WsSendArgs{
			Payload: []byte("test"),
			OpType:  data.OperationSaveBlock,
		})
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	require.Nil(t, err)

	_ = wsClient.Close()
	_ = wsServer.Close()
	wg.Wait()
}

func TestStartServerAddClientAndCloseClientAndServerShouldReceiveClose(t *testing.T) {
	url := "localhost:8833"

	wg1, wg2 := &sync.WaitGroup{}, &sync.WaitGroup{}
	wg1.Add(1)
	wg2.Add(1)
	serverReceivedCloseMessage := false
	log := &mock.LoggerStub{
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
		ProcessPayloadCalled: func(payloadData *data.PayloadData) error {
			require.Equal(t, []byte("test"), payloadData.Payload)
			wg1.Done()
			return nil
		},
	})

	wsServer.Start()

	wsClient, err := createClient(url, &mock.LoggerMock{})
	require.Nil(t, err)
	wsClient.Start()
	time.Sleep(time.Second)

	for {
		err = wsClient.Send(data.WsSendArgs{
			Payload: []byte("test"),
			OpType:  data.OperationSaveBlock,
		})
		if err == nil {
			break
		}
	}

	err = wsClient.Close()
	require.Nil(t, err)
	wg2.Wait()
	_ = wsServer.Close()
	wg1.Wait()
	require.True(t, serverReceivedCloseMessage)
}

func TestStartServerStartClientCloseServer(t *testing.T) {
	url := "localhost:8833"
	wsServer, err := createServer(url, &mock.LoggerMock{})
	require.Nil(t, err)

	var sentMessages []string
	var receivedMessages []string

	wg := &sync.WaitGroup{}
	wg.Add(1)

	numMessagesReceived := 0
	payloadHandler := &testscommon.PayloadHandlerStub{
		ProcessPayloadCalled: func(payloadData *data.PayloadData) error {
			receivedMessages = append(receivedMessages, string(payloadData.Payload))
			numMessagesReceived++
			if numMessagesReceived == 200 {
				wg.Done()
			}
			return nil
		},
	}
	_ = wsServer.SetPayloadHandler(payloadHandler)

	wsServer.Start()

	wsClient, err := createClient(url, &mock.LoggerMock{})
	require.Nil(t, err)
	wsClient.Start()

	for idx := 0; idx < 100; idx++ {
		message := fmt.Sprintf("%d", idx)
		for {
			err = wsClient.Send(data.WsSendArgs{
				Payload: []byte(message),
				OpType:  data.OperationSaveBlock,
			})
			if err == nil {
				sentMessages = append(sentMessages, message)
				break
			} else {
				time.Sleep(300 * time.Millisecond)
			}
		}
	}

	err = wsServer.Close()
	require.Nil(t, err)

	time.Sleep(5 * time.Second)
	// start the server again
	wsServer, err = createServer(url, &mock.LoggerMock{})
	_ = wsServer.SetPayloadHandler(payloadHandler)
	require.Nil(t, err)
	wsServer.Start()

	for idx := 100; idx < 200; idx++ {
		message := fmt.Sprintf("%d", idx)
		for {
			err = wsClient.Send(data.WsSendArgs{
				Payload: []byte(message),
				OpType:  data.OperationSaveBlock,
			})
			if err == nil {
				sentMessages = append(sentMessages, message)
				break
			} else {
				time.Sleep(300 * time.Millisecond)
			}
		}
	}

	wg.Wait()
	err = wsClient.Close()
	require.Nil(t, err)
	err = wsServer.Close()
	require.Nil(t, err)

	require.Equal(t, 200, numMessagesReceived)
	require.Equal(t, sentMessages, receivedMessages)
}
