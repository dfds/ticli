# NewKafkaTopicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KafkaClusterId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Partitions** | **int32** |  | 
**Retention** | **string** |  | 

## Methods

### NewNewKafkaTopicRequest

`func NewNewKafkaTopicRequest(kafkaClusterId string, name string, description string, partitions int32, retention string, ) *NewKafkaTopicRequest`

NewNewKafkaTopicRequest instantiates a new NewKafkaTopicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewKafkaTopicRequestWithDefaults

`func NewNewKafkaTopicRequestWithDefaults() *NewKafkaTopicRequest`

NewNewKafkaTopicRequestWithDefaults instantiates a new NewKafkaTopicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKafkaClusterId

`func (o *NewKafkaTopicRequest) GetKafkaClusterId() string`

GetKafkaClusterId returns the KafkaClusterId field if non-nil, zero value otherwise.

### GetKafkaClusterIdOk

`func (o *NewKafkaTopicRequest) GetKafkaClusterIdOk() (*string, bool)`

GetKafkaClusterIdOk returns a tuple with the KafkaClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaClusterId

`func (o *NewKafkaTopicRequest) SetKafkaClusterId(v string)`

SetKafkaClusterId sets KafkaClusterId field to given value.


### GetName

`func (o *NewKafkaTopicRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewKafkaTopicRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewKafkaTopicRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NewKafkaTopicRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewKafkaTopicRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewKafkaTopicRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPartitions

`func (o *NewKafkaTopicRequest) GetPartitions() int32`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *NewKafkaTopicRequest) GetPartitionsOk() (*int32, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *NewKafkaTopicRequest) SetPartitions(v int32)`

SetPartitions sets Partitions field to given value.


### GetRetention

`func (o *NewKafkaTopicRequest) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *NewKafkaTopicRequest) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *NewKafkaTopicRequest) SetRetention(v string)`

SetRetention sets Retention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


