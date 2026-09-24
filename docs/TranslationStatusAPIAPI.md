# \TranslationStatusAPIAPI

All URIs are relative to *https://api.localizely.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTranslationStatus**](TranslationStatusAPIAPI.md#GetTranslationStatus) | **Get** /v1/projects/{project_id}/status | Get Translation Status for the project



## GetTranslationStatus

> ProjectStatusDto GetTranslationStatus(ctx, projectId).Branch(branch).Execute()

Get Translation Status for the project

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/localizely/localizely-client-go"
)

func main() {
	projectId := "projectId_example" // string | Project ID - Can be found on 'My projects' page
	branch := "branch_example" // string | Name of the branch to get translation status for. Only in case of activated branching feature. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TranslationStatusAPIAPI.GetTranslationStatus(context.Background(), projectId).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TranslationStatusAPIAPI.GetTranslationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTranslationStatus`: ProjectStatusDto
	fmt.Fprintf(os.Stdout, "Response from `TranslationStatusAPIAPI.GetTranslationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - Can be found on &#39;My projects&#39; page | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTranslationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **string** | Name of the branch to get translation status for. Only in case of activated branching feature. | 

### Return type

[**ProjectStatusDto**](ProjectStatusDto.md)

### Authorization

[API auth](../README.md#API auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

