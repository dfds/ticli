# MembershipApprovalListApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]MembershipApprovalApiResource**](MembershipApprovalApiResource.md) |  | [optional] 
**Links** | Pointer to [**MembershipApprovalListLinks**](MembershipApprovalListLinks.md) |  | [optional] 

## Methods

### NewMembershipApprovalListApiResource

`func NewMembershipApprovalListApiResource() *MembershipApprovalListApiResource`

NewMembershipApprovalListApiResource instantiates a new MembershipApprovalListApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipApprovalListApiResourceWithDefaults

`func NewMembershipApprovalListApiResourceWithDefaults() *MembershipApprovalListApiResource`

NewMembershipApprovalListApiResourceWithDefaults instantiates a new MembershipApprovalListApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *MembershipApprovalListApiResource) GetItems() []MembershipApprovalApiResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MembershipApprovalListApiResource) GetItemsOk() (*[]MembershipApprovalApiResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MembershipApprovalListApiResource) SetItems(v []MembershipApprovalApiResource)`

SetItems sets Items field to given value.

### HasItems

`func (o *MembershipApprovalListApiResource) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *MembershipApprovalListApiResource) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *MembershipApprovalListApiResource) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetLinks

`func (o *MembershipApprovalListApiResource) GetLinks() MembershipApprovalListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MembershipApprovalListApiResource) GetLinksOk() (*MembershipApprovalListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MembershipApprovalListApiResource) SetLinks(v MembershipApprovalListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MembershipApprovalListApiResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


