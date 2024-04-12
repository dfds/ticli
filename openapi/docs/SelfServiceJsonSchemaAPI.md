# \SelfServiceJsonSchemaAPI

All URIs are relative to */ssu/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JsonSchemaIdGet**](SelfServiceJsonSchemaAPI.md#JsonSchemaIdGet) | **Get** /json-schema/{id} | 
[**JsonSchemaIdPost**](SelfServiceJsonSchemaAPI.md#JsonSchemaIdPost) | **Post** /json-schema/{id} | 
[**JsonSchemaValidatePost**](SelfServiceJsonSchemaAPI.md#JsonSchemaValidatePost) | **Post** /json-schema/validate | 



## JsonSchemaIdGet

> JsonSchemaIdGet(ctx, id).SchemaVersion(schemaVersion).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	schemaVersion := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SelfServiceJsonSchemaAPI.JsonSchemaIdGet(context.Background(), id).SchemaVersion(schemaVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceJsonSchemaAPI.JsonSchemaIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJsonSchemaIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schemaVersion** | **int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JsonSchemaIdPost

> JsonSchemaIdPost(ctx, id).AddSelfServiceJsonSchemaRequest(addSelfServiceJsonSchemaRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	addSelfServiceJsonSchemaRequest := *openapiclient.NewAddSelfServiceJsonSchemaRequest(map[string]JsonNode{"key": *openapiclient.NewJsonNode()}) // AddSelfServiceJsonSchemaRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SelfServiceJsonSchemaAPI.JsonSchemaIdPost(context.Background(), id).AddSelfServiceJsonSchemaRequest(addSelfServiceJsonSchemaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceJsonSchemaAPI.JsonSchemaIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJsonSchemaIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addSelfServiceJsonSchemaRequest** | [**AddSelfServiceJsonSchemaRequest**](AddSelfServiceJsonSchemaRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JsonSchemaValidatePost

> JsonSchemaValidatePost(ctx).ValidateSelfServiceJsonSchemaRequest(validateSelfServiceJsonSchemaRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	validateSelfServiceJsonSchemaRequest := *openapiclient.NewValidateSelfServiceJsonSchemaRequest(map[string]JsonNode{"key": *openapiclient.NewJsonNode()}) // ValidateSelfServiceJsonSchemaRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SelfServiceJsonSchemaAPI.JsonSchemaValidatePost(context.Background()).ValidateSelfServiceJsonSchemaRequest(validateSelfServiceJsonSchemaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfServiceJsonSchemaAPI.JsonSchemaValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJsonSchemaValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateSelfServiceJsonSchemaRequest** | [**ValidateSelfServiceJsonSchemaRequest**](ValidateSelfServiceJsonSchemaRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

