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

// CollectorResult struct for CollectorResult
type CollectorResult struct {
	Fragment         *bool             `json:"fragment,omitempty"`
	Name             *string           `json:"name,omitempty"`
	ProfiledChildren []CollectorResult `json:"profiledChildren,omitempty"`
	Reason           *string           `json:"reason,omitempty"`
	Time             *int64            `json:"time,omitempty"`
}

// NewCollectorResult instantiates a new CollectorResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectorResult() *CollectorResult {
	this := CollectorResult{}
	return &this
}

// NewCollectorResultWithDefaults instantiates a new CollectorResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectorResultWithDefaults() *CollectorResult {
	this := CollectorResult{}
	return &this
}

// GetFragment returns the Fragment field value if set, zero value otherwise.
func (o *CollectorResult) GetFragment() bool {
	if o == nil || o.Fragment == nil {
		var ret bool
		return ret
	}
	return *o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectorResult) GetFragmentOk() (*bool, bool) {
	if o == nil || o.Fragment == nil {
		return nil, false
	}
	return o.Fragment, true
}

// HasFragment returns a boolean if a field has been set.
func (o *CollectorResult) HasFragment() bool {
	if o != nil && o.Fragment != nil {
		return true
	}

	return false
}

// SetFragment gets a reference to the given bool and assigns it to the Fragment field.
func (o *CollectorResult) SetFragment(v bool) {
	o.Fragment = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CollectorResult) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectorResult) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CollectorResult) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CollectorResult) SetName(v string) {
	o.Name = &v
}

// GetProfiledChildren returns the ProfiledChildren field value if set, zero value otherwise.
func (o *CollectorResult) GetProfiledChildren() []CollectorResult {
	if o == nil || o.ProfiledChildren == nil {
		var ret []CollectorResult
		return ret
	}
	return o.ProfiledChildren
}

// GetProfiledChildrenOk returns a tuple with the ProfiledChildren field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectorResult) GetProfiledChildrenOk() ([]CollectorResult, bool) {
	if o == nil || o.ProfiledChildren == nil {
		return nil, false
	}
	return o.ProfiledChildren, true
}

// HasProfiledChildren returns a boolean if a field has been set.
func (o *CollectorResult) HasProfiledChildren() bool {
	if o != nil && o.ProfiledChildren != nil {
		return true
	}

	return false
}

// SetProfiledChildren gets a reference to the given []CollectorResult and assigns it to the ProfiledChildren field.
func (o *CollectorResult) SetProfiledChildren(v []CollectorResult) {
	o.ProfiledChildren = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CollectorResult) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectorResult) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CollectorResult) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CollectorResult) SetReason(v string) {
	o.Reason = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *CollectorResult) GetTime() int64 {
	if o == nil || o.Time == nil {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectorResult) GetTimeOk() (*int64, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *CollectorResult) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *CollectorResult) SetTime(v int64) {
	o.Time = &v
}

func (o CollectorResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fragment != nil {
		toSerialize["fragment"] = o.Fragment
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProfiledChildren != nil {
		toSerialize["profiledChildren"] = o.ProfiledChildren
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableCollectorResult struct {
	value *CollectorResult
	isSet bool
}

func (v NullableCollectorResult) Get() *CollectorResult {
	return v.value
}

func (v *NullableCollectorResult) Set(val *CollectorResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectorResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectorResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectorResult(val *CollectorResult) *NullableCollectorResult {
	return &NullableCollectorResult{value: val, isSet: true}
}

func (v NullableCollectorResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectorResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
