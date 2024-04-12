# KafkaTopicLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**MessageContracts** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Consumers** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**UpdateDescription** | Pointer to [**ResourceActionLink**](ResourceActionLink.md) |  | [optional] 

## Methods

### NewKafkaTopicLinks

`func NewKafkaTopicLinks() *KafkaTopicLinks`

NewKafkaTopicLinks instantiates a new KafkaTopicLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaTopicLinksWithDefaults

`func NewKafkaTopicLinksWithDefaults() *KafkaTopicLinks`

NewKafkaTopicLinksWithDefaults instantiates a new KafkaTopicLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *KafkaTopicLinks) GetSelf() ResourceLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *KafkaTopicLinks) GetSelfOk() (*ResourceLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *KafkaTopicLinks) SetSelf(v ResourceLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *KafkaTopicLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetMessageContracts

`func (o *KafkaTopicLinks) GetMessageContracts() ResourceLink`

GetMessageContracts returns the MessageContracts field if non-nil, zero value otherwise.

### GetMessageContractsOk

`func (o *KafkaTopicLinks) GetMessageContractsOk() (*ResourceLink, bool)`

GetMessageContractsOk returns a tuple with the MessageContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageContracts

`func (o *KafkaTopicLinks) SetMessageContracts(v ResourceLink)`

SetMessageContracts sets MessageContracts field to given value.

### HasMessageContracts

`func (o *KafkaTopicLinks) HasMessageContracts() bool`

HasMessageContracts returns a boolean if a field has been set.

### GetConsumers

`func (o *KafkaTopicLinks) GetConsumers() ResourceLink`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *KafkaTopicLinks) GetConsumersOk() (*ResourceLink, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *KafkaTopicLinks) SetConsumers(v ResourceLink)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *KafkaTopicLinks) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetUpdateDescription

`func (o *KafkaTopicLinks) GetUpdateDescription() ResourceActionLink`

GetUpdateDescription returns the UpdateDescription field if non-nil, zero value otherwise.

### GetUpdateDescriptionOk

`func (o *KafkaTopicLinks) GetUpdateDescriptionOk() (*ResourceActionLink, bool)`

GetUpdateDescriptionOk returns a tuple with the UpdateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDescription

`func (o *KafkaTopicLinks) SetUpdateDescription(v ResourceActionLink)`

SetUpdateDescription sets UpdateDescription field to given value.

### HasUpdateDescription

`func (o *KafkaTopicLinks) HasUpdateDescription() bool`

HasUpdateDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


