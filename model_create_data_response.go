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

// CreateDataResponse struct for CreateDataResponse
type CreateDataResponse struct {
	Result *string `json:"Result,omitempty"`
}

// NewCreateDataResponse instantiates a new CreateDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDataResponse() *CreateDataResponse {
	this := CreateDataResponse{}
	return &this
}

// NewCreateDataResponseWithDefaults instantiates a new CreateDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDataResponseWithDefaults() *CreateDataResponse {
	this := CreateDataResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateDataResponse) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataResponse) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateDataResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *CreateDataResponse) SetResult(v string) {
	o.Result = &v
}

func (o CreateDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["Result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDataResponse struct {
	value *CreateDataResponse
	isSet bool
}

func (v NullableCreateDataResponse) Get() *CreateDataResponse {
	return v.value
}

func (v *NullableCreateDataResponse) Set(val *CreateDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDataResponse(val *CreateDataResponse) *NullableCreateDataResponse {
	return &NullableCreateDataResponse{value: val, isSet: true}
}

func (v NullableCreateDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


