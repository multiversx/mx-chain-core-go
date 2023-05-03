package core

import (
	"fmt"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/stretchr/testify/require"
)

func TestDumpGoRoutinesToLogShouldNotPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		r := recover()
		if r != nil {
			require.Fail(t, fmt.Sprintf("should have not paniced %v", r))
		}
	}()

	DumpGoRoutinesToLog(0, &mock.LoggerMock{})
}

func TestGetRunningGoRoutines(t *testing.T) {
	t.Parallel()

	res := GetRunningGoRoutines(&mock.LoggerMock{})
	require.NotNil(t, res)
}
