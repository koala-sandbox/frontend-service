/*
 * frontend-service API
 *
 * template for frontend-service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"time"
)

// Item - An item
type Item struct {
	Id int64 `json:"id"`

	Name string `json:"name,omitempty"`

	Date time.Time `json:"date,omitempty"`

	// Some enum property
	EnumProperty string `json:"enum_property,omitempty"`
}

// AssertItemRequired checks if the required fields are not zero-ed
func AssertItemRequired(obj Item) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseItemRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Item (e.g. [][]Item), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseItemRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aItem, ok := obj.(Item)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertItemRequired(aItem)
	})
}
