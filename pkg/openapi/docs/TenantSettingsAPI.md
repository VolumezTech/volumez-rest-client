# \TenantSettingsAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTenantSettings**](TenantSettingsAPI.md#GetTenantSettings) | **Get** /tenant-settings | Get tenant settings
[**ModifyTenantSettings**](TenantSettingsAPI.md#ModifyTenantSettings) | **Put** /tenant-settings | Modify tenant settings
[**ResetTenantSettings**](TenantSettingsAPI.md#ResetTenantSettings) | **Patch** /tenant-settings/reset | Reset tenant settings



## GetTenantSettings

> GetTenantSettingsResponse GetTenantSettings(ctx).Authorization(authorization).Execute()

Get tenant settings

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
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantSettingsAPI.GetTenantSettings(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantSettingsAPI.GetTenantSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantSettings`: GetTenantSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantSettingsAPI.GetTenantSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**GetTenantSettingsResponse**](GetTenantSettingsResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTenantSettings

> map[string]interface{} ModifyTenantSettings(ctx).Authorization(authorization).ModifyTenantSettingsRequest(modifyTenantSettingsRequest).Execute()

Modify tenant settings

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
	authorization := "authorization_example" // string |  (optional)
	modifyTenantSettingsRequest := *openapiclient.NewModifyTenantSettingsRequest() // ModifyTenantSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantSettingsAPI.ModifyTenantSettings(context.Background()).Authorization(authorization).ModifyTenantSettingsRequest(modifyTenantSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantSettingsAPI.ModifyTenantSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTenantSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantSettingsAPI.ModifyTenantSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyTenantSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **modifyTenantSettingsRequest** | [**ModifyTenantSettingsRequest**](ModifyTenantSettingsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetTenantSettings

> map[string]interface{} ResetTenantSettings(ctx).Authorization(authorization).ResetTenantSettingsRequest(resetTenantSettingsRequest).Execute()

Reset tenant settings

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
	authorization := "authorization_example" // string |  (optional)
	resetTenantSettingsRequest := *openapiclient.NewResetTenantSettingsRequest(*openapiclient.NewModifyTenantSettingsRequestSettingsToModify()) // ResetTenantSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantSettingsAPI.ResetTenantSettings(context.Background()).Authorization(authorization).ResetTenantSettingsRequest(resetTenantSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantSettingsAPI.ResetTenantSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetTenantSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantSettingsAPI.ResetTenantSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetTenantSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **resetTenantSettingsRequest** | [**ResetTenantSettingsRequest**](ResetTenantSettingsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

