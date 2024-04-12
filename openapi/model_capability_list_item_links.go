/*
SelfService API

SelfService API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CapabilityListItemLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityListItemLinks{}

// CapabilityListItemLinks struct for CapabilityListItemLinks
type CapabilityListItemLinks struct {
	Self *ResourceLink `json:"self,omitempty"`
}

// NewCapabilityListItemLinks instantiates a new CapabilityListItemLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityListItemLinks() *CapabilityListItemLinks {
	this := CapabilityListItemLinks{}
	return &this
}

// NewCapabilityListItemLinksWithDefaults instantiates a new CapabilityListItemLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityListItemLinksWithDefaults() *CapabilityListItemLinks {
	this := CapabilityListItemLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *CapabilityListItemLinks) GetSelf() ResourceLink {
	if o == nil || IsNil(o.Self) {
		var ret ResourceLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityListItemLinks) GetSelfOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *CapabilityListItemLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ResourceLink and assigns it to the Self field.
func (o *CapabilityListItemLinks) SetSelf(v ResourceLink) {
	o.Self = &v
}

func (o CapabilityListItemLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityListItemLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableCapabilityListItemLinks struct {
	value *CapabilityListItemLinks
	isSet bool
}

func (v NullableCapabilityListItemLinks) Get() *CapabilityListItemLinks {
	return v.value
}

func (v *NullableCapabilityListItemLinks) Set(val *CapabilityListItemLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityListItemLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityListItemLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityListItemLinks(val *CapabilityListItemLinks) *NullableCapabilityListItemLinks {
	return &NullableCapabilityListItemLinks{value: val, isSet: true}
}

func (v NullableCapabilityListItemLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityListItemLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


