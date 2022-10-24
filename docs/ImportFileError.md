# ImportFileError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | Pointer to **int32** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewImportFileError

`func NewImportFileError() *ImportFileError`

NewImportFileError instantiates a new ImportFileError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportFileErrorWithDefaults

`func NewImportFileErrorWithDefaults() *ImportFileError`

NewImportFileErrorWithDefaults instantiates a new ImportFileError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *ImportFileError) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ImportFileError) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *ImportFileError) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *ImportFileError) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetPosition

`func (o *ImportFileError) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ImportFileError) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ImportFileError) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ImportFileError) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ImportFileError) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ImportFileError) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ImportFileError) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ImportFileError) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


