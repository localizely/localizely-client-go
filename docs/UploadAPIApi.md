# \UploadAPIApi

All URIs are relative to *https://api.localizely.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportLocalizationFile**](UploadAPIApi.md#ImportLocalizationFile) | **Post** /v1/projects/{project_id}/files/upload | Upload translations for a language



## ImportLocalizationFile

> ImportLocalizationFile(ctx, projectId).LangCode(langCode).File(file).Branch(branch).Overwrite(overwrite).Reviewed(reviewed).TagAdded(tagAdded).TagUpdated(tagUpdated).TagRemoved(tagRemoved).Execute()

Upload translations for a language

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "projectId_example" // string | Project ID - Can be found on 'My projects' page
    langCode := "langCode_example" // string | Language to upload, specified as language code. e.g. `en`, `en_GB` or `en-GB`
    file := os.NewFile(1234, "some_file") // *os.File | Uploading file. Supported following formats: `Flutter ARB, Android XML, iOS strings, iOS stringsdict, Angular XLF, Gettext PO, Gettext POT, Java properties, Ruby on Rails yaml, .NET resx, flat json, csv, Excel .xlsx, Excel .xls`
    branch := "branch_example" // string | Name of the branch to upload file into. Only in case of activated branching feature. (optional)
    overwrite := true // bool | If translation in given language should be overwritten with modified translation from uploading file. (optional) (default to false)
    reviewed := true // bool | If uploading translations, that are added, should be marked as Reviewed. For uploading translations that are only modified it will have effect only if `overwrite` is set to `true`. (optional) (default to false)
    tagAdded := []string{"Inner_example"} // []string | Optional list of tags to add to new translations from uploading file. <br><br>Multiple tags can be defined in a following way: `&tag_added_keys=NEW&tag_added_keys=NEW_SPRINT05` (optional)
    tagUpdated := []string{"Inner_example"} // []string | Optional list of tags to add to updated translations from uploading file. <br><br>Multiple tags can be defined in a following way: `&tag_updated_keys=UPDATED&tag_updated_keys=UPDATED_SPRINT05` (optional)
    tagRemoved := []string{"Inner_example"} // []string | Optional list of tags to add to removed translations from uploading file. <br><br>Multiple tags can be defined in a following way: `&tag_removed_keys=REMOVED&tag_removed_keys=REMOVED_SPRINT05` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadAPIApi.ImportLocalizationFile(context.Background(), projectId).LangCode(langCode).File(file).Branch(branch).Overwrite(overwrite).Reviewed(reviewed).TagAdded(tagAdded).TagUpdated(tagUpdated).TagRemoved(tagRemoved).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadAPIApi.ImportLocalizationFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiImportLocalizationFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **langCode** | **string** | Language to upload, specified as language code. e.g. &#x60;en&#x60;, &#x60;en_GB&#x60; or &#x60;en-GB&#x60; | 
 **file** | ***os.File** | Uploading file. Supported following formats: &#x60;Flutter ARB, Android XML, iOS strings, iOS stringsdict, Angular XLF, Gettext PO, Gettext POT, Java properties, Ruby on Rails yaml, .NET resx, flat json, csv, Excel .xlsx, Excel .xls&#x60; | 
 **branch** | **string** | Name of the branch to upload file into. Only in case of activated branching feature. | 
 **overwrite** | **bool** | If translation in given language should be overwritten with modified translation from uploading file. | [default to false]
 **reviewed** | **bool** | If uploading translations, that are added, should be marked as Reviewed. For uploading translations that are only modified it will have effect only if &#x60;overwrite&#x60; is set to &#x60;true&#x60;. | [default to false]
 **tagAdded** | **[]string** | Optional list of tags to add to new translations from uploading file. &lt;br&gt;&lt;br&gt;Multiple tags can be defined in a following way: &#x60;&amp;tag_added_keys&#x3D;NEW&amp;tag_added_keys&#x3D;NEW_SPRINT05&#x60; | 
 **tagUpdated** | **[]string** | Optional list of tags to add to updated translations from uploading file. &lt;br&gt;&lt;br&gt;Multiple tags can be defined in a following way: &#x60;&amp;tag_updated_keys&#x3D;UPDATED&amp;tag_updated_keys&#x3D;UPDATED_SPRINT05&#x60; | 
 **tagRemoved** | **[]string** | Optional list of tags to add to removed translations from uploading file. &lt;br&gt;&lt;br&gt;Multiple tags can be defined in a following way: &#x60;&amp;tag_removed_keys&#x3D;REMOVED&amp;tag_removed_keys&#x3D;REMOVED_SPRINT05&#x60; | 

### Return type

 (empty response body)

### Authorization

[API auth](../README.md#API auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

