# \BranchAPIAPI

All URIs are relative to *https://api.localizely.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBranch**](BranchAPIAPI.md#CreateBranch) | **Post** /v1/projects/{project_id}/branches/{branch} | Create a new branch



## CreateBranch

> CreateBranch(ctx, projectId, branch).SourceBranch(sourceBranch).Execute()

Create a new branch

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Project ID - Can be found on 'My projects' page
	branch := "branch_example" // string | Name of the branch to be created
	sourceBranch := "sourceBranch_example" // string | Name of the source branch from which new branch will be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BranchAPIAPI.CreateBranch(context.Background(), projectId, branch).SourceBranch(sourceBranch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BranchAPIAPI.CreateBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - Can be found on &#39;My projects&#39; page | 
**branch** | **string** | Name of the branch to be created | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sourceBranch** | **string** | Name of the source branch from which new branch will be created | 

### Return type

 (empty response body)

### Authorization

[API auth](../README.md#API auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

