package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/overusevery/golang-echo-practice2/e2e/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const HOST = "http://localhost:1323"

func url(path ...string) string {
	concantenated := HOST
	for _, p := range path {
		concantenated = concantenated + p
	}
	return concantenated
}

func TestCustomerCreate(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		statusCode, resCreateJson := post(t, url("/customer"), "../../fixture/e2e/TestCustomerCreate/create_customer_request.json")
		assert.Equal(t, http.StatusOK, statusCode)

		statusCode, resGetJson := get(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")))
		assert.Equal(t, http.StatusOK, statusCode)
		util.CompareJsonWithCustomAssertionJson(t, "../../fixture/e2e/TestCustomerCreate/create_customer_response.customassertion.json", resGetJson)
	})
}

func TestCustomerUpdate(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		statusCode, resCreateJson := post(t, url("/customer"), "../../fixture/e2e/TestCustomerUpdate/create_customer_request.json")
		assert.Equal(t, http.StatusOK, statusCode)

		statusCode, _ = put(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")), "../../fixture/e2e/TestCustomerUpdate/put_customer_request.json")
		assert.Equal(t, http.StatusOK, statusCode)

		statusCode, resGetJson := get(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")))
		assert.Equal(t, http.StatusOK, statusCode)
		util.CompareJsonWithCustomAssertionJson(t, "../../fixture/e2e/TestCustomerUpdate/put_customer_response.customassertion.json", resGetJson)
	})
	t.Run("optimistic concurrency control", func(t *testing.T) {
		statusCode, resCreateJson := post(t, url("/customer"), "../../fixture/e2e/TestCustomerUpdate/create_customer_request.json")
		assert.Equal(t, http.StatusOK, statusCode)

		statusCode, _ = put(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")), "../../fixture/e2e/TestCustomerUpdate/put_customer_request_1.json")
		assert.Equal(t, http.StatusOK, statusCode)

		statusCode, _ = put(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")), "../../fixture/e2e/TestCustomerUpdate/put_customer_request_2.json")
		assert.Equal(t, http.StatusConflict, statusCode)

		statusCode, resGetJson := get(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")))
		assert.Equal(t, http.StatusOK, statusCode)
		util.CompareJsonWithCustomAssertionJson(t, "../../fixture/e2e/TestCustomerUpdate/put_customer_response.customassertion.json", resGetJson)
	})
}

func TestGetCustomer(t *testing.T) {
	t.Run("infra return specific error", func(t *testing.T) {
		t.Run("ErrCustomerNotFound", func(t *testing.T) {
			statusCode, _ := get(t, url("/customer/", "notexistingid"))
			assert.Equal(t, http.StatusNotFound, statusCode)
		})
	})
}

func get(t *testing.T, url string) (int, string) {
	resGet, err := http.Get(url)
	require.NoError(t, err)
	defer resGet.Body.Close()

	resGetJson, err := io.ReadAll(resGet.Body)
	require.NoError(t, err)
	return resGet.StatusCode, string(resGetJson)
}

func post(t *testing.T, url string, jsonPath string) (int, string) {
	request, err := os.ReadFile(jsonPath)
	require.NoError(t, err)

	resCreate, err := http.Post(url, "application/json", bytes.NewBuffer(request))
	require.NoError(t, err)
	defer resCreate.Body.Close()

	body, err := io.ReadAll(resCreate.Body)
	require.NoError(t, err)

	return resCreate.StatusCode, string(body)
}

func put(t *testing.T, url string, jsonPath string) (int, string) {
	request, err := os.ReadFile(jsonPath)
	require.NoError(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(request))
	req.Header.Set("Content-Type", "application/json")
	require.NoError(t, err)
	resUpdate, err := client.Do(req)

	require.NoError(t, err)
	defer resUpdate.Body.Close()

	body, err := io.ReadAll(resUpdate.Body)
	require.NoError(t, err)

	return resUpdate.StatusCode, string(body)
}

func getFieldInJsonString(t *testing.T, jsonString string, field string) string {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	require.NoError(t, err)
	return jsonMap[field].(string)
}
