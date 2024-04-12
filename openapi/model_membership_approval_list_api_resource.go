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

// checks if the MembershipApprovalListApiResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MembershipApprovalListApiResource{}

// MembershipApprovalListApiResource struct for MembershipApprovalListApiResource
type MembershipApprovalListApiResource struct {
	Items []MembershipApprovalApiResource `json:"items,omitempty"`
	Links *MembershipApprovalListLinks `json:"_links,omitempty"`
}

// NewMembershipApprovalListApiResource instantiates a new MembershipApprovalListApiResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMembershipApprovalListApiResource() *MembershipApprovalListApiResource {
	this := MembershipApprovalListApiResource{}
	return &this
}

// NewMembershipApprovalListApiResourceWithDefaults instantiates a new MembershipApprovalListApiResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMembershipApprovalListApiResourceWithDefaults() *MembershipApprovalListApiResource {
	this := MembershipApprovalListApiResource{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MembershipApprovalListApiResource) GetItems() []MembershipApprovalApiResource {
	if o == nil {
		var ret []MembershipApprovalApiResource
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MembershipApprovalListApiResource) GetItemsOk() ([]MembershipApprovalApiResource, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MembershipApprovalListApiResource) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MembershipApprovalApiResource and assigns it to the Items field.
func (o *MembershipApprovalListApiResource) SetItems(v []MembershipApprovalApiResource) {
	o.Items = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MembershipApprovalListApiResource) GetLinks() MembershipApprovalListLinks {
	if o == nil || IsNil(o.Links) {
		var ret MembershipApprovalListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MembershipApprovalListApiResource) GetLinksOk() (*MembershipApprovalListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MembershipApprovalListApiResource) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given MembershipApprovalListLinks and assigns it to the Links field.
func (o *MembershipApprovalListApiResource) SetLinks(v MembershipApprovalListLinks) {
	o.Links = &v
}

func (o MembershipApprovalListApiResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MembershipApprovalListApiResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableMembershipApprovalListApiResource struct {
	value *MembershipApprovalListApiResource
	isSet bool
}

func (v NullableMembershipApprovalListApiResource) Get() *MembershipApprovalListApiResource {
	return v.value
}

func (v *NullableMembershipApprovalListApiResource) Set(val *MembershipApprovalListApiResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMembershipApprovalListApiResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMembershipApprovalListApiResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMembershipApprovalListApiResource(val *MembershipApprovalListApiResource) *NullableMembershipApprovalListApiResource {
	return &NullableMembershipApprovalListApiResource{value: val, isSet: true}
}

func (v NullableMembershipApprovalListApiResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMembershipApprovalListApiResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


