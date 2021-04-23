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

// ModelVersionResponse struct for ModelVersionResponse
type ModelVersionResponse struct {
	Template *ModelVersionRequest `json:"template,omitempty"`
	Version *float32 `json:"version,omitempty"`
}

// NewModelVersionResponse instantiates a new ModelVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelVersionResponse() *ModelVersionResponse {
	this := ModelVersionResponse{}
	return &this
}

// NewModelVersionResponseWithDefaults instantiates a new ModelVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelVersionResponseWithDefaults() *ModelVersionResponse {
	this := ModelVersionResponse{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ModelVersionResponse) GetTemplate() ModelVersionRequest {
	if o == nil || o.Template == nil {
		var ret ModelVersionRequest
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVersionResponse) GetTemplateOk() (*ModelVersionRequest, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ModelVersionResponse) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given ModelVersionRequest and assigns it to the Template field.
func (o *ModelVersionResponse) SetTemplate(v ModelVersionRequest) {
	o.Template = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelVersionResponse) GetVersion() float32 {
	if o == nil || o.Version == nil {
		var ret float32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVersionResponse) GetVersionOk() (*float32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelVersionResponse) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float32 and assigns it to the Version field.
func (o *ModelVersionResponse) SetVersion(v float32) {
	o.Version = &v
}

func (o ModelVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableModelVersionResponse struct {
	value *ModelVersionResponse
	isSet bool
}

func (v NullableModelVersionResponse) Get() *ModelVersionResponse {
	return v.value
}

func (v *NullableModelVersionResponse) Set(val *ModelVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelVersionResponse(val *ModelVersionResponse) *NullableModelVersionResponse {
	return &NullableModelVersionResponse{value: val, isSet: true}
}

func (v NullableModelVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


