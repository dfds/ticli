# MembershipApplicationListApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]MembershipApplicationApiResource**](MembershipApplicationApiResource.md) |  | [optional] 
**Links** | Pointer to [**MembershipApplicationListLinks**](MembershipApplicationListLinks.md) |  | [optional] 

## Methods

### NewMembershipApplicationListApiResource

`func NewMembershipApplicationListApiResource() *MembershipApplicationListApiResource`

NewMembershipApplicationListApiResource instantiates a new MembershipApplicationListApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipApplicationListApiResourceWithDefaults

`func NewMembershipApplicationListApiResourceWithDefaults() *MembershipApplicationListApiResource`

NewMembershipApplicationListApiResourceWithDefaults instantiates a new MembershipApplicationListApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *MembershipApplicationListApiResource) GetItems() []MembershipApplicationApiResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MembershipApplicationListApiResource) GetItemsOk() (*[]MembershipApplicationApiResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MembershipApplicationListApiResource) SetItems(v []MembershipApplicationApiResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *MembershipApplicationListApiResource) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *MembershipApplicationListApiResource) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *MembershipApplicationListApiResource) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetLinks

`func (o *MembershipApplicationListApiResource) GetLinks() MembershipApplicationListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MembershipApplicationListApiResource) GetLinksOk() (*MembershipApplicationListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MembershipApplicationListApiResource) SetLinks(v MembershipApplicationListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MembershipApplicationListApiResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


