package core_test

import (
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/core/mock"
	"github.com/stretchr/testify/assert"
)

const identifier = "identifier"

var log = &mock.LoggerMock{}

func TestStopWatch_Start(t *testing.T) {
	t.Parallel()

	sw := core.NewStopWatch()

	sw.Start(identifier)

	_, has := sw.GetStarted(identifier)

	assert.True(t, has)
	assert.Equal(t, identifier, sw.GetIdentifiers()[0])
}

func TestStopWatch_DoubleStartShouldNotReAddInIdentifiers(t *testing.T) {
	t.Parallel()

	sw := core.NewStopWatch()
	identifier1 := "identifier1"
	identifier2 := "identifier2"

	sw.Start(identifier1)
	sw.Start(identifier2)
	sw.Start(identifier1)

	assert.Equal(t, identifier1, sw.GetIdentifiers()[0])
	assert.Equal(t, identifier2, sw.GetIdentifiers()[1])
	assert.Equal(t, 2, len(sw.GetIdentifiers()))
}

func TestStopWatch_StopNoStartShouldNotAddDuration(t *testing.T) {
	t.Parallel()

	sw := core.NewStopWatch()

	sw.Stop(identifier)

	_, has := sw.GetElapsed(identifier)

	assert.False(t, has)
}

func TestStopWatch_StopWithStartShouldAddDuration(t *testing.T) {
	t.Parallel()

	sw := core.NewStopWatch()

	sw.Start(identifier)
	sw.Stop(identifier)

	_, has := sw.GetElapsed(identifier)

	assert.True(t, has)
}

func TestStopWatch_GetMeasurementsNotFinishedShouldOmit(t *testing.T) {
	t.Parallel()

	sw := core.NewStopWatch()

	sw.Start(identifier)

	measurements := sw.GetMeasurements()
	log.Info("measurements", measurements...)

	assert.Equal(t, 0, len(measurements))
}

func TestStopWatch_GetMeasurementsShouldWork(t *testing.T) {
	t.Parallel()

	sw := core.NewStopWatch()

	sw.Start(identifier)
	sw.Stop(identifier)

	measurements := sw.GetMeasurements()
	log.Info("measurements", measurements...)

	assert.Equal(t, 2, len(measurements))
	assert.Equal(t, identifier, measurements[0])
}

func TestStopWatch_AddShouldWork(t *testing.T) {
	t.Parallel()

	identifier1 := "identifier1"
	duration1 := time.Duration(5)
	identifier2 := "identifier2"
	duration2 := time.Duration(7)

	swSrc := core.NewStopWatch()
	swSrc.SetIdentifiers([]string{identifier1, identifier2})
	swSrc.SetElapsed(identifier1, duration1)
	swSrc.SetElapsed(identifier2, duration2)

	sw := core.NewStopWatch()

	sw.Add(swSrc)

	data, _ := sw.GetContainingDuration()
	assert.Equal(t, duration1, data[identifier1])
	assert.Equal(t, duration2, data[identifier2])

	sw.Add(swSrc)

	data, _ = sw.GetContainingDuration()
	assert.Equal(t, duration1*2, data[identifier1])
	assert.Equal(t, duration2*2, data[identifier2])
}

func TestStopWatch_GetMeasurement(t *testing.T) {
	t.Parallel()

	fooDuration := time.Duration(4243) * time.Millisecond
	sw := core.NewStopWatch()
	sw.SetIdentifiers([]string{"foo"})
	sw.SetElapsed("foo", fooDuration)

	assert.Equal(t, fooDuration, sw.GetMeasurement("foo"))
	assert.Equal(t, time.Duration(0), sw.GetMeasurement("bar"))
}
