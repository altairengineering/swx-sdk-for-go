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

// DeleteClusterResponse struct for DeleteClusterResponse
type DeleteClusterResponse struct {
	ErrorClusterBackend *DeleteClusterResponseErrorClusterBackend `json:"error_cluster_backend,omitempty"`
}

// NewDeleteClusterResponse instantiates a new DeleteClusterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteClusterResponse() *DeleteClusterResponse {
	this := DeleteClusterResponse{}
	return &this
}

// NewDeleteClusterResponseWithDefaults instantiates a new DeleteClusterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteClusterResponseWithDefaults() *DeleteClusterResponse {
	this := DeleteClusterResponse{}
	return &this
}

// GetErrorClusterBackend returns the ErrorClusterBackend field value if set, zero value otherwise.
func (o *DeleteClusterResponse) GetErrorClusterBackend() DeleteClusterResponseErrorClusterBackend {
	if o == nil || o.ErrorClusterBackend == nil {
		var ret DeleteClusterResponseErrorClusterBackend
		return ret
	}
	return *o.ErrorClusterBackend
}

// GetErrorClusterBackendOk returns a tuple with the ErrorClusterBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteClusterResponse) GetErrorClusterBackendOk() (*DeleteClusterResponseErrorClusterBackend, bool) {
	if o == nil || o.ErrorClusterBackend == nil {
		return nil, false
	}
	return o.ErrorClusterBackend, true
}

// HasErrorClusterBackend returns a boolean if a field has been set.
func (o *DeleteClusterResponse) HasErrorClusterBackend() bool {
	if o != nil && o.ErrorClusterBackend != nil {
		return true
	}

	return false
}

// SetErrorClusterBackend gets a reference to the given DeleteClusterResponseErrorClusterBackend and assigns it to the ErrorClusterBackend field.
func (o *DeleteClusterResponse) SetErrorClusterBackend(v DeleteClusterResponseErrorClusterBackend) {
	o.ErrorClusterBackend = &v
}

func (o DeleteClusterResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorClusterBackend != nil {
		toSerialize["error_cluster_backend"] = o.ErrorClusterBackend
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteClusterResponse struct {
	value *DeleteClusterResponse
	isSet bool
}

func (v NullableDeleteClusterResponse) Get() *DeleteClusterResponse {
	return v.value
}

func (v *NullableDeleteClusterResponse) Set(val *DeleteClusterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteClusterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteClusterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteClusterResponse(val *DeleteClusterResponse) *NullableDeleteClusterResponse {
	return &NullableDeleteClusterResponse{value: val, isSet: true}
}

func (v NullableDeleteClusterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteClusterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


