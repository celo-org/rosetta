// Copyright 2021 Celo Org
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

package rpc

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
)

func TestLogErrDetails(t *testing.T) {
	RegisterTestingT(t)

	errCode := int32(-111)
	errMsg := "Fake Error"
	testErr := errors.New("test error")

	rosettaErr := NewErrorResponse(errCode, errMsg)
	copyErr := NewErrorResponse(errCode, errMsg)
	loggedErr := LogErrDetails(rosettaErr, testErr)

	t.Run("Error details populated", func(t *testing.T) {
		Ω(loggedErr).ShouldNot(Equal(copyErr))
		copyLoggedErr := *loggedErr
		copyLoggedErr.Details = nil
		Ω(*loggedErr).ShouldNot(Equal(copyLoggedErr))
		Ω(copyLoggedErr).Should(Equal(*copyErr))
	})

	t.Run("Logging doesn't change original arg", func(t *testing.T) {
		Ω(rosettaErr).Should(Equal(copyErr))
	})
}
