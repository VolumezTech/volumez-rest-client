# \ExportsAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportCreate**](ExportsAPI.md#ExportCreate) | **Post** /exports/ | Create export
[**ExportDelete**](ExportsAPI.md#ExportDelete) | **Delete** /exports/{export} | Delete an export
[**ExportModify**](ExportsAPI.md#ExportModify) | **Patch** /exports/{export} | Modify an export
[**ExportsList**](ExportsAPI.md#ExportsList) | **Get** /exports/ | Get a list of associations



## ExportCreate

> SuccessJobResponse ExportCreate(ctx).Body(body).Authorization(authorization).Execute()

Create export

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "bitbucket.org/volumez/volumez-rest-client"
)

func main() {
	body := *openapiclient.NewExportCreate() // ExportCreate | An export object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportCreate(context.Background()).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCreate`: SuccessJobResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExportCreate**](ExportCreate.md) | An export object | 
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


## ExportDelete

> SuccessJobResponse ExportDelete(ctx, export).Force(force).Authorization(authorization).Execute()

Delete an export

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "bitbucket.org/volumez/volumez-rest-client"
)

func main() {
	export := "export_example" // string | 
	force := true // bool |  (optional) (default to false)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportDelete(context.Background(), export).Force(force).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDelete`: SuccessJobResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**export** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | [default to false]
 **authorization** | **string** |  | 

### Return type

[**SuccessJobResponse**](SuccessJobResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportModify

> SuccessJobResponse ExportModify(ctx, export).Body(body).Authorization(authorization).Execute()

Modify an export

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "bitbucket.org/volumez/volumez-rest-client"
)

func main() {
	export := "export_example" // string | 
	body := *openapiclient.NewExportModify() // ExportModify | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportModify(context.Background(), export).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportModify`: SuccessJobResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**export** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ExportModify**](ExportModify.md) |  | 
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


## ExportsList

> []Export ExportsList(ctx).Startfrom(startfrom).Count(count).Authorization(authorization).Execute()

Get a list of associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "bitbucket.org/volumez/volumez-rest-client"
)

func main() {
	startfrom := "startfrom_example" // string |  (optional)
	count := int32(56) // int32 |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsList(context.Background()).Startfrom(startfrom).Count(count).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsList`: []Export
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startfrom** | **string** |  | 
 **count** | **int32** |  | 
 **authorization** | **string** |  | 

### Return type

[**[]Export**](Export.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

