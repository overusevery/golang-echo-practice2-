package testutil

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertResBodyIsEquWithJson(t *testing.T, res *httptest.ResponseRecorder, expectedJsonPath string) {
	expectedJson, err := os.ReadFile(expectedJsonPath)
	if err != nil {
		panic(err)
	}

	assert.JSONEq(t, string(expectedJson), res.Body.String())
}
