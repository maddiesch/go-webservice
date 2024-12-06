package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestingT interface {
	require.TestingT
	assert.TestingT

	Cleanup(func())
	Name() string
	Helper()
	Log(...any)
	Logf(string, ...any)
}

var _ TestingT = (*testing.T)(nil)
var _ TestingT = (*testing.B)(nil)
