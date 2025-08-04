# \SnapshotsAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsistencyGroupGet**](SnapshotsAPI.md#ConsistencyGroupGet) | **Get** /volumes/snapshot/{snapshot_group_name} | List of snapshots group
[**ConsistencyGroupSnapshotCreate**](SnapshotsAPI.md#ConsistencyGroupSnapshotCreate) | **Post** /volumes/snapshot | Create a new snapshot for given consistency group
[**SnapshotCreate**](SnapshotsAPI.md#SnapshotCreate) | **Post** /volumes/{volume}/snapshots | Create a new snapshot
[**SnapshotDelete**](SnapshotsAPI.md#SnapshotDelete) | **Delete** /volumes/{volume}/snapshots/{snapshot} | Delete a snapshot
[**SnapshotGet**](SnapshotsAPI.md#SnapshotGet) | **Get** /volumes/{volume}/snapshots/{snapshot} | Get the properties of a snapshot
[**SnapshotModify**](SnapshotsAPI.md#SnapshotModify) | **Patch** /volumes/{volume}/snapshots/{snapshot} | Modify a snapshot
[**SnapshotRollback**](SnapshotsAPI.md#SnapshotRollback) | **Patch** /volumes/{volume}/snapshots/{snapshot}/rollback | Roll back to snapshot
[**SnapshotsList**](SnapshotsAPI.md#SnapshotsList) | **Get** /volumes/{volume}/snapshots | Get a list of snapshots
[**SnapshotsListAll**](SnapshotsAPI.md#SnapshotsListAll) | **Get** /snapshots | Get a list of all snapshots



## ConsistencyGroupGet

> []Snapshot ConsistencyGroupGet(ctx, snapshotGroupName).Authorization(authorization).Execute()

List of snapshots group

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
	snapshotGroupName := "snapshotGroupName_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.ConsistencyGroupGet(context.Background(), snapshotGroupName).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.ConsistencyGroupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsistencyGroupGet`: []Snapshot
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.ConsistencyGroupGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotGroupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsistencyGroupGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**[]Snapshot**](Snapshot.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsistencyGroupSnapshotCreate

> RegularResponse ConsistencyGroupSnapshotCreate(ctx).Body(body).Authorization(authorization).Execute()

Create a new snapshot for given consistency group

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
	body := *openapiclient.NewConsistencyGroupSnapshotCreateRequest() // ConsistencyGroupSnapshotCreateRequest | A Snapshot object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.ConsistencyGroupSnapshotCreate(context.Background()).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.ConsistencyGroupSnapshotCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsistencyGroupSnapshotCreate`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.ConsistencyGroupSnapshotCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsistencyGroupSnapshotCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConsistencyGroupSnapshotCreateRequest**](ConsistencyGroupSnapshotCreateRequest.md) | A Snapshot object | 
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


## SnapshotCreate

> RegularResponse SnapshotCreate(ctx, volume).Body(body).Authorization(authorization).Execute()

Create a new snapshot

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
	volume := "volume_example" // string | 
	body := *openapiclient.NewSnapshot("Name_example", "Consistency_example") // Snapshot | A Snapshot object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.SnapshotCreate(context.Background(), volume).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.SnapshotCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnapshotCreate`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.SnapshotCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Snapshot**](Snapshot.md) | A Snapshot object | 
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


## SnapshotDelete

> RegularResponse SnapshotDelete(ctx, volume, snapshot).Force(force).Authorization(authorization).Execute()

Delete a snapshot

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
	volume := "volume_example" // string | 
	snapshot := "snapshot_example" // string | 
	force := true // bool |  (optional) (default to false)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.SnapshotDelete(context.Background(), volume, snapshot).Force(force).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.SnapshotDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnapshotDelete`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.SnapshotDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 
**snapshot** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotDeleteRequest struct via the builder pattern


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


## SnapshotGet

> Snapshot SnapshotGet(ctx, volume, snapshot).Authorization(authorization).Execute()

Get the properties of a snapshot

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
	volume := "volume_example" // string | 
	snapshot := "snapshot_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.SnapshotGet(context.Background(), volume, snapshot).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.SnapshotGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnapshotGet`: Snapshot
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.SnapshotGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 
**snapshot** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotModify

> RegularResponse SnapshotModify(ctx, volume, snapshot).Body(body).Authorization(authorization).Execute()

Modify a snapshot

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
	volume := "volume_example" // string | 
	snapshot := "snapshot_example" // string | 
	body := *openapiclient.NewSnapshot("Name_example", "Consistency_example") // Snapshot | A Snapshot object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.SnapshotModify(context.Background(), volume, snapshot).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.SnapshotModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnapshotModify`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.SnapshotModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 
**snapshot** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Snapshot**](Snapshot.md) | A Snapshot object | 
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


## SnapshotRollback

> RegularResponse SnapshotRollback(ctx, volume, snapshot).Authorization(authorization).Execute()

Roll back to snapshot

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
	volume := "volume_example" // string | 
	snapshot := "snapshot_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.SnapshotRollback(context.Background(), volume, snapshot).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.SnapshotRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnapshotRollback`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.SnapshotRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 
**snapshot** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotRollbackRequest struct via the builder pattern


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


## SnapshotsList

> []Snapshot SnapshotsList(ctx, volume).Authorization(authorization).Execute()

Get a list of snapshots

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
	volume := "volume_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.SnapshotsList(context.Background(), volume).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.SnapshotsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnapshotsList`: []Snapshot
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.SnapshotsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**[]Snapshot**](Snapshot.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnapshotsListAll

> []Snapshot SnapshotsListAll(ctx).Capacity(capacity).Authorization(authorization).Execute()

Get a list of all snapshots

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
	capacity := true // bool |  (optional)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotsAPI.SnapshotsListAll(context.Background()).Capacity(capacity).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.SnapshotsListAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnapshotsListAll`: []Snapshot
	fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.SnapshotsListAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSnapshotsListAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **capacity** | **bool** |  | 
 **authorization** | **string** |  | 

### Return type

[**[]Snapshot**](Snapshot.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

