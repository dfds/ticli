# KafkaTopicApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CapabilityId** | Pointer to **NullableString** |  | [optional] 
**KafkaClusterId** | Pointer to **NullableString** |  | [optional] 
**Partitions** | Pointer to **int32** |  | [optional] 
**Retention** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**KafkaTopicLinks**](KafkaTopicLinks.md) |  | [optional] 

## Methods

### NewKafkaTopicApiResource

`func NewKafkaTopicApiResource() *KafkaTopicApiResource`

NewKafkaTopicApiResource instantiates a new KafkaTopicApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaTopicApiResourceWithDefaults

`func NewKafkaTopicApiResourceWithDefaults() *KafkaTopicApiResource`

NewKafkaTopicApiResourceWithDefaults instantiates a new KafkaTopicApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KafkaTopicApiResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KafkaTopicApiResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KafkaTopicApiResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KafkaTopicApiResource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KafkaTopicApiResource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KafkaTopicApiResource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *KafkaTopicApiResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaTopicApiResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaTopicApiResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KafkaTopicApiResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KafkaTopicApiResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KafkaTopicApiResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *KafkaTopicApiResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KafkaTopicApiResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KafkaTopicApiResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KafkaTopicApiResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KafkaTopicApiResource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KafkaTopicApiResource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCapabilityId

`func (o *KafkaTopicApiResource) GetCapabilityId() string`

GetCapabilityId returns the CapabilityId field if non-nil, zero value otherwise.

### GetCapabilityIdOk

`func (o *KafkaTopicApiResource) GetCapabilityIdOk() (*string, bool)`

GetCapabilityIdOk returns a tuple with the CapabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityId

`func (o *KafkaTopicApiResource) SetCapabilityId(v string)`

SetCapabilityId sets CapabilityId field to given value.

### HasCapabilityId

`func (o *KafkaTopicApiResource) HasCapabilityId() bool`

HasCapabilityId returns a boolean if a field has been set.

### SetCapabilityIdNil

`func (o *KafkaTopicApiResource) SetCapabilityIdNil(b bool)`

 SetCapabilityIdNil sets the value for CapabilityId to be an explicit nil

### UnsetCapabilityId
`func (o *KafkaTopicApiResource) UnsetCapabilityId()`

UnsetCapabilityId ensures that no value is present for CapabilityId, not even an explicit nil
### GetKafkaClusterId

`func (o *KafkaTopicApiResource) GetKafkaClusterId() string`

GetKafkaClusterId returns the KafkaClusterId field if non-nil, zero value otherwise.

### GetKafkaClusterIdOk

`func (o *KafkaTopicApiResource) GetKafkaClusterIdOk() (*string, bool)`

GetKafkaClusterIdOk returns a tuple with the KafkaClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaClusterId

`func (o *KafkaTopicApiResource) SetKafkaClusterId(v string)`

SetKafkaClusterId sets KafkaClusterId field to given value.

### HasKafkaClusterId

`func (o *KafkaTopicApiResource) HasKafkaClusterId() bool`

HasKafkaClusterId returns a boolean if a field has been set.

### SetKafkaClusterIdNil

`func (o *KafkaTopicApiResource) SetKafkaClusterIdNil(b bool)`

 SetKafkaClusterIdNil sets the value for KafkaClusterId to be an explicit nil

### UnsetKafkaClusterId
`func (o *KafkaTopicApiResource) UnsetKafkaClusterId()`

UnsetKafkaClusterId ensures that no value is present for KafkaClusterId, not even an explicit nil
### GetPartitions

`func (o *KafkaTopicApiResource) GetPartitions() int32`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *KafkaTopicApiResource) GetPartitionsOk() (*int32, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *KafkaTopicApiResource) SetPartitions(v int32)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *KafkaTopicApiResource) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetRetention

`func (o *KafkaTopicApiResource) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *KafkaTopicApiResource) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *KafkaTopicApiResource) SetRetention(v string)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *KafkaTopicApiResource) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### SetRetentionNil

`func (o *KafkaTopicApiResource) SetRetentionNil(b bool)`

 SetRetentionNil sets the value for Retention to be an explicit nil

### UnsetRetention
`func (o *KafkaTopicApiResource) UnsetRetention()`

UnsetRetention ensures that no value is present for Retention, not even an explicit nil
### GetStatus

`func (o *KafkaTopicApiResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KafkaTopicApiResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KafkaTopicApiResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KafkaTopicApiResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *KafkaTopicApiResource) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *KafkaTopicApiResource) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLinks

`func (o *KafkaTopicApiResource) GetLinks() KafkaTopicLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *KafkaTopicApiResource) GetLinksOk() (*KafkaTopicLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *KafkaTopicApiResource) SetLinks(v KafkaTopicLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *KafkaTopicApiResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


