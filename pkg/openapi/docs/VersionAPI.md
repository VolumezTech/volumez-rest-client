# \VersionAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VersionGet**](VersionAPI.md#VersionGet) | **Get** /version | Get version of sio



## VersionGet

> VersionResponse VersionGet(ctx).Authorization(authorization).Execute()

Get version of sio

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
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionAPI.VersionGet(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionAPI.VersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionGet`: VersionResponse
	fmt.Fprintf(os.Stdout, "Response from `VersionAPI.VersionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVersionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**VersionResponse**](VersionResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

