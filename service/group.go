package service

import (
	"context"
)

type serviceGroup struct {
	services []Service
	running  RunningLock
}

func Group(services ...Service) *serviceGroup {
	return &serviceGroup{
		services: services,
	}

}

func (sg *serviceGroup) Name() string {
	return "servicegroup"
}

func (sg *serviceGroup) Running() bool {
	return sg.running.Running()
}

func (sg *serviceGroup) Start(ctx context.Context) error {
	if err := sg.running.EnableOrFail(); err != nil {
		return err
	}
	defer sg.running.Disable()

	return RunServices(ctx, sg.services...)
}
