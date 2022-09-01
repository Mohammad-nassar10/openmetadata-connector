/*
OpenMetadata Apis

# Overview  OpenMetadata supports REST APIs for getting metadata in and out of metadata store. The API resources are grouped under following categories: - **Data assets** - includes resources for data entities, such as `databases`, `tables`, and `topics`. Resources for data assets created from data, such as `dashboards`, `reports`, `metrics`, and `ML Features` are part of this collection. `pipelines`, `notebooks`, etc. that are used for creating data assets are also available as resources as of this collection. - **Teams and Users** - includes `users`, `teams`, a special type of user called `bots` that performs many automated tasks such as ingestion. - **Services** - are services that OpenMetadata integrates with. Currently `databaseService` is the only service under this collection that represents data sources. In the future, services related to Dashboards, Reports, ETL pipelines will be added under this collection. - **Glossary** - OpenMetadata supports hierarchical tags that can be used to build business vocabulary to describe and classify data available under `tags` resource.

API version: 0.11.0
Contact: openmetadata-dev@googlegroups.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TableTest struct for TableTest
type TableTest struct {
	Description        *string          `json:"description,omitempty"`
	ExecutionFrequency *string          `json:"executionFrequency,omitempty"`
	Id                 *string          `json:"id,omitempty"`
	Name               string           `json:"name"`
	Owner              *EntityReference `json:"owner,omitempty"`
	Results            []TestCaseResult `json:"results,omitempty"`
	TestCase           TableTestCase    `json:"testCase"`
	UpdatedAt          *int64           `json:"updatedAt,omitempty"`
	UpdatedBy          *string          `json:"updatedBy,omitempty"`
}

// NewTableTest instantiates a new TableTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableTest(name string, testCase TableTestCase) *TableTest {
	this := TableTest{}
	this.Name = name
	this.TestCase = testCase
	return &this
}

// NewTableTestWithDefaults instantiates a new TableTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableTestWithDefaults() *TableTest {
	this := TableTest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TableTest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableTest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TableTest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TableTest) SetDescription(v string) {
	o.Description = &v
}

// GetExecutionFrequency returns the ExecutionFrequency field value if set, zero value otherwise.
func (o *TableTest) GetExecutionFrequency() string {
	if o == nil || o.ExecutionFrequency == nil {
		var ret string
		return ret
	}
	return *o.ExecutionFrequency
}

// GetExecutionFrequencyOk returns a tuple with the ExecutionFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableTest) GetExecutionFrequencyOk() (*string, bool) {
	if o == nil || o.ExecutionFrequency == nil {
		return nil, false
	}
	return o.ExecutionFrequency, true
}

// HasExecutionFrequency returns a boolean if a field has been set.
func (o *TableTest) HasExecutionFrequency() bool {
	if o != nil && o.ExecutionFrequency != nil {
		return true
	}

	return false
}

// SetExecutionFrequency gets a reference to the given string and assigns it to the ExecutionFrequency field.
func (o *TableTest) SetExecutionFrequency(v string) {
	o.ExecutionFrequency = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TableTest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableTest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TableTest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TableTest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *TableTest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TableTest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TableTest) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *TableTest) GetOwner() EntityReference {
	if o == nil || o.Owner == nil {
		var ret EntityReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableTest) GetOwnerOk() (*EntityReference, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *TableTest) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given EntityReference and assigns it to the Owner field.
func (o *TableTest) SetOwner(v EntityReference) {
	o.Owner = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *TableTest) GetResults() []TestCaseResult {
	if o == nil || o.Results == nil {
		var ret []TestCaseResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableTest) GetResultsOk() ([]TestCaseResult, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *TableTest) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TestCaseResult and assigns it to the Results field.
func (o *TableTest) SetResults(v []TestCaseResult) {
	o.Results = v
}

// GetTestCase returns the TestCase field value
func (o *TableTest) GetTestCase() TableTestCase {
	if o == nil {
		var ret TableTestCase
		return ret
	}

	return o.TestCase
}

// GetTestCaseOk returns a tuple with the TestCase field value
// and a boolean to check if the value has been set.
func (o *TableTest) GetTestCaseOk() (*TableTestCase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestCase, true
}

// SetTestCase sets field value
func (o *TableTest) SetTestCase(v TableTestCase) {
	o.TestCase = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TableTest) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableTest) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TableTest) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *TableTest) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *TableTest) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableTest) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *TableTest) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *TableTest) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o TableTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExecutionFrequency != nil {
		toSerialize["executionFrequency"] = o.ExecutionFrequency
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if true {
		toSerialize["testCase"] = o.TestCase
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableTableTest struct {
	value *TableTest
	isSet bool
}

func (v NullableTableTest) Get() *TableTest {
	return v.value
}

func (v *NullableTableTest) Set(val *TableTest) {
	v.value = val
	v.isSet = true
}

func (v NullableTableTest) IsSet() bool {
	return v.isSet
}

func (v *NullableTableTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableTest(val *TableTest) *NullableTableTest {
	return &NullableTableTest{value: val, isSet: true}
}

func (v NullableTableTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
