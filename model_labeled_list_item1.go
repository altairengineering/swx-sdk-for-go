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

// LabeledListItem1 struct for LabeledListItem1
type LabeledListItem1 struct {
	EntityId *string `json:"entity_id,omitempty"`
	Labels *[]LabelListElement1 `json:"labels,omitempty"`
}

// NewLabeledListItem1 instantiates a new LabeledListItem1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabeledListItem1() *LabeledListItem1 {
	this := LabeledListItem1{}
	return &this
}

// NewLabeledListItem1WithDefaults instantiates a new LabeledListItem1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabeledListItem1WithDefaults() *LabeledListItem1 {
	this := LabeledListItem1{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *LabeledListItem1) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabeledListItem1) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *LabeledListItem1) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *LabeledListItem1) SetEntityId(v string) {
	o.EntityId = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *LabeledListItem1) GetLabels() []LabelListElement1 {
	if o == nil || o.Labels == nil {
		var ret []LabelListElement1
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabeledListItem1) GetLabelsOk() (*[]LabelListElement1, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LabeledListItem1) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelListElement1 and assigns it to the Labels field.
func (o *LabeledListItem1) SetLabels(v []LabelListElement1) {
	o.Labels = &v
}

func (o LabeledListItem1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityId != nil {
		toSerialize["entity_id"] = o.EntityId
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableLabeledListItem1 struct {
	value *LabeledListItem1
	isSet bool
}

func (v NullableLabeledListItem1) Get() *LabeledListItem1 {
	return v.value
}

func (v *NullableLabeledListItem1) Set(val *LabeledListItem1) {
	v.value = val
	v.isSet = true
}

func (v NullableLabeledListItem1) IsSet() bool {
	return v.isSet
}

func (v *NullableLabeledListItem1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabeledListItem1(val *LabeledListItem1) *NullableLabeledListItem1 {
	return &NullableLabeledListItem1{value: val, isSet: true}
}

func (v NullableLabeledListItem1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabeledListItem1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


