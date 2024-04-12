# CapabilityDetailsApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**JsonMetadata** | Pointer to **NullableString** |  | [optional] 
**JsonMetadataSchemaVersion** | Pointer to **int32** |  | [optional] 
**Links** | Pointer to [**CapabilityDetailsLinks**](CapabilityDetailsLinks.md) |  | [optional] 

## Methods

### NewCapabilityDetailsApiResource

`func NewCapabilityDetailsApiResource() *CapabilityDetailsApiResource`

NewCapabilityDetailsApiResource instantiates a new CapabilityDetailsApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDetailsApiResourceWithDefaults

`func NewCapabilityDetailsApiResourceWithDefaults() *CapabilityDetailsApiResource`

NewCapabilityDetailsApiResourceWithDefaults instantiates a new CapabilityDetailsApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CapabilityDetailsApiResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapabilityDetailsApiResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapabilityDetailsApiResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapabilityDetailsApiResource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CapabilityDetailsApiResource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CapabilityDetailsApiResource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CapabilityDetailsApiResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilityDetailsApiResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilityDetailsApiResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilityDetailsApiResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CapabilityDetailsApiResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CapabilityDetailsApiResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *CapabilityDetailsApiResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CapabilityDetailsApiResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CapabilityDetailsApiResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CapabilityDetailsApiResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CapabilityDetailsApiResource) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CapabilityDetailsApiResource) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDescription

`func (o *CapabilityDetailsApiResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityDetailsApiResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityDetailsApiResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityDetailsApiResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CapabilityDetailsApiResource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CapabilityDetailsApiResource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetJsonMetadata

`func (o *CapabilityDetailsApiResource) GetJsonMetadata() string`

GetJsonMetadata returns the JsonMetadata field if non-nil, zero value otherwise.

### GetJsonMetadataOk

`func (o *CapabilityDetailsApiResource) GetJsonMetadataOk() (*string, bool)`

GetJsonMetadataOk returns a tuple with the JsonMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonMetadata

`func (o *CapabilityDetailsApiResource) SetJsonMetadata(v string)`

SetJsonMetadata sets JsonMetadata field to given value.

### HasJsonMetadata

`func (o *CapabilityDetailsApiResource) HasJsonMetadata() bool`

HasJsonMetadata returns a boolean if a field has been set.

### SetJsonMetadataNil

`func (o *CapabilityDetailsApiResource) SetJsonMetadataNil(b bool)`

 SetJsonMetadataNil sets the value for JsonMetadata to be an explicit nil

### UnsetJsonMetadata
`func (o *CapabilityDetailsApiResource) UnsetJsonMetadata()`

UnsetJsonMetadata ensures that no value is present for JsonMetadata, not even an explicit nil
### GetJsonMetadataSchemaVersion

`func (o *CapabilityDetailsApiResource) GetJsonMetadataSchemaVersion() int32`

GetJsonMetadataSchemaVersion returns the JsonMetadataSchemaVersion field if non-nil, zero value otherwise.

### GetJsonMetadataSchemaVersionOk

`func (o *CapabilityDetailsApiResource) GetJsonMetadataSchemaVersionOk() (*int32, bool)`

GetJsonMetadataSchemaVersionOk returns a tuple with the JsonMetadataSchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonMetadataSchemaVersion

`func (o *CapabilityDetailsApiResource) SetJsonMetadataSchemaVersion(v int32)`

SetJsonMetadataSchemaVersion sets JsonMetadataSchemaVersion field to given value.

### HasJsonMetadataSchemaVersion

`func (o *CapabilityDetailsApiResource) HasJsonMetadataSchemaVersion() bool`

HasJsonMetadataSchemaVersion returns a boolean if a field has been set.

### GetLinks

`func (o *CapabilityDetailsApiResource) GetLinks() CapabilityDetailsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CapabilityDetailsApiResource) GetLinksOk() (*CapabilityDetailsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CapabilityDetailsApiResource) SetLinks(v CapabilityDetailsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CapabilityDetailsApiResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


