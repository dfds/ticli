# SetCapabilityMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonMetadata** | [**map[string]JsonNode**](JsonNode.md) |  | 

## Methods

### NewSetCapabilityMetadataRequest

`func NewSetCapabilityMetadataRequest(jsonMetadata map[string]JsonNode, ) *SetCapabilityMetadataRequest`

NewSetCapabilityMetadataRequest instantiates a new SetCapabilityMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCapabilityMetadataRequestWithDefaults

`func NewSetCapabilityMetadataRequestWithDefaults() *SetCapabilityMetadataRequest`

NewSetCapabilityMetadataRequestWithDefaults instantiates a new SetCapabilityMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonMetadata

`func (o *SetCapabilityMetadataRequest) GetJsonMetadata() map[string]JsonNode`

GetJsonMetadata returns the JsonMetadata field if non-nil, zero value otherwise.

### GetJsonMetadataOk

`func (o *SetCapabilityMetadataRequest) GetJsonMetadataOk() (*map[string]JsonNode, bool)`

GetJsonMetadataOk returns a tuple with the JsonMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonMetadata

`func (o *SetCapabilityMetadataRequest) SetJsonMetadata(v map[string]JsonNode)`

SetJsonMetadata sets JsonMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


