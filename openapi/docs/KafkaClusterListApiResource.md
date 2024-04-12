# KafkaClusterListApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]KafkaClusterApiResource**](KafkaClusterApiResource.md) |  | [optional] 
**Links** | Pointer to [**KafkaClusterListLinks**](KafkaClusterListLinks.md) |  | [optional] 

## Methods

### NewKafkaClusterListApiResource

`func NewKafkaClusterListApiResource() *KafkaClusterListApiResource`

NewKafkaClusterListApiResource instantiates a new KafkaClusterListApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaClusterListApiResourceWithDefaults

`func NewKafkaClusterListApiResourceWithDefaults() *KafkaClusterListApiResource`

NewKafkaClusterListApiResourceWithDefaults instantiates a new KafkaClusterListApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *KafkaClusterListApiResource) GetItems() []KafkaClusterApiResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KafkaClusterListApiResource) GetItemsOk() (*[]KafkaClusterApiResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KafkaClusterListApiResource) SetItems(v []KafkaClusterApiResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *KafkaClusterListApiResource) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *KafkaClusterListApiResource) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *KafkaClusterListApiResource) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetLinks

`func (o *KafkaClusterListApiResource) GetLinks() KafkaClusterListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *KafkaClusterListApiResource) GetLinksOk() (*KafkaClusterListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *KafkaClusterListApiResource) SetLinks(v KafkaClusterListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *KafkaClusterListApiResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


