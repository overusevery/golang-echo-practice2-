// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Customer API
 *
 * API for customer management
 *
 * API version: 1.0.0
 */

package openapi




type ErrorElement struct {

	Id string `json:"id,omitempty"`

	Msg string `json:"msg,omitempty"`
}

// AssertErrorElementRequired checks if the required fields are not zero-ed
func AssertErrorElementRequired(obj ErrorElement) error {
	return nil
}

// AssertErrorElementConstraints checks if the values respects the defined constraints
func AssertErrorElementConstraints(obj ErrorElement) error {
	return nil
}
