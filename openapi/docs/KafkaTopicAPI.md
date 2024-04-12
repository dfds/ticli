# \KafkaTopicAPI

All URIs are relative to */ssu/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KafkatopicsGet**](KafkaTopicAPI.md#KafkatopicsGet) | **Get** /kafkatopics | 
[**KafkatopicsIdConsumersGet**](KafkaTopicAPI.md#KafkatopicsIdConsumersGet) | **Get** /kafkatopics/{id}/consumers | 
[**KafkatopicsIdDelete**](KafkaTopicAPI.md#KafkatopicsIdDelete) | **Delete** /kafkatopics/{id} | 
[**KafkatopicsIdDescriptionPut**](KafkaTopicAPI.md#KafkatopicsIdDescriptionPut) | **Put** /kafkatopics/{id}/description | 
[**KafkatopicsIdGet**](KafkaTopicAPI.md#KafkatopicsIdGet) | **Get** /kafkatopics/{id} | 
[**KafkatopicsIdMessagecontractsContractIdGet**](KafkaTopicAPI.md#KafkatopicsIdMessagecontractsContractIdGet) | **Get** /kafkatopics/{id}/messagecontracts/{contractId} | 
[**KafkatopicsIdMessagecontractsContractIdRetryPost**](KafkaTopicAPI.md#KafkatopicsIdMessagecontractsContractIdRetryPost) | **Post** /kafkatopics/{id}/messagecontracts/{contractId}/retry | 
[**KafkatopicsIdMessagecontractsGet**](KafkaTopicAPI.md#KafkatopicsIdMessagecontractsGet) | **Get** /kafkatopics/{id}/messagecontracts | 
[**KafkatopicsIdMessagecontractsPost**](KafkaTopicAPI.md#KafkatopicsIdMessagecontractsPost) | **Post** /kafkatopics/{id}/messagecontracts | 
[**KafkatopicsIdMessagecontractsValidatePost**](KafkaTopicAPI.md#KafkatopicsIdMessagecontractsValidatePost) | **Post** /kafkatopics/{id}/messagecontracts-validate | 



## KafkatopicsGet

> KafkaTopicListApiResource KafkatopicsGet(ctx).CapabilityId(capabilityId).ClusterId(clusterId).IncludePrivate(includePrivate).Execute()



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
	capabilityId := "capabilityId_example" // string |  (optional)
	clusterId := "clusterId_example" // string |  (optional)
	includePrivate := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KafkaTopicAPI.KafkatopicsGet(context.Background()).CapabilityId(capabilityId).ClusterId(clusterId).IncludePrivate(includePrivate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicAPI.KafkatopicsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KafkatopicsGet`: KafkaTopicListApiResource
	fmt.Fprintf(os.Stdout, "Response from `KafkaTopicAPI.KafkatopicsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKafkatopicsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **capabilityId** | **string** |  | 
 **clusterId** | **string** |  | 
 **includePrivate** | **bool** |  | 

### Return type

[**KafkaTopicListApiResource**](KafkaTopicListApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KafkatopicsIdConsumersGet

> KafkaTopicApiResource KafkatopicsIdConsumersGet(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KafkaTopicAPI.KafkatopicsIdConsumersGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicAPI.KafkatopicsIdConsumersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KafkatopicsIdConsumersGet`: KafkaTopicApiResource
	fmt.Fprintf(os.Stdout, "Response from `KafkaTopicAPI.KafkatopicsIdConsumersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKafkatopicsIdConsumersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaTopicApiResource**](KafkaTopicApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KafkatopicsIdDelete

> KafkaTopicApiResource KafkatopicsIdDelete(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KafkaTopicAPI.KafkatopicsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicAPI.KafkatopicsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KafkatopicsIdDelete`: KafkaTopicApiResource
	fmt.Fprintf(os.Stdout, "Response from `KafkaTopicAPI.KafkatopicsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKafkatopicsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaTopicApiResource**](KafkaTopicApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KafkatopicsIdDescriptionPut

> KafkaTopicApiResource KafkatopicsIdDescriptionPut(ctx, id).ChangeKafkaTopicDescriptionRequest(changeKafkaTopicDescriptionRequest).Execute()



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
	changeKafkaTopicDescriptionRequest := *openapiclient.NewChangeKafkaTopicDescriptionRequest() // ChangeKafkaTopicDescriptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KafkaTopicAPI.KafkatopicsIdDescriptionPut(context.Background(), id).ChangeKafkaTopicDescriptionRequest(changeKafkaTopicDescriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicAPI.KafkatopicsIdDescriptionPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KafkatopicsIdDescriptionPut`: KafkaTopicApiResource
	fmt.Fprintf(os.Stdout, "Response from `KafkaTopicAPI.KafkatopicsIdDescriptionPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKafkatopicsIdDescriptionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeKafkaTopicDescriptionRequest** | [**ChangeKafkaTopicDescriptionRequest**](ChangeKafkaTopicDescriptionRequest.md) |  | 

### Return type

[**KafkaTopicApiResource**](KafkaTopicApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KafkatopicsIdGet

> KafkaTopicApiResource KafkatopicsIdGet(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KafkaTopicAPI.KafkatopicsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicAPI.KafkatopicsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KafkatopicsIdGet`: KafkaTopicApiResource
	fmt.Fprintf(os.Stdout, "Response from `KafkaTopicAPI.KafkatopicsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKafkatopicsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaTopicApiResource**](KafkaTopicApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KafkatopicsIdMessagecontractsContractIdGet

> KafkatopicsIdMessagecontractsContractIdGet(ctx, id, contractId).Execute()



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
	contractId := "contractId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KafkaTopicAPI.KafkatopicsIdMessagecontractsContractIdGet(context.Background(), id, contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicAPI.KafkatopicsIdMessagecontractsContractIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKafkatopicsIdMessagecontractsContractIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## KafkatopicsIdMessagecontractsContractIdRetryPost

> KafkaTopicApiResource KafkatopicsIdMessagecontractsContractIdRetryPost(ctx, id, contractId).Execute()



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
	contractId := "contractId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KafkaTopicAPI.KafkatopicsIdMessagecontractsContractIdRetryPost(context.Background(), id, contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicAPI.KafkatopicsIdMessagecontractsContractIdRetryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KafkatopicsIdMessagecontractsContractIdRetryPost`: KafkaTopicApiResource
	fmt.Fprintf(os.Stdout, "Response from `KafkaTopicAPI.KafkatopicsIdMessagecontractsContractIdRetryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**contractId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKafkatopicsIdMessagecontractsContractIdRetryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KafkaTopicApiResource**](KafkaTopicApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KafkatopicsIdMessagecontractsGet

> KafkaTopicApiResource KafkatopicsIdMessagecontractsGet(ctx, id).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KafkaTopicAPI.KafkatopicsIdMessagecontractsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicAPI.KafkatopicsIdMessagecontractsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KafkatopicsIdMessagecontractsGet`: KafkaTopicApiResource
	fmt.Fprintf(os.Stdout, "Response from `KafkaTopicAPI.KafkatopicsIdMessagecontractsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKafkatopicsIdMessagecontractsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaTopicApiResource**](KafkaTopicApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KafkatopicsIdMessagecontractsPost

> KafkaTopicApiResource KafkatopicsIdMessagecontractsPost(ctx, id).NewMessageContractRequest(newMessageContractRequest).Execute()



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
	newMessageContractRequest := *openapiclient.NewNewMessageContractRequest("MessageType_example", "Description_example", "Example_example", "Schema_example") // NewMessageContractRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KafkaTopicAPI.KafkatopicsIdMessagecontractsPost(context.Background(), id).NewMessageContractRequest(newMessageContractRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicAPI.KafkatopicsIdMessagecontractsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KafkatopicsIdMessagecontractsPost`: KafkaTopicApiResource
	fmt.Fprintf(os.Stdout, "Response from `KafkaTopicAPI.KafkatopicsIdMessagecontractsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKafkatopicsIdMessagecontractsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newMessageContractRequest** | [**NewMessageContractRequest**](NewMessageContractRequest.md) |  | 

### Return type

[**KafkaTopicApiResource**](KafkaTopicApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KafkatopicsIdMessagecontractsValidatePost

> KafkaTopicApiResource KafkatopicsIdMessagecontractsValidatePost(ctx, id).ValidateMessageContractRequest(validateMessageContractRequest).Execute()



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
	validateMessageContractRequest := *openapiclient.NewValidateMessageContractRequest("MessageType_example", "Schema_example") // ValidateMessageContractRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KafkaTopicAPI.KafkatopicsIdMessagecontractsValidatePost(context.Background(), id).ValidateMessageContractRequest(validateMessageContractRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicAPI.KafkatopicsIdMessagecontractsValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KafkatopicsIdMessagecontractsValidatePost`: KafkaTopicApiResource
	fmt.Fprintf(os.Stdout, "Response from `KafkaTopicAPI.KafkatopicsIdMessagecontractsValidatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKafkatopicsIdMessagecontractsValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validateMessageContractRequest** | [**ValidateMessageContractRequest**](ValidateMessageContractRequest.md) |  | 

### Return type

[**KafkaTopicApiResource**](KafkaTopicApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

