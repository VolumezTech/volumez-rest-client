# \ConnectivitiesAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectivitiesList**](ConnectivitiesAPI.md#ConnectivitiesList) | **Get** /connectivities | Get a list of connectivities
[**ConnectivityCreate**](ConnectivitiesAPI.md#ConnectivityCreate) | **Post** /connectivities | Create a new connectivity
[**ConnectivityDelete**](ConnectivitiesAPI.md#ConnectivityDelete) | **Delete** /connectivities/{connectivity} | Delete a connectivity
[**ConnectivityGet**](ConnectivitiesAPI.md#ConnectivityGet) | **Get** /connectivities/{connectivity} | Get the properties of a connectivity
[**ConnectivityModify**](ConnectivitiesAPI.md#ConnectivityModify) | **Patch** /connectivities/{connectivity} | Modify a connectivity



## ConnectivitiesList

> []Connectivity ConnectivitiesList(ctx).Authorization(authorization).Execute()

Get a list of connectivities

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
	resp, r, err := apiClient.ConnectivitiesAPI.ConnectivitiesList(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectivitiesAPI.ConnectivitiesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectivitiesList`: []Connectivity
	fmt.Fprintf(os.Stdout, "Response from `ConnectivitiesAPI.ConnectivitiesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectivitiesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**[]Connectivity**](Connectivity.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectivityCreate

> RegularResponse ConnectivityCreate(ctx).Body(body).Authorization(authorization).Execute()

Create a new connectivity

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
	body := *openapiclient.NewConnectivity("Name_example", "Zones1_example", "Systemtypes1_example", "Zones2_example", "Systemtypes2_example", "Mediaprotocol_example", "Replicationprotocol_example") // Connectivity | A Connectivity object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectivitiesAPI.ConnectivityCreate(context.Background()).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectivitiesAPI.ConnectivityCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectivityCreate`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectivitiesAPI.ConnectivityCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectivityCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Connectivity**](Connectivity.md) | A Connectivity object | 
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


## ConnectivityDelete

> RegularResponse ConnectivityDelete(ctx, connectivity).Authorization(authorization).Execute()

Delete a connectivity

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
	connectivity := "connectivity_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectivitiesAPI.ConnectivityDelete(context.Background(), connectivity).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectivitiesAPI.ConnectivityDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectivityDelete`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectivitiesAPI.ConnectivityDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectivity** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectivityDeleteRequest struct via the builder pattern


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


## ConnectivityGet

> Connectivity ConnectivityGet(ctx, connectivity).Authorization(authorization).Execute()

Get the properties of a connectivity

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
	connectivity := "connectivity_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectivitiesAPI.ConnectivityGet(context.Background(), connectivity).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectivitiesAPI.ConnectivityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectivityGet`: Connectivity
	fmt.Fprintf(os.Stdout, "Response from `ConnectivitiesAPI.ConnectivityGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectivity** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectivityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**Connectivity**](Connectivity.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectivityModify

> RegularResponse ConnectivityModify(ctx, connectivity).Body(body).Authorization(authorization).Execute()

Modify a connectivity

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
	connectivity := "connectivity_example" // string | 
	body := *openapiclient.NewConnectivity("Name_example", "Zones1_example", "Systemtypes1_example", "Zones2_example", "Systemtypes2_example", "Mediaprotocol_example", "Replicationprotocol_example") // Connectivity | A Connectivity object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectivitiesAPI.ConnectivityModify(context.Background(), connectivity).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectivitiesAPI.ConnectivityModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectivityModify`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectivitiesAPI.ConnectivityModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectivity** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectivityModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Connectivity**](Connectivity.md) | A Connectivity object | 
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

