# \CapacityGroupsAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CapacityGroupGet**](CapacityGroupsAPI.md#CapacityGroupGet) | **Get** /capacitygroups | Get available capacity groups



## CapacityGroupGet

> []CapacityGroup CapacityGroupGet(ctx).Authorization(authorization).Execute()

Get available capacity groups

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
	resp, r, err := apiClient.CapacityGroupsAPI.CapacityGroupGet(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapacityGroupsAPI.CapacityGroupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CapacityGroupGet`: []CapacityGroup
	fmt.Fprintf(os.Stdout, "Response from `CapacityGroupsAPI.CapacityGroupGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCapacityGroupGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**[]CapacityGroup**](CapacityGroup.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

