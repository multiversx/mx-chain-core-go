package factory

import (
	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/core/mock"
	"github.com/ElrondNetwork/elrond-go-core/core/statistics/softwareVersion"
)

type softwareVersionFactory struct {
	statusHandler            core.AppStatusHandler
	stableTagLocation        string
	pollingIntervalInMinutes int
}

// NewSoftwareVersionFactory is responsible for creating a new software version factory object
func NewSoftwareVersionFactory(
	statusHandler core.AppStatusHandler,
	stableTagLocation string,
	pollingIntervalInMinutes int,
) (*softwareVersionFactory, error) {
	if check.IfNil(statusHandler) {
		return nil, core.ErrNilAppStatusHandler
	}

	softwareVersionFactoryObject := &softwareVersionFactory{
		statusHandler:            statusHandler,
		stableTagLocation:        stableTagLocation,
		pollingIntervalInMinutes: pollingIntervalInMinutes,
	}

	return softwareVersionFactoryObject, nil
}

// Create returns an software version checker object
func (svf *softwareVersionFactory) Create() (*softwareVersion.SoftwareVersionChecker, error) {
	stableTagProvider := softwareVersion.NewStableTagProvider(svf.stableTagLocation)
	softwareVersionChecker, err := softwareVersion.NewSoftwareVersionChecker(svf.statusHandler, stableTagProvider, svf.pollingIntervalInMinutes, &mock.LoggerMock{})

	return softwareVersionChecker, err
}
