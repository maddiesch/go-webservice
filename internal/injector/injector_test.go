package injector_test

import (
	"testing"

	"github.com/maddiesch/go-webservice/internal/injector"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	i := injector.New()
	assert.NotNil(t, i)
	err := i.Shutdown()

	// I have no clue why, but assert.NoError doesn't work here
	// assert.NoError(t, err)
	if err != nil {
		t.Logf("test injector shutdown failed %+v", err.Error())
		t.Fail()
	}
}
