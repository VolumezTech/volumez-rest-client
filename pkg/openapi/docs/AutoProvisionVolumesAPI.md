# \AutoProvisionVolumesAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoProvisionVolumes**](AutoProvisionVolumesAPI.md#AutoProvisionVolumes) | **Post** /autoprovisionvolumes | Create a new auto provisioned volume



## AutoProvisionVolumes

> RegularResponse AutoProvisionVolumes(ctx).Body(body).Authorization(authorization).Execute()

Create a new auto provisioned volume

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
	body := *openapiclient.NewAutoProvisionVolume(*openapiclient.NewVolume("Name_example", "Type_example", int32(123), "Policy_example"), "CloudProvider_example", "AccountId_example", "Region_example", []string{"AvailabilityZones_example"}, []string{"Subnets_example"}, "OsType_example") // AutoProvisionVolume | Auto Provision Volume object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoProvisionVolumesAPI.AutoProvisionVolumes(context.Background()).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoProvisionVolumesAPI.AutoProvisionVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoProvisionVolumes`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `AutoProvisionVolumesAPI.AutoProvisionVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoProvisionVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AutoProvisionVolume**](AutoProvisionVolume.md) | Auto Provision Volume object | 
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

