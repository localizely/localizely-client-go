# ProjectLocaleStatsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LangCode** | Pointer to **string** | Language code (ie &#x60;en&#x60; or &#x60;en-US&#x60;) | [optional] 
**LangName** | Pointer to **string** | Language name (ie &#x60;English&#x60; or &#x60;English (US)&#x60;) | [optional] 
**Strings** | Pointer to **int32** | Total number of string keys in the project | [optional] 
**Reviewed** | Pointer to **int32** | Number of reviewed string keys for a language | [optional] 
**ReviewedProgress** | Pointer to **int32** | Reviewed progress for a language, in percentage | [optional] 

## Methods

### NewProjectLocaleStatsDto

`func NewProjectLocaleStatsDto() *ProjectLocaleStatsDto`

NewProjectLocaleStatsDto instantiates a new ProjectLocaleStatsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectLocaleStatsDtoWithDefaults

`func NewProjectLocaleStatsDtoWithDefaults() *ProjectLocaleStatsDto`

NewProjectLocaleStatsDtoWithDefaults instantiates a new ProjectLocaleStatsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLangCode

`func (o *ProjectLocaleStatsDto) GetLangCode() string`

GetLangCode returns the LangCode field if non-nil, zero value otherwise.

### GetLangCodeOk

`func (o *ProjectLocaleStatsDto) GetLangCodeOk() (*string, bool)`

GetLangCodeOk returns a tuple with the LangCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangCode

`func (o *ProjectLocaleStatsDto) SetLangCode(v string)`

SetLangCode sets LangCode field to given value.

### HasLangCode

`func (o *ProjectLocaleStatsDto) HasLangCode() bool`

HasLangCode returns a boolean if a field has been set.

### GetLangName

`func (o *ProjectLocaleStatsDto) GetLangName() string`

GetLangName returns the LangName field if non-nil, zero value otherwise.

### GetLangNameOk

`func (o *ProjectLocaleStatsDto) GetLangNameOk() (*string, bool)`

GetLangNameOk returns a tuple with the LangName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangName

`func (o *ProjectLocaleStatsDto) SetLangName(v string)`

SetLangName sets LangName field to given value.

### HasLangName

`func (o *ProjectLocaleStatsDto) HasLangName() bool`

HasLangName returns a boolean if a field has been set.

### GetStrings

`func (o *ProjectLocaleStatsDto) GetStrings() int32`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *ProjectLocaleStatsDto) GetStringsOk() (*int32, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *ProjectLocaleStatsDto) SetStrings(v int32)`

SetStrings sets Strings field to given value.

### HasStrings

`func (o *ProjectLocaleStatsDto) HasStrings() bool`

HasStrings returns a boolean if a field has been set.

### GetReviewed

`func (o *ProjectLocaleStatsDto) GetReviewed() int32`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *ProjectLocaleStatsDto) GetReviewedOk() (*int32, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *ProjectLocaleStatsDto) SetReviewed(v int32)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *ProjectLocaleStatsDto) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetReviewedProgress

`func (o *ProjectLocaleStatsDto) GetReviewedProgress() int32`

GetReviewedProgress returns the ReviewedProgress field if non-nil, zero value otherwise.

### GetReviewedProgressOk

`func (o *ProjectLocaleStatsDto) GetReviewedProgressOk() (*int32, bool)`

GetReviewedProgressOk returns a tuple with the ReviewedProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedProgress

`func (o *ProjectLocaleStatsDto) SetReviewedProgress(v int32)`

SetReviewedProgress sets ReviewedProgress field to given value.

### HasReviewedProgress

`func (o *ProjectLocaleStatsDto) HasReviewedProgress() bool`

HasReviewedProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


