package utils

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestWaitUntil(t *testing.T) {
	RegisterTestingT(t)

	delay := time.After(1 * time.Second)
	Ω(WaitUntil(10*time.Millisecond, 2*time.Second, func() bool {
		select {
		case <-delay:
			return true
		default:
			return false
		}
	})).Should(BeTrue())

	delay = time.After(3 * time.Second)
	Ω(WaitUntil(10*time.Millisecond, 2*time.Second, func() bool {
		select {
		case <-delay:
			return true
		default:
			return false
		}
	})).Should(BeFalse())
}
