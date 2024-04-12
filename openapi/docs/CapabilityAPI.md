# \CapabilityAPI

All URIs are relative to */ssu/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CapabilitiesGet**](CapabilityAPI.md#CapabilitiesGet) | **Get** /capabilities | 
[**CapabilitiesIdAwsaccountGet**](CapabilityAPI.md#CapabilitiesIdAwsaccountGet) | **Get** /capabilities/{id}/awsaccount | 
[**CapabilitiesIdAwsaccountPost**](CapabilityAPI.md#CapabilitiesIdAwsaccountPost) | **Post** /capabilities/{id}/awsaccount | 
[**CapabilitiesIdCanceldeletionrequestPost**](CapabilityAPI.md#CapabilitiesIdCanceldeletionrequestPost) | **Post** /capabilities/{id}/canceldeletionrequest | 
[**CapabilitiesIdConfigurationlevelGet**](CapabilityAPI.md#CapabilitiesIdConfigurationlevelGet) | **Get** /capabilities/{id}/configurationlevel | 
[**CapabilitiesIdGet**](CapabilityAPI.md#CapabilitiesIdGet) | **Get** /capabilities/{id} | 
[**CapabilitiesIdInvitationsPost**](CapabilityAPI.md#CapabilitiesIdInvitationsPost) | **Post** /capabilities/{id}/invitations | 
[**CapabilitiesIdJoinPost**](CapabilityAPI.md#CapabilitiesIdJoinPost) | **Post** /capabilities/{id}/join | 
[**CapabilitiesIdKafkaclusteraccessClusterIdGet**](CapabilityAPI.md#CapabilitiesIdKafkaclusteraccessClusterIdGet) | **Get** /capabilities/{id}/kafkaclusteraccess/{clusterId} | 
[**CapabilitiesIdKafkaclusteraccessClusterIdPost**](CapabilityAPI.md#CapabilitiesIdKafkaclusteraccessClusterIdPost) | **Post** /capabilities/{id}/kafkaclusteraccess/{clusterId} | 
[**CapabilitiesIdKafkaclusteraccessGet**](CapabilityAPI.md#CapabilitiesIdKafkaclusteraccessGet) | **Get** /capabilities/{id}/kafkaclusteraccess | 
[**CapabilitiesIdLeavePost**](CapabilityAPI.md#CapabilitiesIdLeavePost) | **Post** /capabilities/{id}/leave | 
[**CapabilitiesIdMembersGet**](CapabilityAPI.md#CapabilitiesIdMembersGet) | **Get** /capabilities/{id}/members | 
[**CapabilitiesIdMembershipapplicationsGet**](CapabilityAPI.md#CapabilitiesIdMembershipapplicationsGet) | **Get** /capabilities/{id}/membershipapplications | 
[**CapabilitiesIdMembershipapplicationsPost**](CapabilityAPI.md#CapabilitiesIdMembershipapplicationsPost) | **Post** /capabilities/{id}/membershipapplications | 
[**CapabilitiesIdMetadataGet**](CapabilityAPI.md#CapabilitiesIdMetadataGet) | **Get** /capabilities/{id}/metadata | 
[**CapabilitiesIdMetadataPost**](CapabilityAPI.md#CapabilitiesIdMetadataPost) | **Post** /capabilities/{id}/metadata | 
[**CapabilitiesIdRequestdeletionPost**](CapabilityAPI.md#CapabilitiesIdRequestdeletionPost) | **Post** /capabilities/{id}/requestdeletion | 
[**CapabilitiesIdRequiredMetadataPost**](CapabilityAPI.md#CapabilitiesIdRequiredMetadataPost) | **Post** /capabilities/{id}/required-metadata | 
[**CapabilitiesIdTeamsGet**](CapabilityAPI.md#CapabilitiesIdTeamsGet) | **Get** /capabilities/{id}/teams | 
[**CapabilitiesIdTopicsPost**](CapabilityAPI.md#CapabilitiesIdTopicsPost) | **Post** /capabilities/{id}/topics | 
[**CapabilitiesPost**](CapabilityAPI.md#CapabilitiesPost) | **Post** /capabilities | 



## CapabilitiesGet

> CapabilityListApiResource CapabilitiesGet(ctx).Execute()



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
	resp, r, err := apiClient.CapabilityAPI.CapabilitiesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesGet`: CapabilityListApiResource
	fmt.Fprintf(os.Stdout, "Response from `CapabilityAPI.CapabilitiesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesGetRequest struct via the builder pattern


### Return type

[**CapabilityListApiResource**](CapabilityListApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesIdAwsaccountGet

> AwsAccountApiResource CapabilitiesIdAwsaccountGet(ctx, id).Execute()



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
	resp, r, err := apiClient.CapabilityAPI.CapabilitiesIdAwsaccountGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdAwsaccountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesIdAwsaccountGet`: AwsAccountApiResource
	fmt.Fprintf(os.Stdout, "Response from `CapabilityAPI.CapabilitiesIdAwsaccountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesIdAwsaccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsAccountApiResource**](AwsAccountApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesIdAwsaccountPost

> AwsAccountApiResource CapabilitiesIdAwsaccountPost(ctx, id).Execute()



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
	resp, r, err := apiClient.CapabilityAPI.CapabilitiesIdAwsaccountPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdAwsaccountPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesIdAwsaccountPost`: AwsAccountApiResource
	fmt.Fprintf(os.Stdout, "Response from `CapabilityAPI.CapabilitiesIdAwsaccountPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesIdAwsaccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsAccountApiResource**](AwsAccountApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesIdCanceldeletionrequestPost

> CapabilitiesIdCanceldeletionrequestPost(ctx, id).Execute()



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
	r, err := apiClient.CapabilityAPI.CapabilitiesIdCanceldeletionrequestPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdCanceldeletionrequestPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCapabilitiesIdCanceldeletionrequestPostRequest struct via the builder pattern


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


## CapabilitiesIdConfigurationlevelGet

> CapabilitiesIdConfigurationlevelGet(ctx, id).Execute()



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
	r, err := apiClient.CapabilityAPI.CapabilitiesIdConfigurationlevelGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdConfigurationlevelGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCapabilitiesIdConfigurationlevelGetRequest struct via the builder pattern


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


## CapabilitiesIdGet

> CapabilityDetailsApiResource CapabilitiesIdGet(ctx, id).Execute()



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
	resp, r, err := apiClient.CapabilityAPI.CapabilitiesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesIdGet`: CapabilityDetailsApiResource
	fmt.Fprintf(os.Stdout, "Response from `CapabilityAPI.CapabilitiesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CapabilityDetailsApiResource**](CapabilityDetailsApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesIdInvitationsPost

> CapabilityDetailsApiResource CapabilitiesIdInvitationsPost(ctx, id).InvitationsRequest(invitationsRequest).Execute()



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
	invitationsRequest := *openapiclient.NewInvitationsRequest([]string{"Invitees_example"}) // InvitationsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CapabilityAPI.CapabilitiesIdInvitationsPost(context.Background(), id).InvitationsRequest(invitationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdInvitationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesIdInvitationsPost`: CapabilityDetailsApiResource
	fmt.Fprintf(os.Stdout, "Response from `CapabilityAPI.CapabilitiesIdInvitationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesIdInvitationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitationsRequest** | [**InvitationsRequest**](InvitationsRequest.md) |  | 

### Return type

[**CapabilityDetailsApiResource**](CapabilityDetailsApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesIdJoinPost

> CapabilitiesIdJoinPost(ctx, id).Execute()



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
	r, err := apiClient.CapabilityAPI.CapabilitiesIdJoinPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdJoinPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCapabilitiesIdJoinPostRequest struct via the builder pattern


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


## CapabilitiesIdKafkaclusteraccessClusterIdGet

> CapabilitiesIdKafkaclusteraccessClusterIdGet(ctx, id, clusterId).Execute()



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
	clusterId := "clusterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CapabilityAPI.CapabilitiesIdKafkaclusteraccessClusterIdGet(context.Background(), id, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdKafkaclusteraccessClusterIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesIdKafkaclusteraccessClusterIdGetRequest struct via the builder pattern


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


## CapabilitiesIdKafkaclusteraccessClusterIdPost

> CapabilitiesIdKafkaclusteraccessClusterIdPost(ctx, id, clusterId).Execute()



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
	clusterId := "clusterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CapabilityAPI.CapabilitiesIdKafkaclusteraccessClusterIdPost(context.Background(), id, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdKafkaclusteraccessClusterIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesIdKafkaclusteraccessClusterIdPostRequest struct via the builder pattern


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


## CapabilitiesIdKafkaclusteraccessGet

> CapabilitiesIdKafkaclusteraccessGet(ctx, id).Execute()



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
	r, err := apiClient.CapabilityAPI.CapabilitiesIdKafkaclusteraccessGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdKafkaclusteraccessGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCapabilitiesIdKafkaclusteraccessGetRequest struct via the builder pattern


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


## CapabilitiesIdLeavePost

> CapabilitiesIdLeavePost(ctx, id).Execute()



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
	r, err := apiClient.CapabilityAPI.CapabilitiesIdLeavePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdLeavePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCapabilitiesIdLeavePostRequest struct via the builder pattern


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


## CapabilitiesIdMembersGet

> CapabilityMembersApiResource CapabilitiesIdMembersGet(ctx, id).Execute()



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
	resp, r, err := apiClient.CapabilityAPI.CapabilitiesIdMembersGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdMembersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesIdMembersGet`: CapabilityMembersApiResource
	fmt.Fprintf(os.Stdout, "Response from `CapabilityAPI.CapabilitiesIdMembersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesIdMembersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CapabilityMembersApiResource**](CapabilityMembersApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesIdMembershipapplicationsGet

> MembershipApplicationListApiResource CapabilitiesIdMembershipapplicationsGet(ctx, id).Execute()



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
	resp, r, err := apiClient.CapabilityAPI.CapabilitiesIdMembershipapplicationsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdMembershipapplicationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesIdMembershipapplicationsGet`: MembershipApplicationListApiResource
	fmt.Fprintf(os.Stdout, "Response from `CapabilityAPI.CapabilitiesIdMembershipapplicationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesIdMembershipapplicationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MembershipApplicationListApiResource**](MembershipApplicationListApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesIdMembershipapplicationsPost

> MembershipApplicationApiResource CapabilitiesIdMembershipapplicationsPost(ctx, id).Execute()



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
	resp, r, err := apiClient.CapabilityAPI.CapabilitiesIdMembershipapplicationsPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdMembershipapplicationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesIdMembershipapplicationsPost`: MembershipApplicationApiResource
	fmt.Fprintf(os.Stdout, "Response from `CapabilityAPI.CapabilitiesIdMembershipapplicationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesIdMembershipapplicationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MembershipApplicationApiResource**](MembershipApplicationApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesIdMetadataGet

> CapabilitiesIdMetadataGet(ctx, id).Execute()



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
	r, err := apiClient.CapabilityAPI.CapabilitiesIdMetadataGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdMetadataGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCapabilitiesIdMetadataGetRequest struct via the builder pattern


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


## CapabilitiesIdMetadataPost

> CapabilitiesIdMetadataPost(ctx, id).SetCapabilityMetadataRequest(setCapabilityMetadataRequest).Execute()



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
	setCapabilityMetadataRequest := *openapiclient.NewSetCapabilityMetadataRequest(map[string]JsonNode{"key": *openapiclient.NewJsonNode()}) // SetCapabilityMetadataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CapabilityAPI.CapabilitiesIdMetadataPost(context.Background(), id).SetCapabilityMetadataRequest(setCapabilityMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdMetadataPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCapabilitiesIdMetadataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setCapabilityMetadataRequest** | [**SetCapabilityMetadataRequest**](SetCapabilityMetadataRequest.md) |  | 

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


## CapabilitiesIdRequestdeletionPost

> CapabilitiesIdRequestdeletionPost(ctx, id).User(user).Execute()



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
	user := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CapabilityAPI.CapabilitiesIdRequestdeletionPost(context.Background(), id).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdRequestdeletionPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCapabilitiesIdRequestdeletionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

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


## CapabilitiesIdRequiredMetadataPost

> CapabilitiesIdRequiredMetadataPost(ctx, id).SetCapabilityMetadataRequest(setCapabilityMetadataRequest).Execute()



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
	setCapabilityMetadataRequest := *openapiclient.NewSetCapabilityMetadataRequest(map[string]JsonNode{"key": *openapiclient.NewJsonNode()}) // SetCapabilityMetadataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CapabilityAPI.CapabilitiesIdRequiredMetadataPost(context.Background(), id).SetCapabilityMetadataRequest(setCapabilityMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdRequiredMetadataPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCapabilitiesIdRequiredMetadataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setCapabilityMetadataRequest** | [**SetCapabilityMetadataRequest**](SetCapabilityMetadataRequest.md) |  | 

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


## CapabilitiesIdTeamsGet

> CapabilitiesIdTeamsGet(ctx, id).Execute()



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
	r, err := apiClient.CapabilityAPI.CapabilitiesIdTeamsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdTeamsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCapabilitiesIdTeamsGetRequest struct via the builder pattern


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


## CapabilitiesIdTopicsPost

> KafkaTopicApiResource CapabilitiesIdTopicsPost(ctx, id).NewKafkaTopicRequest(newKafkaTopicRequest).Execute()



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
	newKafkaTopicRequest := *openapiclient.NewNewKafkaTopicRequest("KafkaClusterId_example", "Name_example", "Description_example", int32(123), "Retention_example") // NewKafkaTopicRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CapabilityAPI.CapabilitiesIdTopicsPost(context.Background(), id).NewKafkaTopicRequest(newKafkaTopicRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesIdTopicsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesIdTopicsPost`: KafkaTopicApiResource
	fmt.Fprintf(os.Stdout, "Response from `CapabilityAPI.CapabilitiesIdTopicsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesIdTopicsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newKafkaTopicRequest** | [**NewKafkaTopicRequest**](NewKafkaTopicRequest.md) |  | 

### Return type

[**KafkaTopicApiResource**](KafkaTopicApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CapabilitiesPost

> CapabilityDetailsApiResource CapabilitiesPost(ctx).NewCapabilityRequest(newCapabilityRequest).Execute()



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
	newCapabilityRequest := *openapiclient.NewNewCapabilityRequest("Name_example") // NewCapabilityRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CapabilityAPI.CapabilitiesPost(context.Background()).NewCapabilityRequest(newCapabilityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilityAPI.CapabilitiesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapabilitiesPost`: CapabilityDetailsApiResource
	fmt.Fprintf(os.Stdout, "Response from `CapabilityAPI.CapabilitiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCapabilitiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newCapabilityRequest** | [**NewCapabilityRequest**](NewCapabilityRequest.md) |  | 

### Return type

[**CapabilityDetailsApiResource**](CapabilityDetailsApiResource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

