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
)

func TestCustomerCreate(t *testing.T) {
	request, err := os.ReadFile("../../fixture/create_customer_request.json")
	if err != nil {
		panic(err)
	}
	resCreate, err := http.Post("http://localhost:1323/customer", "application/json", bytes.NewBuffer(request))
	if err != nil {
		panic(err)
	}
	defer resCreate.Body.Close()

	body, _ := io.ReadAll(resCreate.Body)

	var resCreateJson map[string]interface{}

	if err := json.Unmarshal(body, &resCreateJson); err != nil {
		fmt.Println(err)
		return
	}

	resGet, err := http.Get(fmt.Sprintf("http://localhost:1323/customer/%v", resCreateJson["id"]))
	if err != nil {
		panic(err)
	}
	defer resGet.Body.Close()

	resGetJson, err := io.ReadAll(resGet.Body)
	if err != nil {
		panic(err)
	}

	expectedJson, err := os.ReadFile("../../fixture/create_customer_response.customassertion.json")
	if err != nil {
		panic(err)
	}
	util.CompareJsonWithCustomAssertionJson(t, string(expectedJson), string(resGetJson))
}

func TestCustomerGet(t *testing.T) {
	TestCustomerCreate(t)
}
