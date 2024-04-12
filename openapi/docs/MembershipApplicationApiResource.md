# MembershipApplicationApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Applicant** | Pointer to **NullableString** |  | [optional] 
**SubmittedAt** | Pointer to **NullableString** |  | [optional] 
**ExpiresOn** | Pointer to **NullableString** |  | [optional] 
**Approvals** | Pointer to [**MembershipApprovalListApiResource**](MembershipApprovalListApiResource.md) |  | [optional] 
**Links** | Pointer to [**MembershipApplicationLinks**](MembershipApplicationLinks.md) |  | [optional] 

## Methods

### NewMembershipApplicationApiResource

`func NewMembershipApplicationApiResource() *MembershipApplicationApiResource`

NewMembershipApplicationApiResource instantiates a new MembershipApplicationApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipApplicationApiResourceWithDefaults

`func NewMembershipApplicationApiResourceWithDefaults() *MembershipApplicationApiResource`

NewMembershipApplicationApiResourceWithDefaults instantiates a new MembershipApplicationApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MembershipApplicationApiResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MembershipApplicationApiResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MembershipApplicationApiResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MembershipApplicationApiResource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MembershipApplicationApiResource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MembershipApplicationApiResource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetApplicant

`func (o *MembershipApplicationApiResource) GetApplicant() string`

GetApplicant returns the Applicant field if non-nil, zero value otherwise.

### GetApplicantOk

`func (o *MembershipApplicationApiResource) GetApplicantOk() (*string, bool)`

GetApplicantOk returns a tuple with the Applicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicant

`func (o *MembershipApplicationApiResource) SetApplicant(v string)`

SetApplicant sets Applicant field to given value.

### HasApplicant

`func (o *MembershipApplicationApiResource) HasApplicant() bool`

HasApplicant returns a boolean if a field has been set.

### SetApplicantNil

`func (o *MembershipApplicationApiResource) SetApplicantNil(b bool)`

 SetApplicantNil sets the value for Applicant to be an explicit nil

### UnsetApplicant
`func (o *MembershipApplicationApiResource) UnsetApplicant()`

UnsetApplicant ensures that no value is present for Applicant, not even an explicit nil
### GetSubmittedAt

`func (o *MembershipApplicationApiResource) GetSubmittedAt() string`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *MembershipApplicationApiResource) GetSubmittedAtOk() (*string, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *MembershipApplicationApiResource) SetSubmittedAt(v string)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *MembershipApplicationApiResource) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### SetSubmittedAtNil

`func (o *MembershipApplicationApiResource) SetSubmittedAtNil(b bool)`

 SetSubmittedAtNil sets the value for SubmittedAt to be an explicit nil

### UnsetSubmittedAt
`func (o *MembershipApplicationApiResource) UnsetSubmittedAt()`

UnsetSubmittedAt ensures that no value is present for SubmittedAt, not even an explicit nil
### GetExpiresOn

`func (o *MembershipApplicationApiResource) GetExpiresOn() string`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *MembershipApplicationApiResource) GetExpiresOnOk() (*string, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *MembershipApplicationApiResource) SetExpiresOn(v string)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *MembershipApplicationApiResource) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### SetExpiresOnNil

`func (o *MembershipApplicationApiResource) SetExpiresOnNil(b bool)`

 SetExpiresOnNil sets the value for ExpiresOn to be an explicit nil

### UnsetExpiresOn
`func (o *MembershipApplicationApiResource) UnsetExpiresOn()`

UnsetExpiresOn ensures that no value is present for ExpiresOn, not even an explicit nil
### GetApprovals

`func (o *MembershipApplicationApiResource) GetApprovals() MembershipApprovalListApiResource`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *MembershipApplicationApiResource) GetApprovalsOk() (*MembershipApprovalListApiResource, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *MembershipApplicationApiResource) SetApprovals(v MembershipApprovalListApiResource)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *MembershipApplicationApiResource) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetLinks

`func (o *MembershipApplicationApiResource) GetLinks() MembershipApplicationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MembershipApplicationApiResource) GetLinksOk() (*MembershipApplicationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MembershipApplicationApiResource) SetLinks(v MembershipApplicationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MembershipApplicationApiResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


