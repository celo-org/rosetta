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

package airgap

import (
	"encoding/json"
)

type filterQueryRawData struct {
	Event     *string         `json:"event,omitempty"`
	FromBlock *string         `json:"from_block,omitempty"`
	ToBlock   *string         `json:"to_block,omitempty"`
	Topics    [][]interface{} `json:"topics,omitempty"`
}

func (data *filterQueryRawData) transform(args *FilterQueryParams) error {
	var err error
	args.FromBlock, err = stringToBigInt(data.FromBlock)
	if err != nil {
		return err
	}
	args.ToBlock, err = stringToBigInt(data.ToBlock)
	if err != nil {
		return err
	}
	args.Event, err = stringToEvent(data.Event)
	if err != nil {
		return err
	}

	args.Topics = data.Topics
	return nil
}

func (args *FilterQueryParams) transform() *filterQueryRawData {
	var data filterQueryRawData
	data.FromBlock = bigIntToString(args.FromBlock)
	data.ToBlock = bigIntToString(args.ToBlock)
	if args.Event != nil {
		var str = args.Event.String()
		data.Event = &str
	}
	data.Topics = args.Topics
	return &data
}

func (args FilterQueryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(args.transform())
}

func (args *FilterQueryParams) UnmarshalJSON(b []byte) error {
	var err error

	var data filterQueryRawData
	if err = json.Unmarshal(b, &data); err != nil {
		return err
	}
	return data.transform(args)
}
