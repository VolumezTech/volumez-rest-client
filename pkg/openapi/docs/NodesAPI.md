# \NodesAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NodeDelete**](NodesAPI.md#NodeDelete) | **Delete** /nodes/{node} | Delete a node
[**NodeDescribe**](NodesAPI.md#NodeDescribe) | **Get** /nodes/{node}/describe | Node describe
[**NodeDrain**](NodesAPI.md#NodeDrain) | **Post** /nodes/{node}/drain | Node drain
[**NodeGet**](NodesAPI.md#NodeGet) | **Get** /nodes/{node} | Get the properties of a node
[**NodeHwScan**](NodesAPI.md#NodeHwScan) | **Post** /nodes/{node}/hw | Node hardware scan
[**NodeSetTags**](NodesAPI.md#NodeSetTags) | **Patch** /nodes/tags/{node} | Set the tags of a node
[**NodeUpgrade**](NodesAPI.md#NodeUpgrade) | **Post** /nodes/upgrade/{node} | performing node version upgrade
[**NodesList**](NodesAPI.md#NodesList) | **Get** /nodes | Get a list of nodes



## NodeDelete

> RegularResponse NodeDelete(ctx, node).Force(force).DelayDelete(delayDelete).Authorization(authorization).Execute()

Delete a node

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
	node := "node_example" // string |  node to delete
	force := true // bool |  (optional) (default to false)
	delayDelete := true // bool |  (optional) (default to false)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.NodeDelete(context.Background(), node).Force(force).DelayDelete(delayDelete).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.NodeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeDelete`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.NodeDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  node to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | [default to false]
 **delayDelete** | **bool** |  | [default to false]
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


## NodeDescribe

> NodeDescribeResponse NodeDescribe(ctx, node).Authorization(authorization).Execute()

Node describe

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
	node := "node_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.NodeDescribe(context.Background(), node).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.NodeDescribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeDescribe`: NodeDescribeResponse
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.NodeDescribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeDescribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**NodeDescribeResponse**](NodeDescribeResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NodeDrain

> RegularResponse NodeDrain(ctx, node).Cleanup(cleanup).Authorization(authorization).Execute()

Node drain

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
	node := "node_example" // string | 
	cleanup := true // bool |  (optional) (default to false)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.NodeDrain(context.Background(), node).Cleanup(cleanup).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.NodeDrain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeDrain`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.NodeDrain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeDrainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cleanup** | **bool** |  | [default to false]
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


## NodeGet

> Node NodeGet(ctx, node).Authorization(authorization).Execute()

Get the properties of a node

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
	node := "node_example" // string | Name of node to return
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.NodeGet(context.Background(), node).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.NodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeGet`: Node
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.NodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | Name of node to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**Node**](Node.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NodeHwScan

> SuccessJobResponse NodeHwScan(ctx, node).Properties(properties).Authorization(authorization).Execute()

Node hardware scan

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
	node := "node_example" // string | 
	properties := []string{"Inner_example"} // []string |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.NodeHwScan(context.Background(), node).Properties(properties).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.NodeHwScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeHwScan`: SuccessJobResponse
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.NodeHwScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeHwScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **properties** | **[]string** |  | 
 **authorization** | **string** |  | 

### Return type

[**SuccessJobResponse**](SuccessJobResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NodeSetTags

> RegularResponse NodeSetTags(ctx, node).Tags(tags).Authorization(authorization).Execute()

Set the tags of a node

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
	node := "node_example" // string | Name of node to return
	tags := map[string]string{"key": "Inner_example"} // map[string]string | user defined tags
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.NodeSetTags(context.Background(), node).Tags(tags).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.NodeSetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeSetTags`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.NodeSetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | Name of node to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeSetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tags** | **map[string]string** | user defined tags | 
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


## NodeUpgrade

> RegularResponse NodeUpgrade(ctx, node).Authorization(authorization).Body(body).Execute()

performing node version upgrade

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
	node := "node_example" // string | Name of node to upgrade
	authorization := "authorization_example" // string |  (optional)
	body := *openapiclient.NewNodeVersion() // NodeVersion | Connector Version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodesAPI.NodeUpgrade(context.Background(), node).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.NodeUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeUpgrade`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.NodeUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | Name of node to upgrade | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **body** | [**NodeVersion**](NodeVersion.md) | Connector Version | 

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


## NodesList

> []Node NodesList(ctx).Authorization(authorization).Execute()

Get a list of nodes

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
	resp, r, err := apiClient.NodesAPI.NodesList(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.NodesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodesList`: []Node
	fmt.Fprintf(os.Stdout, "Response from `NodesAPI.NodesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNodesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**[]Node**](Node.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

