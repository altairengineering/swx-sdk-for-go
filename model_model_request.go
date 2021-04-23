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

// ModelRequest struct for ModelRequest
type ModelRequest struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewModelRequest instantiates a new ModelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRequest() *ModelRequest {
	this := ModelRequest{}
	return &this
}

// NewModelRequestWithDefaults instantiates a new ModelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRequestWithDefaults() *ModelRequest {
	this := ModelRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelRequest) SetName(v string) {
	o.Name = &v
}

func (o ModelRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableModelRequest struct {
	value *ModelRequest
	isSet bool
}

func (v NullableModelRequest) Get() *ModelRequest {
	return v.value
}

func (v *NullableModelRequest) Set(val *ModelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRequest(val *ModelRequest) *NullableModelRequest {
	return &NullableModelRequest{value: val, isSet: true}
}

func (v NullableModelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


