# \UserAPI

All URIs are relative to *https://api.dev.volumez.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUser**](UserAPI.md#AddUser) | **Post** /add-user | Add user to tenant group users
[**ApproveChangePassword**](UserAPI.md#ApproveChangePassword) | **Post** /tenant/user/changepassword | Change client password
[**ChangePasswordLoggedIn**](UserAPI.md#ChangePasswordLoggedIn) | **Post** /tenant/user/changepasswordloggedin | Change clients password when user is logged in
[**DeleteAzureSSOMapping**](UserAPI.md#DeleteAzureSSOMapping) | **Delete** /sso/azure/mapping | Delete Azure SSO Mapping
[**DeleteTenantHost**](UserAPI.md#DeleteTenantHost) | **Delete** /tenant/{tenantID}/tenanthosts/{tenantHost} | 
[**DisableUser**](UserAPI.md#DisableUser) | **Post** /disable-user/{email} | disable user
[**EnableUser**](UserAPI.md#EnableUser) | **Post** /enable-user/{email} | enable user
[**GetAzureSSOMapping**](UserAPI.md#GetAzureSSOMapping) | **Get** /sso/azure/mapping | Get Azure SSO Mapping
[**GetSubscription**](UserAPI.md#GetSubscription) | **Get** /azuremarketplace/subscription | get subscription
[**GetTenantAccessToken**](UserAPI.md#GetTenantAccessToken) | **Get** /tenant/token | 
[**GetTenantHost**](UserAPI.md#GetTenantHost) | **Get** /tenant/tenanthost | 
[**GetTenantIDFromTenantToken**](UserAPI.md#GetTenantIDFromTenantToken) | **Get** /tenant/tenantid | 
[**GetTenantRefreshToken**](UserAPI.md#GetTenantRefreshToken) | **Post** /tenant/refreshtoken | 
[**GetTenantUser**](UserAPI.md#GetTenantUser) | **Get** /tenant/user | Get user details
[**HandleSSOCallback**](UserAPI.md#HandleSSOCallback) | **Get** /sso/callback/{userPoolID}/{applicationClientId} | Handle SSO callback
[**InviteUser**](UserAPI.md#InviteUser) | **Post** /invite-user/{email} | invite a user to join tenant (send add user signup email
[**PutAzureSSOMapping**](UserAPI.md#PutAzureSSOMapping) | **Put** /sso/azure/mapping | Put Azure SSO Mapping
[**RefreshTenantAPIAccessCredentials**](UserAPI.md#RefreshTenantAPIAccessCredentials) | **Get** /tenant/apiaccess/credentials/refresh | 
[**RequestChangePassword**](UserAPI.md#RequestChangePassword) | **Post** /tenant/user/requestchangepassword | Request change client password
[**SignIn**](UserAPI.md#SignIn) | **Post** /signin | User SignIn
[**SignOut**](UserAPI.md#SignOut) | **Post** /signout | Signs out user from all devices
[**SignUp**](UserAPI.md#SignUp) | **Post** /signup | Create Tenant&#39;s first user
[**TenantHostAccessCredentials**](UserAPI.md#TenantHostAccessCredentials) | **Get** /tenant/tenanthost/credentials | 
[**TenantUserCreate**](UserAPI.md#TenantUserCreate) | **Post** /tenant/user | Create Tenant&#39;s additional user
[**TenantUsers**](UserAPI.md#TenantUsers) | **Get** /users/{tenantId} | List all tenant group users
[**UserConfirm**](UserAPI.md#UserConfirm) | **Get** /tenant/user/confirmation | Confirm user signup



## AddUser

> RegularResponse AddUser(ctx).AddUserRequest(addUserRequest).Execute()

Add user to tenant group users

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
	addUserRequest := *openapiclient.NewAddUserRequest("Email_example", "Password_example", "Name_example", "InvitedUserToken_example") // AddUserRequest | New user attributes to add in Cognito

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.AddUser(context.Background()).AddUserRequest(addUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AddUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUser`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUserRequest** | [**AddUserRequest**](AddUserRequest.md) | New user attributes to add in Cognito | 

### Return type

[**RegularResponse**](RegularResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApproveChangePassword

> RegularResponse ApproveChangePassword(ctx).ChangePasswordApproveRequest(changePasswordApproveRequest).Execute()

Change client password

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
	changePasswordApproveRequest := *openapiclient.NewChangePasswordApproveRequest() // ChangePasswordApproveRequest | new user password

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.ApproveChangePassword(context.Background()).ChangePasswordApproveRequest(changePasswordApproveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ApproveChangePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveChangePassword`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ApproveChangePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApproveChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changePasswordApproveRequest** | [**ChangePasswordApproveRequest**](ChangePasswordApproveRequest.md) | new user password | 

### Return type

[**RegularResponse**](RegularResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangePasswordLoggedIn

> RegularResponse ChangePasswordLoggedIn(ctx).ChangePasswordRequestLoggedIn(changePasswordRequestLoggedIn).Authorization(authorization).Execute()

Change clients password when user is logged in

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
	changePasswordRequestLoggedIn := *openapiclient.NewChangePasswordRequestLoggedIn() // ChangePasswordRequestLoggedIn | new user password when logged in
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.ChangePasswordLoggedIn(context.Background()).ChangePasswordRequestLoggedIn(changePasswordRequestLoggedIn).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ChangePasswordLoggedIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangePasswordLoggedIn`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ChangePasswordLoggedIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordLoggedInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changePasswordRequestLoggedIn** | [**ChangePasswordRequestLoggedIn**](ChangePasswordRequestLoggedIn.md) | new user password when logged in | 
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


## DeleteAzureSSOMapping

> DeleteAzureSSOMapping(ctx).Authorization(authorization).Execute()

Delete Azure SSO Mapping

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
	r, err := apiClient.UserAPI.DeleteAzureSSOMapping(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteAzureSSOMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureSSOMappingRequest struct via the builder pattern


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


## DeleteTenantHost

> TenantHostDeleteResponse DeleteTenantHost(ctx, tenantHost, tenantID).Authorization(authorization).Execute()



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
	authorization := "authorization_example" // string | 
	tenantHost := "tenantHost_example" // string | 
	tenantID := "tenantID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.DeleteTenantHost(context.Background(), tenantHost, tenantID).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteTenantHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantHost`: TenantHostDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.DeleteTenantHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHost** | **string** |  | 
**tenantID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 



### Return type

[**TenantHostDeleteResponse**](TenantHostDeleteResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableUser

> RegularResponse DisableUser(ctx, email).Authorization(authorization).Execute()

disable user

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
	email := "email_example" // string | User email to disable
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.DisableUser(context.Background(), email).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DisableUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableUser`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.DisableUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | User email to disable | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableUserRequest struct via the builder pattern


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


## EnableUser

> RegularResponse EnableUser(ctx, email).Authorization(authorization).Execute()

enable user

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
	email := "email_example" // string | User email to enable
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.EnableUser(context.Background(), email).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.EnableUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableUser`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.EnableUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | User email to enable | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableUserRequest struct via the builder pattern


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


## GetAzureSSOMapping

> GetAzureSSOMappingResponse GetAzureSSOMapping(ctx).Authorization(authorization).Execute()

Get Azure SSO Mapping

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
	resp, r, err := apiClient.UserAPI.GetAzureSSOMapping(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetAzureSSOMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAzureSSOMapping`: GetAzureSSOMappingResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetAzureSSOMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureSSOMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**GetAzureSSOMappingResponse**](GetAzureSSOMappingResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> RegularResponse GetSubscription(ctx).Token(token).Execute()

get subscription

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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetSubscription(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscription`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** |  | 

### Return type

[**RegularResponse**](RegularResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantAccessToken

> TenantTokenResponse GetTenantAccessToken(ctx).Authorization(authorization).Execute()



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
	authorization := "authorization_example" // string | Tenant token authorization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetTenantAccessToken(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTenantAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantAccessToken`: TenantTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTenantAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Tenant token authorization | 

### Return type

[**TenantTokenResponse**](TenantTokenResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantHost

> GetTenantHostResponse GetTenantHost(ctx).Tenanthosttoken(tenanthosttoken).Tenantaccesstoken(tenantaccesstoken).Execute()



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
	tenanthosttoken := "tenanthosttoken_example" // string | 
	tenantaccesstoken := "tenantaccesstoken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetTenantHost(context.Background()).Tenanthosttoken(tenanthosttoken).Tenantaccesstoken(tenantaccesstoken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTenantHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantHost`: GetTenantHostResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTenantHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenanthosttoken** | **string** |  | 
 **tenantaccesstoken** | **string** |  | 

### Return type

[**GetTenantHostResponse**](GetTenantHostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantIDFromTenantToken

> GetTenantIDResponse GetTenantIDFromTenantToken(ctx).Tenanttoken(tenanttoken).Execute()



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
	tenanttoken := "tenanttoken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetTenantIDFromTenantToken(context.Background()).Tenanttoken(tenanttoken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTenantIDFromTenantToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantIDFromTenantToken`: GetTenantIDResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTenantIDFromTenantToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantIDFromTenantTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenanttoken** | **string** |  | 

### Return type

[**GetTenantIDResponse**](GetTenantIDResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantRefreshToken

> RefreshToken GetTenantRefreshToken(ctx).GetTenantRefreshTokenRequest(getTenantRefreshTokenRequest).Execute()



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
	getTenantRefreshTokenRequest := *openapiclient.NewGetTenantRefreshTokenRequest("Accesstoken_example", "Hostname_example") // GetTenantRefreshTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetTenantRefreshToken(context.Background()).GetTenantRefreshTokenRequest(getTenantRefreshTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTenantRefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantRefreshToken`: RefreshToken
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTenantRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getTenantRefreshTokenRequest** | [**GetTenantRefreshTokenRequest**](GetTenantRefreshTokenRequest.md) |  | 

### Return type

[**RefreshToken**](RefreshToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantUser

> GetTenantUserResponse GetTenantUser(ctx).Authorization(authorization).Execute()

Get user details

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
	resp, r, err := apiClient.UserAPI.GetTenantUser(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetTenantUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantUser`: GetTenantUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetTenantUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 

### Return type

[**GetTenantUserResponse**](GetTenantUserResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSSOCallback

> HandleSSOCallback(ctx, userPoolID, applicationClientId).Code(code).Execute()

Handle SSO callback



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
	userPoolID := "userPoolID_example" // string | User Pool ID
	applicationClientId := "applicationClientId_example" // string | Application Client ID
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.HandleSSOCallback(context.Background(), userPoolID, applicationClientId).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.HandleSSOCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userPoolID** | **string** | User Pool ID | 
**applicationClientId** | **string** | Application Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleSSOCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **code** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUser

> InviteUserResponse InviteUser(ctx, email).Authorization(authorization).Execute()

invite a user to join tenant (send add user signup email

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
	email := "email_example" // string | User email to invite
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.InviteUser(context.Background(), email).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.InviteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteUser`: InviteUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.InviteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | User email to invite | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**InviteUserResponse**](InviteUserResponse.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAzureSSOMapping

> map[string]interface{} PutAzureSSOMapping(ctx).PutAzureSSOMappingRequest(putAzureSSOMappingRequest).Authorization(authorization).Execute()

Put Azure SSO Mapping

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
	putAzureSSOMappingRequest := *openapiclient.NewPutAzureSSOMappingRequest("SecurityGroupID_example") // PutAzureSSOMappingRequest | Put Azure SSO Mapping for tenant
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PutAzureSSOMapping(context.Background()).PutAzureSSOMappingRequest(putAzureSSOMappingRequest).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PutAzureSSOMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAzureSSOMapping`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PutAzureSSOMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAzureSSOMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putAzureSSOMappingRequest** | [**PutAzureSSOMappingRequest**](PutAzureSSOMappingRequest.md) | Put Azure SSO Mapping for tenant | 
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


## RefreshTenantAPIAccessCredentials

> RefreshTokenResponse RefreshTenantAPIAccessCredentials(ctx).Refreshtoken(refreshtoken).Execute()



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
	refreshtoken := "refreshtoken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RefreshTenantAPIAccessCredentials(context.Background()).Refreshtoken(refreshtoken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RefreshTenantAPIAccessCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshTenantAPIAccessCredentials`: RefreshTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RefreshTenantAPIAccessCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTenantAPIAccessCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshtoken** | **string** |  | 

### Return type

[**RefreshTokenResponse**](RefreshTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestChangePassword

> RegularResponse RequestChangePassword(ctx).RequestChangePassword(requestChangePassword).Execute()

Request change client password

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
	requestChangePassword := *openapiclient.NewRequestChangePassword() // RequestChangePassword | new user password

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RequestChangePassword(context.Background()).RequestChangePassword(requestChangePassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RequestChangePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestChangePassword`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RequestChangePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestChangePassword** | [**RequestChangePassword**](RequestChangePassword.md) | new user password | 

### Return type

[**RegularResponse**](RegularResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignIn

> SignInResponse SignIn(ctx).SignInRequest(signInRequest).Execute()

User SignIn

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
	signInRequest := *openapiclient.NewSignInRequest("Email_example", "Password_example") // SignInRequest | A signin object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.SignIn(context.Background()).SignInRequest(signInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.SignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignIn`: SignInResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.SignIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signInRequest** | [**SignInRequest**](SignInRequest.md) | A signin object | 

### Return type

[**SignInResponse**](SignInResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignOut

> RegularResponse SignOut(ctx).SignOutRequest(signOutRequest).Execute()

Signs out user from all devices

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
	signOutRequest := *openapiclient.NewSignOutRequest("AccessToken_example") // SignOutRequest | Access Token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.SignOut(context.Background()).SignOutRequest(signOutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.SignOut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignOut`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.SignOut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signOutRequest** | [**SignOutRequest**](SignOutRequest.md) | Access Token | 

### Return type

[**RegularResponse**](RegularResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignUp

> SignUpResponse SignUp(ctx).SignUpRequest(signUpRequest).Execute()

Create Tenant's first user

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
	signUpRequest := *openapiclient.NewSignUpRequest("Email_example", "Password_example", "Name_example") // SignUpRequest | A user signup object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.SignUp(context.Background()).SignUpRequest(signUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.SignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignUp`: SignUpResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.SignUp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signUpRequest** | [**SignUpRequest**](SignUpRequest.md) | A user signup object | 

### Return type

[**SignUpResponse**](SignUpResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantHostAccessCredentials

> RefreshTokenResponse TenantHostAccessCredentials(ctx).Refreshtoken(refreshtoken).Tenantaccesstoken(tenantaccesstoken).Execute()



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
	refreshtoken := "refreshtoken_example" // string | 
	tenantaccesstoken := "tenantaccesstoken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.TenantHostAccessCredentials(context.Background()).Refreshtoken(refreshtoken).Tenantaccesstoken(tenantaccesstoken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.TenantHostAccessCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantHostAccessCredentials`: RefreshTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.TenantHostAccessCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantHostAccessCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshtoken** | **string** |  | 
 **tenantaccesstoken** | **string** |  | 

### Return type

[**RefreshTokenResponse**](RefreshTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantUserCreate

> RegularResponse TenantUserCreate(ctx).Authorization(authorization).CreateAddTenantUserRequest(createAddTenantUserRequest).Execute()

Create Tenant's additional user

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
	createAddTenantUserRequest := *openapiclient.NewCreateAddTenantUserRequest("Email_example", "Password_example", "Name_example", "TenantId_example") // CreateAddTenantUserRequest | new tenant's user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.TenantUserCreate(context.Background()).Authorization(authorization).CreateAddTenantUserRequest(createAddTenantUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.TenantUserCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantUserCreate`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.TenantUserCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantUserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **createAddTenantUserRequest** | [**CreateAddTenantUserRequest**](CreateAddTenantUserRequest.md) | new tenant&#39;s user | 

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


## TenantUsers

> []TenantUser TenantUsers(ctx, tenantId).Authorization(authorization).Execute()

List all tenant group users

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
	tenantId := "tenantId_example" // string | Tenant Id to get all users for
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.TenantUsers(context.Background(), tenantId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.TenantUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantUsers`: []TenantUser
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.TenantUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant Id to get all users for | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**[]TenantUser**](TenantUser.md)

### Authorization

[storage.io-authorizer](../README.md#storage.io-authorizer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserConfirm

> RegularResponse UserConfirm(ctx).ClientId(clientId).UserName(userName).ConfirmationCode(confirmationCode).Execute()

Confirm user signup

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
	clientId := "clientId_example" // string | Cognito Client ID (optional)
	userName := "userName_example" // string | Cognito User Name (optional)
	confirmationCode := "confirmationCode_example" // string | Cognito Signup Confirmation Code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserConfirm(context.Background()).ClientId(clientId).UserName(userName).ConfirmationCode(confirmationCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserConfirm`: RegularResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserConfirm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | Cognito Client ID | 
 **userName** | **string** | Cognito User Name | 
 **confirmationCode** | **string** | Cognito Signup Confirmation Code | 

### Return type

[**RegularResponse**](RegularResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

