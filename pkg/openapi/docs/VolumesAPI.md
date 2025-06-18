# \VolumesAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchVolumesPlan**](VolumesAPI.md#BatchVolumesPlan) | **Post** /volumes/plan | check if volumes can be created
[**VolumeCreate**](VolumesAPI.md#VolumeCreate) | **Post** /volumes | Create a new volume
[**VolumeDelete**](VolumesAPI.md#VolumeDelete) | **Delete** /volumes/{volume} | Delete a volume
[**VolumeDescribe**](VolumesAPI.md#VolumeDescribe) | **Get** /volumes/{volume}/describe | describe existing volume 
[**VolumeGet**](VolumesAPI.md#VolumeGet) | **Get** /volumes/{volume} | Get the properties of a volume
[**VolumeModify**](VolumesAPI.md#VolumeModify) | **Patch** /volumes/{volume} | Modify a volume
[**VolumeRecoverInitiate**](VolumesAPI.md#VolumeRecoverInitiate) | **Post** /volumes/{volume}/recover | Initiate recover on volume
[**VolumesList**](VolumesAPI.md#VolumesList) | **Get** /volumes | Get a list of volumes



## BatchVolumesPlan

> VolumePlanOutput BatchVolumesPlan(ctx).Body(body).Verbose(verbose).Authorization(authorization).Execute()

check if volumes can be created

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
	body := *openapiclient.NewBatchVolumesPlanRequest() // BatchVolumesPlanRequest | 
	verbose := true // bool | if true will return the volume plan if false will omit the plan from the response (optional) (default to true)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.BatchVolumesPlan(context.Background()).Body(body).Verbose(verbose).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.BatchVolumesPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchVolumesPlan`: VolumePlanOutput
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.BatchVolumesPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchVolumesPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BatchVolumesPlanRequest**](BatchVolumesPlanRequest.md) |  | 
 **verbose** | **bool** | if true will return the volume plan if false will omit the plan from the response | [default to true]
 **authorization** | **string** |  | 

### Return type

[**VolumePlanOutput**](VolumePlanOutput.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumeCreate

> RegularResponse VolumeCreate(ctx).Body(body).Authorization(authorization).Execute()

Create a new volume

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
	body := *openapiclient.NewVolume("Name_example", "Type_example", int32(123), "Policy_example") // Volume | A Volume object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeCreate(context.Background()).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumeCreate`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVolumeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Volume**](Volume.md) | A Volume object | 
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


## VolumeDelete

> RegularResponse VolumeDelete(ctx, volume).Force(force).Authorization(authorization).Execute()

Delete a volume

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
	volume := "volume_example" // string | 
	force := true // bool |  (optional) (default to false)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeDelete(context.Background(), volume).Force(force).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumeDelete`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | [default to false]
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


## VolumeDescribe

> VolumeGroup VolumeDescribe(ctx, volume).Authorization(authorization).Execute()

describe existing volume 

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
	volume := "volume_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeDescribe(context.Background(), volume).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeDescribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumeDescribe`: VolumeGroup
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeDescribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeDescribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**VolumeGroup**](VolumeGroup.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumeGet

> Volume VolumeGet(ctx, volume).Authorization(authorization).Execute()

Get the properties of a volume

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
	volume := "volume_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeGet(context.Background(), volume).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumeGet`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumeModify

> RegularResponse VolumeModify(ctx, volume).Body(body).Authorization(authorization).Execute()

Modify a volume

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
	volume := "volume_example" // string | 
	body := *openapiclient.NewVolume("Name_example", "Type_example", int32(123), "Policy_example") // Volume | A Volume object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeModify(context.Background(), volume).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumeModify`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Volume**](Volume.md) | A Volume object | 
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


## VolumeRecoverInitiate

> RegularResponse VolumeRecoverInitiate(ctx, volume).Authorization(authorization).Execute()

Initiate recover on volume

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
	volume := "volume_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumeRecoverInitiate(context.Background(), volume).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumeRecoverInitiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumeRecoverInitiate`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumeRecoverInitiate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeRecoverInitiateRequest struct via the builder pattern


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


## VolumesList

> []Volume VolumesList(ctx).Capacity(capacity).Authorization(authorization).Execute()

Get a list of volumes

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
	capacity := true // bool |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumesList(context.Background()).Capacity(capacity).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesList`: []Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVolumesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **capacity** | **bool** |  | 
 **authorization** | **string** |  | 

### Return type

[**[]Volume**](Volume.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

