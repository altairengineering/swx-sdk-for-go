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

// ModelListResponse struct for ModelListResponse
type ModelListResponse struct {
	Data *[]ModelResponse `json:"data,omitempty"`
	Paging *ActionDelayListResponsePaging `json:"paging,omitempty"`
}

// NewModelListResponse instantiates a new ModelListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelListResponse() *ModelListResponse {
	this := ModelListResponse{}
	return &this
}

// NewModelListResponseWithDefaults instantiates a new ModelListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelListResponseWithDefaults() *ModelListResponse {
	this := ModelListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModelListResponse) GetData() []ModelResponse {
	if o == nil || o.Data == nil {
		var ret []ModelResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelListResponse) GetDataOk() (*[]ModelResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModelListResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelResponse and assigns it to the Data field.
func (o *ModelListResponse) SetData(v []ModelResponse) {
	o.Data = &v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *ModelListResponse) GetPaging() ActionDelayListResponsePaging {
	if o == nil || o.Paging == nil {
		var ret ActionDelayListResponsePaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelListResponse) GetPagingOk() (*ActionDelayListResponsePaging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *ModelListResponse) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ActionDelayListResponsePaging and assigns it to the Paging field.
func (o *ModelListResponse) SetPaging(v ActionDelayListResponsePaging) {
	o.Paging = &v
}

func (o ModelListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableModelListResponse struct {
	value *ModelListResponse
	isSet bool
}

func (v NullableModelListResponse) Get() *ModelListResponse {
	return v.value
}

func (v *NullableModelListResponse) Set(val *ModelListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelListResponse(val *ModelListResponse) *NullableModelListResponse {
	return &NullableModelListResponse{value: val, isSet: true}
}

func (v NullableModelListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


