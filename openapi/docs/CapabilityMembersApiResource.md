# CapabilityMembersApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]MemberApiResource**](MemberApiResource.md) |  | [optional] 
**Links** | Pointer to [**CapabilityMembersLinks**](CapabilityMembersLinks.md) |  | [optional] 

## Methods

### NewCapabilityMembersApiResource

`func NewCapabilityMembersApiResource() *CapabilityMembersApiResource`

NewCapabilityMembersApiResource instantiates a new CapabilityMembersApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityMembersApiResourceWithDefaults

`func NewCapabilityMembersApiResourceWithDefaults() *CapabilityMembersApiResource`

NewCapabilityMembersApiResourceWithDefaults instantiates a new CapabilityMembersApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CapabilityMembersApiResource) GetItems() []MemberApiResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CapabilityMembersApiResource) GetItemsOk() (*[]MemberApiResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CapabilityMembersApiResource) SetItems(v []MemberApiResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *CapabilityMembersApiResource) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *CapabilityMembersApiResource) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *CapabilityMembersApiResource) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetLinks

`func (o *CapabilityMembersApiResource) GetLinks() CapabilityMembersLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CapabilityMembersApiResource) GetLinksOk() (*CapabilityMembersLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CapabilityMembersApiResource) SetLinks(v CapabilityMembersLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CapabilityMembersApiResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


