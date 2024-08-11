package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/overusevery/golang-echo-practice2/e2e/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomerCreate(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		resCreateJson := post(t, "http://localhost:1323/customer", "../../fixture/e2e/TestCustomerCreate/create_customer_request.json", http.StatusOK)
		resGetJson := get(t, fmt.Sprintf("http://localhost:1323/customer/%v", getFieldInJsonString(t, resCreateJson, "id")), http.StatusOK)
		util.CompareJsonWithCustomAssertionJson(t, "../../fixture/e2e/TestCustomerCreate/create_customer_response.customassertion.json", resGetJson)
	})
}

func TestCustomerUpdate(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		resCreateJson := post(t, "http://localhost:1323/customer", "../../fixture/e2e/TestCustomerUpdate/create_customer_request.json", http.StatusOK)
		_ = put(t, fmt.Sprintf("http://localhost:1323/customer/%v", getFieldInJsonString(t, resCreateJson, "id")), "../../fixture/e2e/TestCustomerUpdate/put_customer_request.json", http.StatusOK)
		resGetJson := get(t, fmt.Sprintf("http://localhost:1323/customer/%v", getFieldInJsonString(t, resCreateJson, "id")), http.StatusOK)
		util.CompareJsonWithCustomAssertionJson(t, "../../fixture/e2e/TestCustomerUpdate/put_customer_response.customassertion.json", resGetJson)
	})
	t.Run("optimistic concurrency control", func(t *testing.T) {
		resCreateJson := post(t, "http://localhost:1323/customer", "../../fixture/e2e/TestCustomerUpdate/create_customer_request.json", http.StatusOK)
		_ = put(t, fmt.Sprintf("http://localhost:1323/customer/%v", getFieldInJsonString(t, resCreateJson, "id")), "../../fixture/e2e/TestCustomerUpdate/put_customer_request_1.json", http.StatusOK)
		_ = put(t, fmt.Sprintf("http://localhost:1323/customer/%v", getFieldInJsonString(t, resCreateJson, "id")), "../../fixture/e2e/TestCustomerUpdate/put_customer_request_2.json", http.StatusConflict)
		resGetJson := get(t, fmt.Sprintf("http://localhost:1323/customer/%v", getFieldInJsonString(t, resCreateJson, "id")), http.StatusOK)
		util.CompareJsonWithCustomAssertionJson(t, "../../fixture/e2e/TestCustomerUpdate/put_customer_response.customassertion.json", resGetJson)
	})
}

func TestGetCustomer(t *testing.T) {
	t.Run("infra return specific error", func(t *testing.T) {
		t.Run("ErrCustomerNotFound", func(t *testing.T) {
			_ = get(t, fmt.Sprintf("http://localhost:1323/customer/%v", "notexistingid"), http.StatusNotFound)
		})
	})
}

func get(t *testing.T, url string, expectedStatus int) string {
	resGet, err := http.Get(url)
	require.NoError(t, err)
	defer resGet.Body.Close()

	assert.Equal(t, expectedStatus, resGet.StatusCode)
	resGetJson, err := io.ReadAll(resGet.Body)
	require.NoError(t, err)
	return string(resGetJson)
}

func post(t *testing.T, url string, jsonPath string, expectedStatus int) string {
	request, err := os.ReadFile(jsonPath)
	require.NoError(t, err)

	resCreate, err := http.Post(url, "application/json", bytes.NewBuffer(request))
	require.NoError(t, err)
	defer resCreate.Body.Close()

	assert.Equal(t, expectedStatus, resCreate.StatusCode)

	body, err := io.ReadAll(resCreate.Body)
	require.NoError(t, err)

	return string(body)
}

func put(t *testing.T, url string, jsonPath string, expectedStatus int) string {
	request, err := os.ReadFile(jsonPath)
	require.NoError(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(request))
	req.Header.Set("Content-Type", "application/json")
	require.NoError(t, err)
	resUpdate, err := client.Do(req)

	require.NoError(t, err)
	defer resUpdate.Body.Close()

	assert.Equal(t, expectedStatus, resUpdate.StatusCode)

	body, err := io.ReadAll(resUpdate.Body)
	require.NoError(t, err)

	return string(body)
}

func getFieldInJsonString(t *testing.T, jsonString string, field string) string {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	require.NoError(t, err)
	return jsonMap[field].(string)
}
