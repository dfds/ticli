# AddTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**LinkedCapabilityIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddTeamRequest

`func NewAddTeamRequest() *AddTeamRequest`

NewAddTeamRequest instantiates a new AddTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTeamRequestWithDefaults

`func NewAddTeamRequestWithDefaults() *AddTeamRequest`

NewAddTeamRequestWithDefaults instantiates a new AddTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddTeamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddTeamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddTeamRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddTeamRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AddTeamRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddTeamRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AddTeamRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTeamRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTeamRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTeamRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddTeamRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddTeamRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLinkedCapabilityIds

`func (o *AddTeamRequest) GetLinkedCapabilityIds() []string`

GetLinkedCapabilityIds returns the LinkedCapabilityIds field if non-nil, zero value otherwise.

### GetLinkedCapabilityIdsOk

`func (o *AddTeamRequest) GetLinkedCapabilityIdsOk() (*[]string, bool)`

GetLinkedCapabilityIdsOk returns a tuple with the LinkedCapabilityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedCapabilityIds

`func (o *AddTeamRequest) SetLinkedCapabilityIds(v []string)`

SetLinkedCapabilityIds sets LinkedCapabilityIds field to given value.

### HasLinkedCapabilityIds

`func (o *AddTeamRequest) HasLinkedCapabilityIds() bool`

HasLinkedCapabilityIds returns a boolean if a field has been set.

### SetLinkedCapabilityIdsNil

`func (o *AddTeamRequest) SetLinkedCapabilityIdsNil(b bool)`

 SetLinkedCapabilityIdsNil sets the value for LinkedCapabilityIds to be an explicit nil

### UnsetLinkedCapabilityIds
`func (o *AddTeamRequest) UnsetLinkedCapabilityIds()`

UnsetLinkedCapabilityIds ensures that no value is present for LinkedCapabilityIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


