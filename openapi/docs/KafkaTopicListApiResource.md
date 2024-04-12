# KafkaTopicListApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]KafkaTopicApiResource**](KafkaTopicApiResource.md) |  | [optional] 
**Embedded** | Pointer to [**KafkaTopicListEmbeddedResources**](KafkaTopicListEmbeddedResources.md) |  | [optional] 
**Links** | Pointer to [**KafkaTopicListLinks**](KafkaTopicListLinks.md) |  | [optional] 

## Methods

### NewKafkaTopicListApiResource

`func NewKafkaTopicListApiResource() *KafkaTopicListApiResource`

NewKafkaTopicListApiResource instantiates a new KafkaTopicListApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaTopicListApiResourceWithDefaults

`func NewKafkaTopicListApiResourceWithDefaults() *KafkaTopicListApiResource`

NewKafkaTopicListApiResourceWithDefaults instantiates a new KafkaTopicListApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *KafkaTopicListApiResource) GetItems() []KafkaTopicApiResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KafkaTopicListApiResource) GetItemsOk() (*[]KafkaTopicApiResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KafkaTopicListApiResource) SetItems(v []KafkaTopicApiResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *KafkaTopicListApiResource) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *KafkaTopicListApiResource) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *KafkaTopicListApiResource) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetEmbedded

`func (o *KafkaTopicListApiResource) GetEmbedded() KafkaTopicListEmbeddedResources`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *KafkaTopicListApiResource) GetEmbeddedOk() (*KafkaTopicListEmbeddedResources, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *KafkaTopicListApiResource) SetEmbedded(v KafkaTopicListEmbeddedResources)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *KafkaTopicListApiResource) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *KafkaTopicListApiResource) GetLinks() KafkaTopicListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *KafkaTopicListApiResource) GetLinksOk() (*KafkaTopicListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *KafkaTopicListApiResource) SetLinks(v KafkaTopicListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *KafkaTopicListApiResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


