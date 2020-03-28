package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestReadConfig(t *testing.T) {
	g := NewGomegaWithT(t)

	ReadConfig()
	g.Expect(HttpServer.RequestTimeout.String()).To(Equal("25s"))
}
