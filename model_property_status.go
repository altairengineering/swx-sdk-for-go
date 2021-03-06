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

// PropertyStatus struct for PropertyStatus
type PropertyStatus struct {
	Description *string `json:"description,omitempty"`
	Links *[]ActionRunStatsLinks `json:"links,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewPropertyStatus instantiates a new PropertyStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyStatus() *PropertyStatus {
	this := PropertyStatus{}
	return &this
}

// NewPropertyStatusWithDefaults instantiates a new PropertyStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyStatusWithDefaults() *PropertyStatus {
	this := PropertyStatus{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PropertyStatus) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PropertyStatus) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PropertyStatus) SetDescription(v string) {
	o.Description = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PropertyStatus) GetLinks() []ActionRunStatsLinks {
	if o == nil || o.Links == nil {
		var ret []ActionRunStatsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyStatus) GetLinksOk() (*[]ActionRunStatsLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PropertyStatus) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ActionRunStatsLinks and assigns it to the Links field.
func (o *PropertyStatus) SetLinks(v []ActionRunStatsLinks) {
	o.Links = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PropertyStatus) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyStatus) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PropertyStatus) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PropertyStatus) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PropertyStatus) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyStatus) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PropertyStatus) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PropertyStatus) SetType(v string) {
	o.Type = &v
}

func (o PropertyStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyStatus struct {
	value *PropertyStatus
	isSet bool
}

func (v NullablePropertyStatus) Get() *PropertyStatus {
	return v.value
}

func (v *NullablePropertyStatus) Set(val *PropertyStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyStatus(val *PropertyStatus) *NullablePropertyStatus {
	return &NullablePropertyStatus{value: val, isSet: true}
}

func (v NullablePropertyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


