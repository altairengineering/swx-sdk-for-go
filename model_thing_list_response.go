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

// ThingListResponse struct for ThingListResponse
type ThingListResponse struct {
	Data *[]ThingResponse `json:"data,omitempty"`
	Paging *ActionDelayListResponsePaging `json:"paging,omitempty"`
}

// NewThingListResponse instantiates a new ThingListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThingListResponse() *ThingListResponse {
	this := ThingListResponse{}
	return &this
}

// NewThingListResponseWithDefaults instantiates a new ThingListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThingListResponseWithDefaults() *ThingListResponse {
	this := ThingListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ThingListResponse) GetData() []ThingResponse {
	if o == nil || o.Data == nil {
		var ret []ThingResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingListResponse) GetDataOk() (*[]ThingResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ThingListResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ThingResponse and assigns it to the Data field.
func (o *ThingListResponse) SetData(v []ThingResponse) {
	o.Data = &v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *ThingListResponse) GetPaging() ActionDelayListResponsePaging {
	if o == nil || o.Paging == nil {
		var ret ActionDelayListResponsePaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingListResponse) GetPagingOk() (*ActionDelayListResponsePaging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *ThingListResponse) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ActionDelayListResponsePaging and assigns it to the Paging field.
func (o *ThingListResponse) SetPaging(v ActionDelayListResponsePaging) {
	o.Paging = &v
}

func (o ThingListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableThingListResponse struct {
	value *ThingListResponse
	isSet bool
}

func (v NullableThingListResponse) Get() *ThingListResponse {
	return v.value
}

func (v *NullableThingListResponse) Set(val *ThingListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThingListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThingListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThingListResponse(val *ThingListResponse) *NullableThingListResponse {
	return &NullableThingListResponse{value: val, isSet: true}
}

func (v NullableThingListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThingListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


