package clientServerSender

import (
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/data/typeConverters/uint64ByteSlice"
	outportData "github.com/multiversx/mx-chain-core-go/websocketOutportDriver/data"
	"github.com/stretchr/testify/require"
)

func TestNewSenderSendAndClose(t *testing.T) {
	t.Parallel()

	args := ArgsWSClientServerSender{
		IsServer:                 true,
		Url:                      "localhost:22111",
		Uint64ByteSliceConverter: uint64ByteSlice.NewBigEndianConverter(),
		RetryDuration:            5 * time.Second,
	}

	clientServerSender, err := NewClientServerSender(args)
	require.Nil(t, err)
	require.NotNil(t, clientServerSender)

	err = clientServerSender.SendMessage([]byte("message"))
	require.Equal(t, outportData.ErrNoClientToSendTo, err)

	err = clientServerSender.Close()
	require.Nil(t, err)
}
