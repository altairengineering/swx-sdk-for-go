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

// LabelListElement1 struct for LabelListElement1
type LabelListElement1 struct {
	Color *string `json:"color,omitempty"`
	Id *string `json:"id,omitempty"`
	LabelDescription *string `json:"label_description,omitempty"`
	LabelName *string `json:"label_name,omitempty"`
	Mqtt *bool `json:"mqtt,omitempty"`
	Space *string `json:"space,omitempty"`
}

// NewLabelListElement1 instantiates a new LabelListElement1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelListElement1() *LabelListElement1 {
	this := LabelListElement1{}
	return &this
}

// NewLabelListElement1WithDefaults instantiates a new LabelListElement1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelListElement1WithDefaults() *LabelListElement1 {
	this := LabelListElement1{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *LabelListElement1) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelListElement1) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *LabelListElement1) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *LabelListElement1) SetColor(v string) {
	o.Color = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LabelListElement1) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelListElement1) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LabelListElement1) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LabelListElement1) SetId(v string) {
	o.Id = &v
}

// GetLabelDescription returns the LabelDescription field value if set, zero value otherwise.
func (o *LabelListElement1) GetLabelDescription() string {
	if o == nil || o.LabelDescription == nil {
		var ret string
		return ret
	}
	return *o.LabelDescription
}

// GetLabelDescriptionOk returns a tuple with the LabelDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelListElement1) GetLabelDescriptionOk() (*string, bool) {
	if o == nil || o.LabelDescription == nil {
		return nil, false
	}
	return o.LabelDescription, true
}

// HasLabelDescription returns a boolean if a field has been set.
func (o *LabelListElement1) HasLabelDescription() bool {
	if o != nil && o.LabelDescription != nil {
		return true
	}

	return false
}

// SetLabelDescription gets a reference to the given string and assigns it to the LabelDescription field.
func (o *LabelListElement1) SetLabelDescription(v string) {
	o.LabelDescription = &v
}

// GetLabelName returns the LabelName field value if set, zero value otherwise.
func (o *LabelListElement1) GetLabelName() string {
	if o == nil || o.LabelName == nil {
		var ret string
		return ret
	}
	return *o.LabelName
}

// GetLabelNameOk returns a tuple with the LabelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelListElement1) GetLabelNameOk() (*string, bool) {
	if o == nil || o.LabelName == nil {
		return nil, false
	}
	return o.LabelName, true
}

// HasLabelName returns a boolean if a field has been set.
func (o *LabelListElement1) HasLabelName() bool {
	if o != nil && o.LabelName != nil {
		return true
	}

	return false
}

// SetLabelName gets a reference to the given string and assigns it to the LabelName field.
func (o *LabelListElement1) SetLabelName(v string) {
	o.LabelName = &v
}

// GetMqtt returns the Mqtt field value if set, zero value otherwise.
func (o *LabelListElement1) GetMqtt() bool {
	if o == nil || o.Mqtt == nil {
		var ret bool
		return ret
	}
	return *o.Mqtt
}

// GetMqttOk returns a tuple with the Mqtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelListElement1) GetMqttOk() (*bool, bool) {
	if o == nil || o.Mqtt == nil {
		return nil, false
	}
	return o.Mqtt, true
}

// HasMqtt returns a boolean if a field has been set.
func (o *LabelListElement1) HasMqtt() bool {
	if o != nil && o.Mqtt != nil {
		return true
	}

	return false
}

// SetMqtt gets a reference to the given bool and assigns it to the Mqtt field.
func (o *LabelListElement1) SetMqtt(v bool) {
	o.Mqtt = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *LabelListElement1) GetSpace() string {
	if o == nil || o.Space == nil {
		var ret string
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelListElement1) GetSpaceOk() (*string, bool) {
	if o == nil || o.Space == nil {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *LabelListElement1) HasSpace() bool {
	if o != nil && o.Space != nil {
		return true
	}

	return false
}

// SetSpace gets a reference to the given string and assigns it to the Space field.
func (o *LabelListElement1) SetSpace(v string) {
	o.Space = &v
}

func (o LabelListElement1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LabelDescription != nil {
		toSerialize["label_description"] = o.LabelDescription
	}
	if o.LabelName != nil {
		toSerialize["label_name"] = o.LabelName
	}
	if o.Mqtt != nil {
		toSerialize["mqtt"] = o.Mqtt
	}
	if o.Space != nil {
		toSerialize["space"] = o.Space
	}
	return json.Marshal(toSerialize)
}

type NullableLabelListElement1 struct {
	value *LabelListElement1
	isSet bool
}

func (v NullableLabelListElement1) Get() *LabelListElement1 {
	return v.value
}

func (v *NullableLabelListElement1) Set(val *LabelListElement1) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelListElement1) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelListElement1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelListElement1(val *LabelListElement1) *NullableLabelListElement1 {
	return &NullableLabelListElement1{value: val, isSet: true}
}

func (v NullableLabelListElement1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelListElement1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


