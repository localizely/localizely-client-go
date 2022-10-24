# ErrorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ErrorData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewErrorDto

`func NewErrorDto() *ErrorDto`

NewErrorDto instantiates a new ErrorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDtoWithDefaults

`func NewErrorDtoWithDefaults() *ErrorDto`

NewErrorDtoWithDefaults instantiates a new ErrorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ErrorDto) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorDto) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorDto) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ErrorDto) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ErrorDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ErrorDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ErrorDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ErrorDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorData

`func (o *ErrorDto) GetErrorData() map[string]map[string]interface{}`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *ErrorDto) GetErrorDataOk() (*map[string]map[string]interface{}, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *ErrorDto) SetErrorData(v map[string]map[string]interface{})`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *ErrorDto) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


