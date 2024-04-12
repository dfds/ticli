# \ECRRepositoriesAPI

All URIs are relative to */ssu/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EcrRepositoriesGet**](ECRRepositoriesAPI.md#EcrRepositoriesGet) | **Get** /ecr/repositories | 
[**EcrRepositoriesPost**](ECRRepositoriesAPI.md#EcrRepositoriesPost) | **Post** /ecr/repositories | 
[**EcrSynchronizePost**](ECRRepositoriesAPI.md#EcrSynchronizePost) | **Post** /ecr/synchronize | 



## EcrRepositoriesGet

> EcrRepositoriesGet(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ECRRepositoriesAPI.EcrRepositoriesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ECRRepositoriesAPI.EcrRepositoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEcrRepositoriesGetRequest struct via the builder pattern


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


## EcrRepositoriesPost

> EcrRepositoriesPost(ctx).NewECRRepositoryRequest(newECRRepositoryRequest).Execute()



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
	newECRRepositoryRequest := *openapiclient.NewNewECRRepositoryRequest("Name_example", "Description_example") // NewECRRepositoryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ECRRepositoriesAPI.EcrRepositoriesPost(context.Background()).NewECRRepositoryRequest(newECRRepositoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ECRRepositoriesAPI.EcrRepositoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEcrRepositoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newECRRepositoryRequest** | [**NewECRRepositoryRequest**](NewECRRepositoryRequest.md) |  | 

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


## EcrSynchronizePost

> EcrSynchronizePost(ctx).UpdateOnMismatch(updateOnMismatch).Execute()



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
	updateOnMismatch := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ECRRepositoriesAPI.EcrSynchronizePost(context.Background()).UpdateOnMismatch(updateOnMismatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ECRRepositoriesAPI.EcrSynchronizePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEcrSynchronizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOnMismatch** | **bool** |  | 

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

