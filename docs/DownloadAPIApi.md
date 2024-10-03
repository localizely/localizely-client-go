# \DownloadAPIAPI

All URIs are relative to *https://api.localizely.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocalizationFile**](DownloadAPIAPI.md#GetLocalizationFile) | **Get** /v1/projects/{project_id}/files/download | Download translations for a language in a specified file format



## GetLocalizationFile

> GetLocalizationFile(ctx, projectId).Type_(type_).Branch(branch).LangCodes(langCodes).JavaPropertiesEncoding(javaPropertiesEncoding).IncludeTags(includeTags).ExcludeTags(excludeTags).ExportEmptyAs(exportEmptyAs).Execute()

Download translations for a language in a specified file format



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
	type_ := "type__example" // string | File format
	branch := "branch_example" // string | Name of the branch to download file from. Only in case of activated branching feature. (optional)
	langCodes := "langCodes_example" // string | Language to download, specified as language code. e.g. `en`, `en_GB` or `en-GB`. For multiple languages use comma separator. If omitted, all languages are downloaded. (optional)
	javaPropertiesEncoding := "javaPropertiesEncoding_example" // string | (Only for Java .properties files download) Character encoding. Default is `latin_1`. (optional)
	includeTags := []string{"Inner_example"} // []string | Optional list of tags to be downloaded. <br>If not set, all string keys will be considered for download. <br><br>Multiple tags can be defined in a following way: `&include_tags=ANDROID&include_tags=ANDROID_SPRINT05`. (optional)
	excludeTags := []string{"Inner_example"} // []string | Optional list of tags to be excluded from download. <br>If not set, all string keys will be considered for download. <br><br> Multiple tags can be defined in a following way: `&exclude_tags=REMOVED&exclude_tags=REMOVED_SPRINT05`. (optional)
	exportEmptyAs := "exportEmptyAs_example" // string | Optional. How you would like empty translations to be exported. Allowed values are `empty` to keep empty, `main` to replace with the main language value, or `skip` to omit. (optional) (default to "empty")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DownloadAPIAPI.GetLocalizationFile(context.Background(), projectId).Type_(type_).Branch(branch).LangCodes(langCodes).JavaPropertiesEncoding(javaPropertiesEncoding).IncludeTags(includeTags).ExcludeTags(excludeTags).ExportEmptyAs(exportEmptyAs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPIAPI.GetLocalizationFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID - Can be found on &#39;My projects&#39; page | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | File format | 
 **branch** | **string** | Name of the branch to download file from. Only in case of activated branching feature. | 
 **langCodes** | **string** | Language to download, specified as language code. e.g. &#x60;en&#x60;, &#x60;en_GB&#x60; or &#x60;en-GB&#x60;. For multiple languages use comma separator. If omitted, all languages are downloaded. | 
 **javaPropertiesEncoding** | **string** | (Only for Java .properties files download) Character encoding. Default is &#x60;latin_1&#x60;. | 
 **includeTags** | **[]string** | Optional list of tags to be downloaded. &lt;br&gt;If not set, all string keys will be considered for download. &lt;br&gt;&lt;br&gt;Multiple tags can be defined in a following way: &#x60;&amp;include_tags&#x3D;ANDROID&amp;include_tags&#x3D;ANDROID_SPRINT05&#x60;. | 
 **excludeTags** | **[]string** | Optional list of tags to be excluded from download. &lt;br&gt;If not set, all string keys will be considered for download. &lt;br&gt;&lt;br&gt; Multiple tags can be defined in a following way: &#x60;&amp;exclude_tags&#x3D;REMOVED&amp;exclude_tags&#x3D;REMOVED_SPRINT05&#x60;. | 
 **exportEmptyAs** | **string** | Optional. How you would like empty translations to be exported. Allowed values are &#x60;empty&#x60; to keep empty, &#x60;main&#x60; to replace with the main language value, or &#x60;skip&#x60; to omit. | [default to &quot;empty&quot;]

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

