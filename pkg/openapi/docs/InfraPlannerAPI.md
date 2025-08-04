# \InfraPlannerAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInfraPlan**](InfraPlannerAPI.md#CreateInfraPlan) | **Post** /infra-planner/create-infra-plan | 
[**CreatePublicInfraPlan**](InfraPlannerAPI.md#CreatePublicInfraPlan) | **Post** /infra-planner/create-infra-plan/public | 
[**ProviderPricingInfo**](InfraPlannerAPI.md#ProviderPricingInfo) | **Post** /infra-planner/provider-pricing-info | 



## CreateInfraPlan

> CreateInfraPlanResponse CreateInfraPlan(ctx).CreateInfraPlanRequest(createInfraPlanRequest).Authorization(authorization).ShouldValidate(shouldValidate).Execute()



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
	createInfraPlanRequest := *openapiclient.NewCreateInfraPlanRequest(openapiclient.CloudProviderType("aws"), *openapiclient.NewCreateInfraPlanRequestPolicy(), int32(123)) // CreateInfraPlanRequest | 
	authorization := "authorization_example" // string |  (optional)
	shouldValidate := true // bool | validate the given resources (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfraPlannerAPI.CreateInfraPlan(context.Background()).CreateInfraPlanRequest(createInfraPlanRequest).Authorization(authorization).ShouldValidate(shouldValidate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraPlannerAPI.CreateInfraPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInfraPlan`: CreateInfraPlanResponse
	fmt.Fprintf(os.Stdout, "Response from `InfraPlannerAPI.CreateInfraPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInfraPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInfraPlanRequest** | [**CreateInfraPlanRequest**](CreateInfraPlanRequest.md) |  | 
 **authorization** | **string** |  | 
 **shouldValidate** | **bool** | validate the given resources | [default to true]

### Return type

[**CreateInfraPlanResponse**](CreateInfraPlanResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePublicInfraPlan

> CreateInfraPlanResponse CreatePublicInfraPlan(ctx).CreateInfraPlanRequest(createInfraPlanRequest).Execute()



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
	createInfraPlanRequest := *openapiclient.NewCreateInfraPlanRequest(openapiclient.CloudProviderType("aws"), *openapiclient.NewCreateInfraPlanRequestPolicy(), int32(123)) // CreateInfraPlanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfraPlannerAPI.CreatePublicInfraPlan(context.Background()).CreateInfraPlanRequest(createInfraPlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraPlannerAPI.CreatePublicInfraPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePublicInfraPlan`: CreateInfraPlanResponse
	fmt.Fprintf(os.Stdout, "Response from `InfraPlannerAPI.CreatePublicInfraPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicInfraPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInfraPlanRequest** | [**CreateInfraPlanRequest**](CreateInfraPlanRequest.md) |  | 

### Return type

[**CreateInfraPlanResponse**](CreateInfraPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProviderPricingInfo

> ProviderPricingInfoResponse ProviderPricingInfo(ctx).ProviderPricingInfoRequest(providerPricingInfoRequest).Execute()



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
	providerPricingInfoRequest := *openapiclient.NewProviderPricingInfoRequest(openapiclient.CloudProviderType("aws"), int32(123), int32(123)) // ProviderPricingInfoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfraPlannerAPI.ProviderPricingInfo(context.Background()).ProviderPricingInfoRequest(providerPricingInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfraPlannerAPI.ProviderPricingInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProviderPricingInfo`: ProviderPricingInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `InfraPlannerAPI.ProviderPricingInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProviderPricingInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerPricingInfoRequest** | [**ProviderPricingInfoRequest**](ProviderPricingInfoRequest.md) |  | 

### Return type

[**ProviderPricingInfoResponse**](ProviderPricingInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

