# InvalidImportFileErrorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ErrorData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Errors** | Pointer to [**[]ImportFileError**](ImportFileError.md) |  | [optional] 

## Methods

### NewInvalidImportFileErrorDto

`func NewInvalidImportFileErrorDto() *InvalidImportFileErrorDto`

NewInvalidImportFileErrorDto instantiates a new InvalidImportFileErrorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidImportFileErrorDtoWithDefaults

`func NewInvalidImportFileErrorDtoWithDefaults() *InvalidImportFileErrorDto`

NewInvalidImportFileErrorDtoWithDefaults instantiates a new InvalidImportFileErrorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *InvalidImportFileErrorDto) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *InvalidImportFileErrorDto) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *InvalidImportFileErrorDto) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *InvalidImportFileErrorDto) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *InvalidImportFileErrorDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InvalidImportFileErrorDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InvalidImportFileErrorDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InvalidImportFileErrorDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorData

`func (o *InvalidImportFileErrorDto) GetErrorData() map[string]map[string]interface{}`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *InvalidImportFileErrorDto) GetErrorDataOk() (*map[string]map[string]interface{}, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *InvalidImportFileErrorDto) SetErrorData(v map[string]map[string]interface{})`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *InvalidImportFileErrorDto) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.

### GetErrors

`func (o *InvalidImportFileErrorDto) GetErrors() []ImportFileError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InvalidImportFileErrorDto) GetErrorsOk() (*[]ImportFileError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InvalidImportFileErrorDto) SetErrors(v []ImportFileError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InvalidImportFileErrorDto) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


