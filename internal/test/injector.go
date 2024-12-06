package test

import (
	"github.com/maddiesch/go-webservice/internal/injector"
	"github.com/samber/do/v2"
)

func NewInjector(t TestingT) do.Injector {
	t.Helper()

	i := injector.New()

	t.Cleanup(func() {
		err := i.Shutdown()

		if err != nil {
			t.Logf("test injector shutdown failed", err.Error())
		}
	})

	return i
}
