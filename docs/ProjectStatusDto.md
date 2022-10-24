# ProjectStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strings** | Pointer to **int32** | Total number of string keys in the project | [optional] 
**ReviewedProgress** | Pointer to **int32** | Total reviewed progress across all languages, in percentage | [optional] 
**Languages** | Pointer to [**[]ProjectLocaleStatsDto**](ProjectLocaleStatsDto.md) | Translation status per language | [optional] 

## Methods

### NewProjectStatusDto

`func NewProjectStatusDto() *ProjectStatusDto`

NewProjectStatusDto instantiates a new ProjectStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectStatusDtoWithDefaults

`func NewProjectStatusDtoWithDefaults() *ProjectStatusDto`

NewProjectStatusDtoWithDefaults instantiates a new ProjectStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrings

`func (o *ProjectStatusDto) GetStrings() int32`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *ProjectStatusDto) GetStringsOk() (*int32, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *ProjectStatusDto) SetStrings(v int32)`

SetStrings sets Strings field to given value.

### HasStrings

`func (o *ProjectStatusDto) HasStrings() bool`

HasStrings returns a boolean if a field has been set.

### GetReviewedProgress

`func (o *ProjectStatusDto) GetReviewedProgress() int32`

GetReviewedProgress returns the ReviewedProgress field if non-nil, zero value otherwise.

### GetReviewedProgressOk

`func (o *ProjectStatusDto) GetReviewedProgressOk() (*int32, bool)`

GetReviewedProgressOk returns a tuple with the ReviewedProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedProgress

`func (o *ProjectStatusDto) SetReviewedProgress(v int32)`

SetReviewedProgress sets ReviewedProgress field to given value.

### HasReviewedProgress

`func (o *ProjectStatusDto) HasReviewedProgress() bool`

HasReviewedProgress returns a boolean if a field has been set.

### GetLanguages

`func (o *ProjectStatusDto) GetLanguages() []ProjectLocaleStatsDto`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ProjectStatusDto) GetLanguagesOk() (*[]ProjectLocaleStatsDto, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ProjectStatusDto) SetLanguages(v []ProjectLocaleStatsDto)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ProjectStatusDto) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


