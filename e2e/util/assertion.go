package util

import (
	"encoding/json"
	"os"

	"github.com/stretchr/testify/assert"
)

// compares two JSON strings with a custom assertion format.
// It replaces values in the actual JSON that match the keyword "<auto-generated>" in the expected JSON
// before performing the comparison. The function returns true if the actual JSON matches the expected JSON.
//
// Parameters:
//
//	t - the testing framework interface
//	expected - the expected JSON string
//	actual - the actual JSON string
//
// Returns:
//
//	bool - true if the actual JSON matches the expected JSON after applying custom assertion rules, false otherwise.
//
// Example:
//
//	expected := `{
//	    "id": "<auto-generated>",
//	    "name": "John Doe"
//	}`
//	actual := `{
//	    "id": 12345,
//	    "name": "John Doe"
//	}`
//	result := compareJsonWithCustomAssertionJson(t, expected, actual) //->true
func CompareJsonWithCustomAssertionJson(t assert.TestingT, expectedJsonPath string, actual string) bool {
	expected, err := os.ReadFile("../../fixture/create_customer_response.customassertion.json")
	if err != nil {
		t.Errorf("something odd!:%v", err)
		return false
	}
	const keyword = "<auto-generated>"
	var expectedJson map[string]interface{}
	if err := json.Unmarshal([]byte(expected), &expectedJson); err != nil {
		panic(err)
	}

	var actualJson map[string]interface{}
	if err := json.Unmarshal([]byte(actual), &actualJson); err != nil {
		panic(err)
	}

	for key, value := range expectedJson {
		if str, ok := value.(string); ok {
			if str == keyword {
				actualJson[key] = keyword
			}
		}
	}
	jsonData, err := json.Marshal(actualJson)
	if err != nil {
		panic(err)
	}

	return assert.JSONEq(t, string(expected), string(jsonData))

}
