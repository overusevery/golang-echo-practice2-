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
	// client level optimisitic conncurrency may be not neccessary feature
	// t.Run("optimistic concurrency control", func(t *testing.T) {
	// 	statusCode, resCreateJson := post(t, url("/customer"), "../../fixture/e2e/TestCustomerUpdate/create_customer_request.json")
	// 	assert.Equal(t, http.StatusOK, statusCode)

	// 	statusCode, _ = put(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")), "../../fixture/e2e/TestCustomerUpdate/put_customer_request_1.json")
	// 	assert.Equal(t, http.StatusOK, statusCode)

	// 	statusCode, _ = put(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")), "../../fixture/e2e/TestCustomerUpdate/put_customer_request_2.json")
	// 	assert.Equal(t, http.StatusConflict, statusCode)

	// 	statusCode, resGetJson := get(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")))
	// 	assert.Equal(t, http.StatusOK, statusCode)
	// 	util.CompareJsonWithCustomAssertionJson(t, "../../fixture/e2e/TestCustomerUpdate/put_customer_response.customassertion.json", resGetJson)
	// })
	t.Run("infra return specific error", func(t *testing.T) {
		t.Run("ErrCustomerNotFound", func(t *testing.T) {
			statusCode, _ := put(t, url("/customer/", "notexists"), "../../fixture/e2e/TestCustomerUpdate/put_customer_request.json")
			assert.Equal(t, http.StatusNotFound, statusCode)
		})
	})
}

func TestCustomerDelete(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		statusCode, resCreateJson := post(t, url("/customer"), "../../fixture/e2e/TestCustomerDelete/create_customer_request.json")
		assert.Equal(t, http.StatusOK, statusCode)

		statusCode, _ = delete(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")))
		assert.Equal(t, http.StatusOK, statusCode)

		statusCode, _ = get(t, url("/customer/", getFieldInJsonString(t, resCreateJson, "id")))
		assert.Equal(t, http.StatusNotFound, statusCode)
	})
	t.Run("infra return specific error", func(t *testing.T) {
		t.Run("ErrCustomerNotFound", func(t *testing.T) {
			statusCode, _ := delete(t, url("/customer/", "notexistingid"))
			assert.Equal(t, http.StatusNotFound, statusCode)
		})
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

// Data
//
//	{
//		"sub": "11111111-1111-1111-1111-111111111111",
//		"iss": "someiss",
//		"client_id": "someclient_id",
//		"scope": "mybackendapi/getcustomer mybackendapi/editcustomer",
//		"exp": 1824767332,
//		"iat": 1724763732,
//		"jti": "22222222-2222-2222-2222-222222222222"
//	}
var authToken = "Bearer eyJraWQiOiJzb21la2lkIiwiYWxnIjoiSFMyNTYifQ.eyJzdWIiOiIxMTExMTExMS0xMTExLTExMTEtMTExMS0xMTExMTExMTExMTEiLCJpc3MiOiJzb21laXNzIiwiY2xpZW50X2lkIjoic29tZWNsaWVudF9pZCIsInNjb3BlIjoibXliYWNrZW5kYXBpL2dldGN1c3RvbWVyIG15YmFja2VuZGFwaS9lZGl0Y3VzdG9tZXIiLCJleHAiOjE4MjQ3NjczMzIsImlhdCI6MTcyNDc2MzczMiwianRpIjoiMjIyMjIyMjItMjIyMi0yMjIyLTIyMjItMjIyMjIyMjIyMjIyIn0.AUPdh5v9fvna4U8NiRKK5aq4AgFzwu1WAMwKC7FSiCY" //nolint:gosec,lll, this is just example dummy token

func get(t *testing.T, url string) (int, string) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)
	resGet, err := client.Do(req)

	require.NoError(t, err)
	defer resGet.Body.Close()

	resGetJson, err := io.ReadAll(resGet.Body)
	require.NoError(t, err)
	return resGet.StatusCode, string(resGetJson)
}

func post(t *testing.T, url string, jsonPath string) (int, string) {
	request, err := os.ReadFile(jsonPath)
	require.NoError(t, err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(request))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)
	resCreate, err := client.Do(req)

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
	req.Header.Set("Authorization", authToken)
	require.NoError(t, err)
	resUpdate, err := client.Do(req)

	require.NoError(t, err)
	defer resUpdate.Body.Close()

	body, err := io.ReadAll(resUpdate.Body)
	require.NoError(t, err)

	return resUpdate.StatusCode, string(body)
}

func delete(t *testing.T, url string) (int, string) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Authorization", authToken)
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
