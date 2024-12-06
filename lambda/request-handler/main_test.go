package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maddiesch/go-webservice/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandler(t *testing.T) {
	i := test.NewInjector(t)
	h := createRequestHandler(i)

	t.Run("GET /", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		h.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Hello, World!\n", res.Body.String())
	})
}
