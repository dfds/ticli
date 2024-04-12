# CapabilityListItemApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**JsonMetadata** | Pointer to **NullableString** |  | [optional] 
**AwsAccountId** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**CapabilityListItemLinks**](CapabilityListItemLinks.md) |  | [optional] 

## Methods

### NewCapabilityListItemApiResource

`func NewCapabilityListItemApiResource() *CapabilityListItemApiResource`

NewCapabilityListItemApiResource instantiates a new CapabilityListItemApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityListItemApiResourceWithDefaults

`func NewCapabilityListItemApiResourceWithDefaults() *CapabilityListItemApiResource`

NewCapabilityListItemApiResourceWithDefaults instantiates a new CapabilityListItemApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CapabilityListItemApiResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapabilityListItemApiResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapabilityListItemApiResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapabilityListItemApiResource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CapabilityListItemApiResource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CapabilityListItemApiResource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CapabilityListItemApiResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilityListItemApiResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilityListItemApiResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilityListItemApiResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CapabilityListItemApiResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CapabilityListItemApiResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *CapabilityListItemApiResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CapabilityListItemApiResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CapabilityListItemApiResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CapabilityListItemApiResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CapabilityListItemApiResource) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CapabilityListItemApiResource) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDescription

`func (o *CapabilityListItemApiResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityListItemApiResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityListItemApiResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityListItemApiResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CapabilityListItemApiResource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CapabilityListItemApiResource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetJsonMetadata

`func (o *CapabilityListItemApiResource) GetJsonMetadata() string`

GetJsonMetadata returns the JsonMetadata field if non-nil, zero value otherwise.

### GetJsonMetadataOk

`func (o *CapabilityListItemApiResource) GetJsonMetadataOk() (*string, bool)`

GetJsonMetadataOk returns a tuple with the JsonMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonMetadata

`func (o *CapabilityListItemApiResource) SetJsonMetadata(v string)`

SetJsonMetadata sets JsonMetadata field to given value.

### HasJsonMetadata

`func (o *CapabilityListItemApiResource) HasJsonMetadata() bool`

HasJsonMetadata returns a boolean if a field has been set.

### SetJsonMetadataNil

`func (o *CapabilityListItemApiResource) SetJsonMetadataNil(b bool)`

 SetJsonMetadataNil sets the value for JsonMetadata to be an explicit nil

### UnsetJsonMetadata
`func (o *CapabilityListItemApiResource) UnsetJsonMetadata()`

UnsetJsonMetadata ensures that no value is present for JsonMetadata, not even an explicit nil
### GetAwsAccountId

`func (o *CapabilityListItemApiResource) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *CapabilityListItemApiResource) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *CapabilityListItemApiResource) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.

### HasAwsAccountId

`func (o *CapabilityListItemApiResource) HasAwsAccountId() bool`

HasAwsAccountId returns a boolean if a field has been set.

### SetAwsAccountIdNil

`func (o *CapabilityListItemApiResource) SetAwsAccountIdNil(b bool)`

 SetAwsAccountIdNil sets the value for AwsAccountId to be an explicit nil

### UnsetAwsAccountId
`func (o *CapabilityListItemApiResource) UnsetAwsAccountId()`

UnsetAwsAccountId ensures that no value is present for AwsAccountId, not even an explicit nil
### GetLinks

`func (o *CapabilityListItemApiResource) GetLinks() CapabilityListItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CapabilityListItemApiResource) GetLinksOk() (*CapabilityListItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CapabilityListItemApiResource) SetLinks(v CapabilityListItemLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CapabilityListItemApiResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


