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

// checks if the MembershipApplicationLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipApplicationLinks{}

// MembershipApplicationLinks struct for MembershipApplicationLinks
type MembershipApplicationLinks struct {
	Self *ResourceLink `json:"self,omitempty"`
}

// NewMembershipApplicationLinks instantiates a new MembershipApplicationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipApplicationLinks() *MembershipApplicationLinks {
	this := MembershipApplicationLinks{}
	return &this
}

// NewMembershipApplicationLinksWithDefaults instantiates a new MembershipApplicationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipApplicationLinksWithDefaults() *MembershipApplicationLinks {
	this := MembershipApplicationLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *MembershipApplicationLinks) GetSelf() ResourceLink {
	if o == nil || IsNil(o.Self) {
		var ret ResourceLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipApplicationLinks) GetSelfOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *MembershipApplicationLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ResourceLink and assigns it to the Self field.
func (o *MembershipApplicationLinks) SetSelf(v ResourceLink) {
	o.Self = &v
}

func (o MembershipApplicationLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipApplicationLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableMembershipApplicationLinks struct {
	value *MembershipApplicationLinks
	isSet bool
}

func (v NullableMembershipApplicationLinks) Get() *MembershipApplicationLinks {
	return v.value
}

func (v *NullableMembershipApplicationLinks) Set(val *MembershipApplicationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipApplicationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipApplicationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipApplicationLinks(val *MembershipApplicationLinks) *NullableMembershipApplicationLinks {
	return &NullableMembershipApplicationLinks{value: val, isSet: true}
}

func (v NullableMembershipApplicationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipApplicationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


