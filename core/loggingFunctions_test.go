package core

import (
	"fmt"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDumpGoRoutinesToLogShouldNotPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		r := recover()
		if r != nil {
			assert.Fail(t, fmt.Sprintf("should have not paniced %v", r))
		}
	}()

	DumpGoRoutinesToLog(0, nil)
}

func TestGetRunningGoRoutines(t *testing.T) {
	t.Parallel()

	res := GetRunningGoRoutines(&mock.LoggerFake{})
	require.NotNil(t, res)
}
