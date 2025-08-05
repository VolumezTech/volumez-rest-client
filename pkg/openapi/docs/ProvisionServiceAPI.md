# \ProvisionServiceAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Provision**](ProvisionServiceAPI.md#Provision) | **Post** /provision | Volumez connector provisioning



## Provision

> ProvisionResponse Provision(ctx).ProvisionRequest(provisionRequest).Authorization(authorization).Execute()

Volumez connector provisioning

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
	provisionRequest := *openapiclient.NewProvisionRequest() // ProvisionRequest | A provisioning object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisionServiceAPI.Provision(context.Background()).ProvisionRequest(provisionRequest).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisionServiceAPI.Provision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Provision`: ProvisionResponse
	fmt.Fprintf(os.Stdout, "Response from `ProvisionServiceAPI.Provision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisionRequest** | [**ProvisionRequest**](ProvisionRequest.md) | A provisioning object | 
 **authorization** | **string** |  | 

### Return type

[**ProvisionResponse**](ProvisionResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

