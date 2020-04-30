// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
