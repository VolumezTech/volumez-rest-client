# \AssociationsAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociationCreate**](AssociationsAPI.md#AssociationCreate) | **Post** /associations | Create association
[**AssociationDelete**](AssociationsAPI.md#AssociationDelete) | **Delete** /associations/{association} | Delete an association
[**AssociationModify**](AssociationsAPI.md#AssociationModify) | **Patch** /associations/{association} | Association modify
[**AssociationsList**](AssociationsAPI.md#AssociationsList) | **Get** /associations | Get a list of associations



## AssociationCreate

> SuccessJobResponse AssociationCreate(ctx).Body(body).Authorization(authorization).Execute()

Create association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "bitbucket.org/volumez/volumez-openapi-client"
)

func main() {
	body := *openapiclient.NewAssociationCreate() // AssociationCreate | An association object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssociationsAPI.AssociationCreate(context.Background()).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociationsAPI.AssociationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociationCreate`: SuccessJobResponse
	fmt.Fprintf(os.Stdout, "Response from `AssociationsAPI.AssociationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssociationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AssociationCreate**](AssociationCreate.md) | An association object | 
 **authorization** | **string** |  | 

### Return type

[**SuccessJobResponse**](SuccessJobResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociationDelete

> RegularResponse AssociationDelete(ctx, association).Authorization(authorization).Execute()

Delete an association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "bitbucket.org/volumez/volumez-openapi-client"
)

func main() {
	association := "association_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssociationsAPI.AssociationDelete(context.Background(), association).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociationsAPI.AssociationDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociationDelete`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `AssociationsAPI.AssociationDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**association** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociationDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**RegularResponse**](RegularResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociationModify

> RegularResponse AssociationModify(ctx, association).Body(body).Authorization(authorization).Execute()

Association modify

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "bitbucket.org/volumez/volumez-openapi-client"
)

func main() {
	association := "association_example" // string | 
	body := *openapiclient.NewAssociationModify() // AssociationModify | An association object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssociationsAPI.AssociationModify(context.Background(), association).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociationsAPI.AssociationModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociationModify`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `AssociationsAPI.AssociationModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**association** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociationModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AssociationModify**](AssociationModify.md) | An association object | 
 **authorization** | **string** |  | 

### Return type

[**RegularResponse**](RegularResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociationsList

> []Association AssociationsList(ctx).Startfrom(startfrom).Count(count).Authorization(authorization).Execute()

Get a list of associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "bitbucket.org/volumez/volumez-openapi-client"
)

func main() {
	startfrom := "startfrom_example" // string |  (optional)
	count := int32(56) // int32 |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssociationsAPI.AssociationsList(context.Background()).Startfrom(startfrom).Count(count).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssociationsAPI.AssociationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociationsList`: []Association
	fmt.Fprintf(os.Stdout, "Response from `AssociationsAPI.AssociationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssociationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startfrom** | **string** |  | 
 **count** | **int32** |  | 
 **authorization** | **string** |  | 

### Return type

[**[]Association**](Association.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

