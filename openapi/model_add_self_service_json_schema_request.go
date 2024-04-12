/*
SelfService API

SelfService API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AddSelfServiceJsonSchemaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSelfServiceJsonSchemaRequest{}

// AddSelfServiceJsonSchemaRequest struct for AddSelfServiceJsonSchemaRequest
type AddSelfServiceJsonSchemaRequest struct {
	Schema map[string]JsonNode `json:"schema"`
}

type _AddSelfServiceJsonSchemaRequest AddSelfServiceJsonSchemaRequest

// NewAddSelfServiceJsonSchemaRequest instantiates a new AddSelfServiceJsonSchemaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSelfServiceJsonSchemaRequest(schema map[string]JsonNode) *AddSelfServiceJsonSchemaRequest {
	this := AddSelfServiceJsonSchemaRequest{}
	this.Schema = schema
	return &this
}

// NewAddSelfServiceJsonSchemaRequestWithDefaults instantiates a new AddSelfServiceJsonSchemaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSelfServiceJsonSchemaRequestWithDefaults() *AddSelfServiceJsonSchemaRequest {
	this := AddSelfServiceJsonSchemaRequest{}
	return &this
}

// GetSchema returns the Schema field value
func (o *AddSelfServiceJsonSchemaRequest) GetSchema() map[string]JsonNode {
	if o == nil {
		var ret map[string]JsonNode
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *AddSelfServiceJsonSchemaRequest) GetSchemaOk() (*map[string]JsonNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value
func (o *AddSelfServiceJsonSchemaRequest) SetSchema(v map[string]JsonNode) {
	o.Schema = v
}

func (o AddSelfServiceJsonSchemaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSelfServiceJsonSchemaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schema"] = o.Schema
	return toSerialize, nil
}

func (o *AddSelfServiceJsonSchemaRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schema",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddSelfServiceJsonSchemaRequest := _AddSelfServiceJsonSchemaRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddSelfServiceJsonSchemaRequest)

	if err != nil {
		return err
	}

	*o = AddSelfServiceJsonSchemaRequest(varAddSelfServiceJsonSchemaRequest)

	return err
}

type NullableAddSelfServiceJsonSchemaRequest struct {
	value *AddSelfServiceJsonSchemaRequest
	isSet bool
}

func (v NullableAddSelfServiceJsonSchemaRequest) Get() *AddSelfServiceJsonSchemaRequest {
	return v.value
}

func (v *NullableAddSelfServiceJsonSchemaRequest) Set(val *AddSelfServiceJsonSchemaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSelfServiceJsonSchemaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSelfServiceJsonSchemaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSelfServiceJsonSchemaRequest(val *AddSelfServiceJsonSchemaRequest) *NullableAddSelfServiceJsonSchemaRequest {
	return &NullableAddSelfServiceJsonSchemaRequest{value: val, isSet: true}
}

func (v NullableAddSelfServiceJsonSchemaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSelfServiceJsonSchemaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


