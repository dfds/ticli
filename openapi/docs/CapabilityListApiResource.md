# CapabilityListApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]CapabilityListItemApiResource**](CapabilityListItemApiResource.md) |  | [optional] 
**Links** | Pointer to [**CapabilityListLinks**](CapabilityListLinks.md) |  | [optional] 

## Methods

### NewCapabilityListApiResource

`func NewCapabilityListApiResource() *CapabilityListApiResource`

NewCapabilityListApiResource instantiates a new CapabilityListApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityListApiResourceWithDefaults

`func NewCapabilityListApiResourceWithDefaults() *CapabilityListApiResource`

NewCapabilityListApiResourceWithDefaults instantiates a new CapabilityListApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CapabilityListApiResource) GetItems() []CapabilityListItemApiResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CapabilityListApiResource) GetItemsOk() (*[]CapabilityListItemApiResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CapabilityListApiResource) SetItems(v []CapabilityListItemApiResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *CapabilityListApiResource) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *CapabilityListApiResource) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *CapabilityListApiResource) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetLinks

`func (o *CapabilityListApiResource) GetLinks() CapabilityListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CapabilityListApiResource) GetLinksOk() (*CapabilityListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CapabilityListApiResource) SetLinks(v CapabilityListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CapabilityListApiResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


