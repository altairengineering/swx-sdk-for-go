/*
 * Accounts & Users Service - Public API
 *
 * IN PROGRESS->This is the guide to use the different endpoints to manage the clusters.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AuthZError struct for AuthZError
type AuthZError struct {
	Error *AuthZErrorError `json:"error,omitempty"`
}

// NewAuthZError instantiates a new AuthZError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthZError() *AuthZError {
	this := AuthZError{}
	return &this
}

// NewAuthZErrorWithDefaults instantiates a new AuthZError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthZErrorWithDefaults() *AuthZError {
	this := AuthZError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AuthZError) GetError() AuthZErrorError {
	if o == nil || o.Error == nil {
		var ret AuthZErrorError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthZError) GetErrorOk() (*AuthZErrorError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AuthZError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given AuthZErrorError and assigns it to the Error field.
func (o *AuthZError) SetError(v AuthZErrorError) {
	o.Error = &v
}

func (o AuthZError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableAuthZError struct {
	value *AuthZError
	isSet bool
}

func (v NullableAuthZError) Get() *AuthZError {
	return v.value
}

func (v *NullableAuthZError) Set(val *AuthZError) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthZError) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthZError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthZError(val *AuthZError) *NullableAuthZError {
	return &NullableAuthZError{value: val, isSet: true}
}

func (v NullableAuthZError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthZError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

