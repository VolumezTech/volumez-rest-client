# \SystemAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeaturesList**](SystemAPI.md#GetFeaturesList) | **Get** /system/features | 
[**GetMachineInfo**](SystemAPI.md#GetMachineInfo) | **Get** /system/machineinfo | get system info



## GetFeaturesList

> []string GetFeaturesList(ctx).Authorization(authorization).Execute()



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
	resp, r, err := apiClient.SystemAPI.GetFeaturesList(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetFeaturesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeaturesList`: []string
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetFeaturesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

**[]string**

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMachineInfo

> MachineInfo GetMachineInfo(ctx).Authorization(authorization).Execute()

get system info

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
	resp, r, err := apiClient.SystemAPI.GetMachineInfo(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetMachineInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMachineInfo`: MachineInfo
	fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetMachineInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**MachineInfo**](MachineInfo.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

