package test

import (
	"testing"

	"github.com/maddiesch/go-webservice/internal/injector"
	"github.com/samber/do/v2"
)

func NewInjector(t testing.TB) do.Injector {
	t.Helper()

	i := injector.New()

	t.Cleanup(func() {
		t.Helper()

		err := i.Shutdown()

		if err != nil {
			t.Logf("test injector shutdown failed %+v", err.Error())
		}
	})

	return i
}
