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

// checks if the CapabilityDetailsApiResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityDetailsApiResource{}

// CapabilityDetailsApiResource struct for CapabilityDetailsApiResource
type CapabilityDetailsApiResource struct {
	Id NullableString `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Status NullableString `json:"status,omitempty"`
	Description NullableString `json:"description,omitempty"`
	JsonMetadata NullableString `json:"jsonMetadata,omitempty"`
	JsonMetadataSchemaVersion *int32 `json:"jsonMetadataSchemaVersion,omitempty"`
	Links *CapabilityDetailsLinks `json:"_links,omitempty"`
}

// NewCapabilityDetailsApiResource instantiates a new CapabilityDetailsApiResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityDetailsApiResource() *CapabilityDetailsApiResource {
	this := CapabilityDetailsApiResource{}
	return &this
}

// NewCapabilityDetailsApiResourceWithDefaults instantiates a new CapabilityDetailsApiResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityDetailsApiResourceWithDefaults() *CapabilityDetailsApiResource {
	this := CapabilityDetailsApiResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityDetailsApiResource) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityDetailsApiResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CapabilityDetailsApiResource) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CapabilityDetailsApiResource) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CapabilityDetailsApiResource) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CapabilityDetailsApiResource) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityDetailsApiResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityDetailsApiResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CapabilityDetailsApiResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CapabilityDetailsApiResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CapabilityDetailsApiResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CapabilityDetailsApiResource) UnsetName() {
	o.Name.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityDetailsApiResource) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityDetailsApiResource) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *CapabilityDetailsApiResource) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *CapabilityDetailsApiResource) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *CapabilityDetailsApiResource) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *CapabilityDetailsApiResource) UnsetStatus() {
	o.Status.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityDetailsApiResource) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityDetailsApiResource) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilityDetailsApiResource) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CapabilityDetailsApiResource) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CapabilityDetailsApiResource) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CapabilityDetailsApiResource) UnsetDescription() {
	o.Description.Unset()
}

// GetJsonMetadata returns the JsonMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityDetailsApiResource) GetJsonMetadata() string {
	if o == nil || IsNil(o.JsonMetadata.Get()) {
		var ret string
		return ret
	}
	return *o.JsonMetadata.Get()
}

// GetJsonMetadataOk returns a tuple with the JsonMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityDetailsApiResource) GetJsonMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JsonMetadata.Get(), o.JsonMetadata.IsSet()
}

// HasJsonMetadata returns a boolean if a field has been set.
func (o *CapabilityDetailsApiResource) HasJsonMetadata() bool {
	if o != nil && o.JsonMetadata.IsSet() {
		return true
	}

	return false
}

// SetJsonMetadata gets a reference to the given NullableString and assigns it to the JsonMetadata field.
func (o *CapabilityDetailsApiResource) SetJsonMetadata(v string) {
	o.JsonMetadata.Set(&v)
}
// SetJsonMetadataNil sets the value for JsonMetadata to be an explicit nil
func (o *CapabilityDetailsApiResource) SetJsonMetadataNil() {
	o.JsonMetadata.Set(nil)
}

// UnsetJsonMetadata ensures that no value is present for JsonMetadata, not even an explicit nil
func (o *CapabilityDetailsApiResource) UnsetJsonMetadata() {
	o.JsonMetadata.Unset()
}

// GetJsonMetadataSchemaVersion returns the JsonMetadataSchemaVersion field value if set, zero value otherwise.
func (o *CapabilityDetailsApiResource) GetJsonMetadataSchemaVersion() int32 {
	if o == nil || IsNil(o.JsonMetadataSchemaVersion) {
		var ret int32
		return ret
	}
	return *o.JsonMetadataSchemaVersion
}

// GetJsonMetadataSchemaVersionOk returns a tuple with the JsonMetadataSchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDetailsApiResource) GetJsonMetadataSchemaVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.JsonMetadataSchemaVersion) {
		return nil, false
	}
	return o.JsonMetadataSchemaVersion, true
}

// HasJsonMetadataSchemaVersion returns a boolean if a field has been set.
func (o *CapabilityDetailsApiResource) HasJsonMetadataSchemaVersion() bool {
	if o != nil && !IsNil(o.JsonMetadataSchemaVersion) {
		return true
	}

	return false
}

// SetJsonMetadataSchemaVersion gets a reference to the given int32 and assigns it to the JsonMetadataSchemaVersion field.
func (o *CapabilityDetailsApiResource) SetJsonMetadataSchemaVersion(v int32) {
	o.JsonMetadataSchemaVersion = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CapabilityDetailsApiResource) GetLinks() CapabilityDetailsLinks {
	if o == nil || IsNil(o.Links) {
		var ret CapabilityDetailsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityDetailsApiResource) GetLinksOk() (*CapabilityDetailsLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CapabilityDetailsApiResource) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CapabilityDetailsLinks and assigns it to the Links field.
func (o *CapabilityDetailsApiResource) SetLinks(v CapabilityDetailsLinks) {
	o.Links = &v
}

func (o CapabilityDetailsApiResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityDetailsApiResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.JsonMetadata.IsSet() {
		toSerialize["jsonMetadata"] = o.JsonMetadata.Get()
	}
	if !IsNil(o.JsonMetadataSchemaVersion) {
		toSerialize["jsonMetadataSchemaVersion"] = o.JsonMetadataSchemaVersion
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCapabilityDetailsApiResource struct {
	value *CapabilityDetailsApiResource
	isSet bool
}

func (v NullableCapabilityDetailsApiResource) Get() *CapabilityDetailsApiResource {
	return v.value
}

func (v *NullableCapabilityDetailsApiResource) Set(val *CapabilityDetailsApiResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityDetailsApiResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityDetailsApiResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityDetailsApiResource(val *CapabilityDetailsApiResource) *NullableCapabilityDetailsApiResource {
	return &NullableCapabilityDetailsApiResource{value: val, isSet: true}
}

func (v NullableCapabilityDetailsApiResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityDetailsApiResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


