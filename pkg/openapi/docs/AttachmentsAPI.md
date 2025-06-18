# \AttachmentsAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachmentCreate**](AttachmentsAPI.md#AttachmentCreate) | **Post** /volumes/{volume}/snapshots/{snapshot}/attachments | Create a new attachment
[**AttachmentDelete**](AttachmentsAPI.md#AttachmentDelete) | **Delete** /volumes/{volume}/snapshots/{snapshot}/attachments/{node} | Delete an attachment
[**AttachmentGet**](AttachmentsAPI.md#AttachmentGet) | **Get** /volumes/{volume}/snapshots/{snapshot}/attachments/{node} | Get the properties of an attachment
[**AttachmentModify**](AttachmentsAPI.md#AttachmentModify) | **Patch** /volumes/{volume}/snapshots/{snapshot}/attachments/{node} | Modify an attachment
[**AttachmentsList**](AttachmentsAPI.md#AttachmentsList) | **Get** /volumes/{volume}/snapshots/{snapshot}/attachments | Get a list of attachments for a given volume and snapshot
[**AttachmentsListAll**](AttachmentsAPI.md#AttachmentsListAll) | **Get** /attachments | Get a list of all attachments
[**AttachmentsListForVolume**](AttachmentsAPI.md#AttachmentsListForVolume) | **Get** /volumes/{volume}/attachments | Get a list of attachments for a given volume



## AttachmentCreate

> RegularResponse AttachmentCreate(ctx, volume, snapshot).Body(body).Authorization(authorization).Execute()

Create a new attachment

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
	snapshot := "snapshot_example" // string | 
	body := *openapiclient.NewAttachment("Node_example") // Attachment | An Attachment object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentCreate(context.Background(), volume, snapshot).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentCreate`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 
**snapshot** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**Attachment**](Attachment.md) | An Attachment object | 
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


## AttachmentDelete

> RegularResponse AttachmentDelete(ctx, volume, snapshot, node).Force(force).Authorization(authorization).Execute()

Delete an attachment

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
	snapshot := "snapshot_example" // string | 
	node := "node_example" // string | 
	force := true // bool |  (optional) (default to false)
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentDelete(context.Background(), volume, snapshot, node).Force(force).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentDelete`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 
**snapshot** | **string** |  | 
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentDeleteRequest struct via the builder pattern


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


## AttachmentGet

> Attachment AttachmentGet(ctx, volume, snapshot, node).Authorization(authorization).Execute()

Get the properties of an attachment

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
	snapshot := "snapshot_example" // string | 
	node := "node_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentGet(context.Background(), volume, snapshot, node).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentGet`: Attachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 
**snapshot** | **string** |  | 
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

### Return type

[**Attachment**](Attachment.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentModify

> RegularResponse AttachmentModify(ctx, volume, snapshot, node).Body(body).Authorization(authorization).Execute()

Modify an attachment

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
	snapshot := "snapshot_example" // string | 
	node := "node_example" // string | 
	body := *openapiclient.NewAttachment("Node_example") // Attachment | An Attachment object
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentModify(context.Background(), volume, snapshot, node).Body(body).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentModify`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 
**snapshot** | **string** |  | 
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**Attachment**](Attachment.md) | An Attachment object | 
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


## AttachmentsList

> []Attachment AttachmentsList(ctx, volume, snapshot).Authorization(authorization).Execute()

Get a list of attachments for a given volume and snapshot

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
	snapshot := "snapshot_example" // string | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttachmentsAPI.AttachmentsList(context.Background(), volume, snapshot).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentsList`: []Attachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 
**snapshot** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 

### Return type

[**[]Attachment**](Attachment.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentsListAll

> []Attachment AttachmentsListAll(ctx).Authorization(authorization).Execute()

Get a list of all attachments

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
	resp, r, err := apiClient.AttachmentsAPI.AttachmentsListAll(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentsListAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentsListAll`: []Attachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentsListAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentsListAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**[]Attachment**](Attachment.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentsListForVolume

> []Attachment AttachmentsListForVolume(ctx, volume).Authorization(authorization).Execute()

Get a list of attachments for a given volume

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
	resp, r, err := apiClient.AttachmentsAPI.AttachmentsListForVolume(context.Background(), volume).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttachmentsAPI.AttachmentsListForVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachmentsListForVolume`: []Attachment
	fmt.Fprintf(os.Stdout, "Response from `AttachmentsAPI.AttachmentsListForVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volume** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentsListForVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**[]Attachment**](Attachment.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

