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
