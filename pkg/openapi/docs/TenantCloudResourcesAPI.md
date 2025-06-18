# \TenantCloudResourcesAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](TenantCloudResourcesAPI.md#CreateRole) | **Post** /tenant-cloud-resources/role | Create tenant cloud role
[**DeleteRole**](TenantCloudResourcesAPI.md#DeleteRole) | **Delete** /tenant-cloud-resources/{cloudProviderAccountId}/role | Delete tenant role resource by the given cloudProviderAccountId query param
[**GetAllRoles**](TenantCloudResourcesAPI.md#GetAllRoles) | **Get** /tenant-cloud-resources/role | Get All Tenant roles
[**GetVMMetadataByRegion**](TenantCloudResourcesAPI.md#GetVMMetadataByRegion) | **Get** /tenant-cloud-resources/{cloudProviderAccountId}/vm/{region}/metadata/{nodeId} | Get tenant&#39;s cloud VM&#39; metadata by region
[**GetVMVPCs**](TenantCloudResourcesAPI.md#GetVMVPCs) | **Get** /tenant-cloud-resources/{cloudProviderAccountId}/vm/vpcs | Get all tenant&#39;s cloud VM&#39;s available VPCs
[**GetVMVPCsByRegion**](TenantCloudResourcesAPI.md#GetVMVPCsByRegion) | **Get** /tenant-cloud-resources/{cloudProviderAccountId}/vm/{region}/vpcs | Get tenant&#39;s cloud VM&#39;s available VPCs by region
[**GetVMZones**](TenantCloudResourcesAPI.md#GetVMZones) | **Get** /tenant-cloud-resources/vm/zones | Get cloud provider zones for requested region
[**UpdateRole**](TenantCloudResourcesAPI.md#UpdateRole) | **Put** /tenant-cloud-resources/{cloudProviderAccountId}/role | Update tenant cloud role for the given cloudProviderAccountId



## CreateRole

> map[string]interface{} CreateRole(ctx).CreateRoleRequest(createRoleRequest).Authorization(authorization).Execute()

Create tenant cloud role

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
	createRoleRequest := *openapiclient.NewCreateRoleRequest() // CreateRoleRequest | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantCloudResourcesAPI.CreateRole(context.Background()).CreateRoleRequest(createRoleRequest).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantCloudResourcesAPI.CreateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRole`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TenantCloudResourcesAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRoleRequest** | [**CreateRoleRequest**](CreateRoleRequest.md) |  | 
 **authorization** | **string** |  | 

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


## DeleteRole

> DeleteRole(ctx, cloudProviderAccountId).Authorization(authorization).Execute()

Delete tenant role resource by the given cloudProviderAccountId query param

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
	cloudProviderAccountId := "cloudProviderAccountId_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantCloudResourcesAPI.DeleteRole(context.Background(), cloudProviderAccountId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantCloudResourcesAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRoles

> GetRolesResponse GetAllRoles(ctx).Authorization(authorization).Execute()

Get All Tenant roles

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
	resp, r, err := apiClient.TenantCloudResourcesAPI.GetAllRoles(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantCloudResourcesAPI.GetAllRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllRoles`: GetRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantCloudResourcesAPI.GetAllRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**GetRolesResponse**](GetRolesResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMMetadataByRegion

> GetVMMetadataResponse GetVMMetadataByRegion(ctx, cloudProviderAccountId, region, nodeId).Authorization(authorization).Execute()

Get tenant's cloud VM' metadata by region

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
	cloudProviderAccountId := "cloudProviderAccountId_example" // string | 
	region := "region_example" // string | 
	nodeId := "nodeId_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantCloudResourcesAPI.GetVMMetadataByRegion(context.Background(), cloudProviderAccountId, region, nodeId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantCloudResourcesAPI.GetVMMetadataByRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMMetadataByRegion`: GetVMMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantCloudResourcesAPI.GetVMMetadataByRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderAccountId** | **string** |  | 
**region** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMMetadataByRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

### Return type

[**GetVMMetadataResponse**](GetVMMetadataResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMVPCs

> map[string][]AvailableVPCItem GetVMVPCs(ctx, cloudProviderAccountId).Authorization(authorization).Execute()

Get all tenant's cloud VM's available VPCs

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
	cloudProviderAccountId := "cloudProviderAccountId_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantCloudResourcesAPI.GetVMVPCs(context.Background(), cloudProviderAccountId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantCloudResourcesAPI.GetVMVPCs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMVPCs`: map[string][]AvailableVPCItem
	fmt.Fprintf(os.Stdout, "Response from `TenantCloudResourcesAPI.GetVMVPCs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMVPCsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**map[string][]AvailableVPCItem**](array.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMVPCsByRegion

> map[string][]AvailableVPCItem GetVMVPCsByRegion(ctx, cloudProviderAccountId, region).Authorization(authorization).Execute()

Get tenant's cloud VM's available VPCs by region

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
	cloudProviderAccountId := "cloudProviderAccountId_example" // string | 
	region := "region_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantCloudResourcesAPI.GetVMVPCsByRegion(context.Background(), cloudProviderAccountId, region).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantCloudResourcesAPI.GetVMVPCsByRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMVPCsByRegion`: map[string][]AvailableVPCItem
	fmt.Fprintf(os.Stdout, "Response from `TenantCloudResourcesAPI.GetVMVPCsByRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderAccountId** | **string** |  | 
**region** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMVPCsByRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 

### Return type

[**map[string][]AvailableVPCItem**](array.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMZones

> GetVmRegionZonesResponse GetVMZones(ctx).CloudProvider(cloudProvider).Authorization(authorization).Execute()

Get cloud provider zones for requested region

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
	cloudProvider := openapiclient.CloudProviderType("aws") // CloudProviderType | Cloud provider type (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantCloudResourcesAPI.GetVMZones(context.Background()).CloudProvider(cloudProvider).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantCloudResourcesAPI.GetVMZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVMZones`: GetVmRegionZonesResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantCloudResourcesAPI.GetVMZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVMZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProvider** | [**CloudProviderType**](CloudProviderType.md) | Cloud provider type | 
 **authorization** | **string** |  | 

### Return type

[**GetVmRegionZonesResponse**](GetVmRegionZonesResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> UpdateRole(ctx, cloudProviderAccountId).UpdateRoleRequest(updateRoleRequest).Authorization(authorization).Execute()

Update tenant cloud role for the given cloudProviderAccountId

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
	cloudProviderAccountId := "cloudProviderAccountId_example" // string | 
	updateRoleRequest := *openapiclient.NewUpdateRoleRequest() // UpdateRoleRequest | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantCloudResourcesAPI.UpdateRole(context.Background(), cloudProviderAccountId).UpdateRoleRequest(updateRoleRequest).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantCloudResourcesAPI.UpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProviderAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleRequest** | [**UpdateRoleRequest**](UpdateRoleRequest.md) |  | 
 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

