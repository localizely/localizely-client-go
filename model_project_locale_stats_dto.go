/*
Localizely API

<h2>Getting started</h2><p>Localizely API is built on <a href=\"https://en.wikipedia.org/wiki/Representational_state_transfer\" target=\"_blank\">REST</a>. <br>You can use this API for importing & exporting your localization files in order to automate the process with `curl` scripts or external <a href=\"https://en.wikipedia.org/wiki/Continuous_integration\" target=\"_blank\">CI</a> tools. <br>Response is returned in JSON form even in case of error. <br></p><p>If you Authenticate with your API token on this page by clicking \"Authorize\" button, you can make API calls directly from here with \"Try it out\", and generate such `curl` examples. </p><h2>API Authentication</h2><p>Authenticate your account by sending your API token as a request header `X-Api-Token`. <br>The token can be found under <a href=\"https://app.localizely.com/account\" target=\"_blank\">My Profile</a> page. <br>A user must have an Admin role in the project in order to access the project with his token. <br>API requests without authentication will fail.</p><p><b>Base url:</b> `https://api.localizely.com`</p>

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package localizely

import (
	"encoding/json"
)

// ProjectLocaleStatsDto Translation status per language
type ProjectLocaleStatsDto struct {
	// Language code (ie `en` or `en-US`)
	LangCode *string `json:"langCode,omitempty"`
	// Language name (ie `English` or `English (US)`)
	LangName *string `json:"langName,omitempty"`
	// Total number of string keys in the project
	Strings *int32 `json:"strings,omitempty"`
	// Number of reviewed string keys for a language
	Reviewed *int32 `json:"reviewed,omitempty"`
	// Reviewed progress for a language, in percentage
	ReviewedProgress *int32 `json:"reviewedProgress,omitempty"`
}

// NewProjectLocaleStatsDto instantiates a new ProjectLocaleStatsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectLocaleStatsDto() *ProjectLocaleStatsDto {
	this := ProjectLocaleStatsDto{}
	return &this
}

// NewProjectLocaleStatsDtoWithDefaults instantiates a new ProjectLocaleStatsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectLocaleStatsDtoWithDefaults() *ProjectLocaleStatsDto {
	this := ProjectLocaleStatsDto{}
	return &this
}

// GetLangCode returns the LangCode field value if set, zero value otherwise.
func (o *ProjectLocaleStatsDto) GetLangCode() string {
	if o == nil || o.LangCode == nil {
		var ret string
		return ret
	}
	return *o.LangCode
}

// GetLangCodeOk returns a tuple with the LangCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLocaleStatsDto) GetLangCodeOk() (*string, bool) {
	if o == nil || o.LangCode == nil {
		return nil, false
	}
	return o.LangCode, true
}

// HasLangCode returns a boolean if a field has been set.
func (o *ProjectLocaleStatsDto) HasLangCode() bool {
	if o != nil && o.LangCode != nil {
		return true
	}

	return false
}

// SetLangCode gets a reference to the given string and assigns it to the LangCode field.
func (o *ProjectLocaleStatsDto) SetLangCode(v string) {
	o.LangCode = &v
}

// GetLangName returns the LangName field value if set, zero value otherwise.
func (o *ProjectLocaleStatsDto) GetLangName() string {
	if o == nil || o.LangName == nil {
		var ret string
		return ret
	}
	return *o.LangName
}

// GetLangNameOk returns a tuple with the LangName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLocaleStatsDto) GetLangNameOk() (*string, bool) {
	if o == nil || o.LangName == nil {
		return nil, false
	}
	return o.LangName, true
}

// HasLangName returns a boolean if a field has been set.
func (o *ProjectLocaleStatsDto) HasLangName() bool {
	if o != nil && o.LangName != nil {
		return true
	}

	return false
}

// SetLangName gets a reference to the given string and assigns it to the LangName field.
func (o *ProjectLocaleStatsDto) SetLangName(v string) {
	o.LangName = &v
}

// GetStrings returns the Strings field value if set, zero value otherwise.
func (o *ProjectLocaleStatsDto) GetStrings() int32 {
	if o == nil || o.Strings == nil {
		var ret int32
		return ret
	}
	return *o.Strings
}

// GetStringsOk returns a tuple with the Strings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLocaleStatsDto) GetStringsOk() (*int32, bool) {
	if o == nil || o.Strings == nil {
		return nil, false
	}
	return o.Strings, true
}

// HasStrings returns a boolean if a field has been set.
func (o *ProjectLocaleStatsDto) HasStrings() bool {
	if o != nil && o.Strings != nil {
		return true
	}

	return false
}

// SetStrings gets a reference to the given int32 and assigns it to the Strings field.
func (o *ProjectLocaleStatsDto) SetStrings(v int32) {
	o.Strings = &v
}

// GetReviewed returns the Reviewed field value if set, zero value otherwise.
func (o *ProjectLocaleStatsDto) GetReviewed() int32 {
	if o == nil || o.Reviewed == nil {
		var ret int32
		return ret
	}
	return *o.Reviewed
}

// GetReviewedOk returns a tuple with the Reviewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLocaleStatsDto) GetReviewedOk() (*int32, bool) {
	if o == nil || o.Reviewed == nil {
		return nil, false
	}
	return o.Reviewed, true
}

// HasReviewed returns a boolean if a field has been set.
func (o *ProjectLocaleStatsDto) HasReviewed() bool {
	if o != nil && o.Reviewed != nil {
		return true
	}

	return false
}

// SetReviewed gets a reference to the given int32 and assigns it to the Reviewed field.
func (o *ProjectLocaleStatsDto) SetReviewed(v int32) {
	o.Reviewed = &v
}

// GetReviewedProgress returns the ReviewedProgress field value if set, zero value otherwise.
func (o *ProjectLocaleStatsDto) GetReviewedProgress() int32 {
	if o == nil || o.ReviewedProgress == nil {
		var ret int32
		return ret
	}
	return *o.ReviewedProgress
}

// GetReviewedProgressOk returns a tuple with the ReviewedProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectLocaleStatsDto) GetReviewedProgressOk() (*int32, bool) {
	if o == nil || o.ReviewedProgress == nil {
		return nil, false
	}
	return o.ReviewedProgress, true
}

// HasReviewedProgress returns a boolean if a field has been set.
func (o *ProjectLocaleStatsDto) HasReviewedProgress() bool {
	if o != nil && o.ReviewedProgress != nil {
		return true
	}

	return false
}

// SetReviewedProgress gets a reference to the given int32 and assigns it to the ReviewedProgress field.
func (o *ProjectLocaleStatsDto) SetReviewedProgress(v int32) {
	o.ReviewedProgress = &v
}

func (o ProjectLocaleStatsDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LangCode != nil {
		toSerialize["langCode"] = o.LangCode
	}
	if o.LangName != nil {
		toSerialize["langName"] = o.LangName
	}
	if o.Strings != nil {
		toSerialize["strings"] = o.Strings
	}
	if o.Reviewed != nil {
		toSerialize["reviewed"] = o.Reviewed
	}
	if o.ReviewedProgress != nil {
		toSerialize["reviewedProgress"] = o.ReviewedProgress
	}
	return json.Marshal(toSerialize)
}

type NullableProjectLocaleStatsDto struct {
	value *ProjectLocaleStatsDto
	isSet bool
}

func (v NullableProjectLocaleStatsDto) Get() *ProjectLocaleStatsDto {
	return v.value
}

func (v *NullableProjectLocaleStatsDto) Set(val *ProjectLocaleStatsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectLocaleStatsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectLocaleStatsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectLocaleStatsDto(val *ProjectLocaleStatsDto) *NullableProjectLocaleStatsDto {
	return &NullableProjectLocaleStatsDto{value: val, isSet: true}
}

func (v NullableProjectLocaleStatsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectLocaleStatsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

