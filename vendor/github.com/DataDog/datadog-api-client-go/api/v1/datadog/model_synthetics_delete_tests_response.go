/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsDeleteTestsResponse Response object for deleting Synthetic test.
type SyntheticsDeleteTestsResponse struct {
	// Array of objects containing a deleted Synthetic test ID with the associated deletion timestamp.
	DeletedTests *[]SyntheticsDeleteTestsResponseDeletedTests `json:"deleted_tests,omitempty"`
}

// NewSyntheticsDeleteTestsResponse instantiates a new SyntheticsDeleteTestsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsDeleteTestsResponse() *SyntheticsDeleteTestsResponse {
	this := SyntheticsDeleteTestsResponse{}
	return &this
}

// NewSyntheticsDeleteTestsResponseWithDefaults instantiates a new SyntheticsDeleteTestsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsDeleteTestsResponseWithDefaults() *SyntheticsDeleteTestsResponse {
	this := SyntheticsDeleteTestsResponse{}
	return &this
}

// GetDeletedTests returns the DeletedTests field value if set, zero value otherwise.
func (o *SyntheticsDeleteTestsResponse) GetDeletedTests() []SyntheticsDeleteTestsResponseDeletedTests {
	if o == nil || o.DeletedTests == nil {
		var ret []SyntheticsDeleteTestsResponseDeletedTests
		return ret
	}
	return *o.DeletedTests
}

// GetDeletedTestsOk returns a tuple with the DeletedTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDeleteTestsResponse) GetDeletedTestsOk() (*[]SyntheticsDeleteTestsResponseDeletedTests, bool) {
	if o == nil || o.DeletedTests == nil {
		return nil, false
	}
	return o.DeletedTests, true
}

// HasDeletedTests returns a boolean if a field has been set.
func (o *SyntheticsDeleteTestsResponse) HasDeletedTests() bool {
	if o != nil && o.DeletedTests != nil {
		return true
	}

	return false
}

// SetDeletedTests gets a reference to the given []SyntheticsDeleteTestsResponseDeletedTests and assigns it to the DeletedTests field.
func (o *SyntheticsDeleteTestsResponse) SetDeletedTests(v []SyntheticsDeleteTestsResponseDeletedTests) {
	o.DeletedTests = &v
}

func (o SyntheticsDeleteTestsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeletedTests != nil {
		toSerialize["deleted_tests"] = o.DeletedTests
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsDeleteTestsResponse struct {
	value *SyntheticsDeleteTestsResponse
	isSet bool
}

func (v NullableSyntheticsDeleteTestsResponse) Get() *SyntheticsDeleteTestsResponse {
	return v.value
}

func (v *NullableSyntheticsDeleteTestsResponse) Set(val *SyntheticsDeleteTestsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsDeleteTestsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsDeleteTestsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsDeleteTestsResponse(val *SyntheticsDeleteTestsResponse) *NullableSyntheticsDeleteTestsResponse {
	return &NullableSyntheticsDeleteTestsResponse{value: val, isSet: true}
}

func (v NullableSyntheticsDeleteTestsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsDeleteTestsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
