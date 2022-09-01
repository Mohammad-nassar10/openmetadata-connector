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

// SearchResponseSections struct for SearchResponseSections
type SearchResponseSections struct {
	Fragment        *bool  `json:"fragment,omitempty"`
	NumReducePhases *int32 `json:"numReducePhases,omitempty"`
}

// NewSearchResponseSections instantiates a new SearchResponseSections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResponseSections() *SearchResponseSections {
	this := SearchResponseSections{}
	return &this
}

// NewSearchResponseSectionsWithDefaults instantiates a new SearchResponseSections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResponseSectionsWithDefaults() *SearchResponseSections {
	this := SearchResponseSections{}
	return &this
}

// GetFragment returns the Fragment field value if set, zero value otherwise.
func (o *SearchResponseSections) GetFragment() bool {
	if o == nil || o.Fragment == nil {
		var ret bool
		return ret
	}
	return *o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseSections) GetFragmentOk() (*bool, bool) {
	if o == nil || o.Fragment == nil {
		return nil, false
	}
	return o.Fragment, true
}

// HasFragment returns a boolean if a field has been set.
func (o *SearchResponseSections) HasFragment() bool {
	if o != nil && o.Fragment != nil {
		return true
	}

	return false
}

// SetFragment gets a reference to the given bool and assigns it to the Fragment field.
func (o *SearchResponseSections) SetFragment(v bool) {
	o.Fragment = &v
}

// GetNumReducePhases returns the NumReducePhases field value if set, zero value otherwise.
func (o *SearchResponseSections) GetNumReducePhases() int32 {
	if o == nil || o.NumReducePhases == nil {
		var ret int32
		return ret
	}
	return *o.NumReducePhases
}

// GetNumReducePhasesOk returns a tuple with the NumReducePhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseSections) GetNumReducePhasesOk() (*int32, bool) {
	if o == nil || o.NumReducePhases == nil {
		return nil, false
	}
	return o.NumReducePhases, true
}

// HasNumReducePhases returns a boolean if a field has been set.
func (o *SearchResponseSections) HasNumReducePhases() bool {
	if o != nil && o.NumReducePhases != nil {
		return true
	}

	return false
}

// SetNumReducePhases gets a reference to the given int32 and assigns it to the NumReducePhases field.
func (o *SearchResponseSections) SetNumReducePhases(v int32) {
	o.NumReducePhases = &v
}

func (o SearchResponseSections) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fragment != nil {
		toSerialize["fragment"] = o.Fragment
	}
	if o.NumReducePhases != nil {
		toSerialize["numReducePhases"] = o.NumReducePhases
	}
	return json.Marshal(toSerialize)
}

type NullableSearchResponseSections struct {
	value *SearchResponseSections
	isSet bool
}

func (v NullableSearchResponseSections) Get() *SearchResponseSections {
	return v.value
}

func (v *NullableSearchResponseSections) Set(val *SearchResponseSections) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResponseSections) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResponseSections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResponseSections(val *SearchResponseSections) *NullableSearchResponseSections {
	return &NullableSearchResponseSections{value: val, isSet: true}
}

func (v NullableSearchResponseSections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResponseSections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
