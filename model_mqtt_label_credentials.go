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

// MQTTLabelCredentials struct for MQTTLabelCredentials
type MQTTLabelCredentials struct {
	Enabled *bool `json:"enabled,omitempty"`
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewMQTTLabelCredentials instantiates a new MQTTLabelCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMQTTLabelCredentials() *MQTTLabelCredentials {
	this := MQTTLabelCredentials{}
	return &this
}

// NewMQTTLabelCredentialsWithDefaults instantiates a new MQTTLabelCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMQTTLabelCredentialsWithDefaults() *MQTTLabelCredentials {
	this := MQTTLabelCredentials{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MQTTLabelCredentials) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MQTTLabelCredentials) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MQTTLabelCredentials) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MQTTLabelCredentials) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *MQTTLabelCredentials) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MQTTLabelCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *MQTTLabelCredentials) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *MQTTLabelCredentials) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *MQTTLabelCredentials) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MQTTLabelCredentials) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *MQTTLabelCredentials) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *MQTTLabelCredentials) SetUsername(v string) {
	o.Username = &v
}

func (o MQTTLabelCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableMQTTLabelCredentials struct {
	value *MQTTLabelCredentials
	isSet bool
}

func (v NullableMQTTLabelCredentials) Get() *MQTTLabelCredentials {
	return v.value
}

func (v *NullableMQTTLabelCredentials) Set(val *MQTTLabelCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableMQTTLabelCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableMQTTLabelCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMQTTLabelCredentials(val *MQTTLabelCredentials) *NullableMQTTLabelCredentials {
	return &NullableMQTTLabelCredentials{value: val, isSet: true}
}

func (v NullableMQTTLabelCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMQTTLabelCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


