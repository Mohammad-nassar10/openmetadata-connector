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

// Aggregations struct for Aggregations
type Aggregations struct {
	AsMap    *map[string]Aggregation `json:"asMap,omitempty"`
	Fragment *bool                   `json:"fragment,omitempty"`
}

// NewAggregations instantiates a new Aggregations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregations() *Aggregations {
	this := Aggregations{}
	return &this
}

// NewAggregationsWithDefaults instantiates a new Aggregations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationsWithDefaults() *Aggregations {
	this := Aggregations{}
	return &this
}

// GetAsMap returns the AsMap field value if set, zero value otherwise.
func (o *Aggregations) GetAsMap() map[string]Aggregation {
	if o == nil || o.AsMap == nil {
		var ret map[string]Aggregation
		return ret
	}
	return *o.AsMap
}

// GetAsMapOk returns a tuple with the AsMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregations) GetAsMapOk() (*map[string]Aggregation, bool) {
	if o == nil || o.AsMap == nil {
		return nil, false
	}
	return o.AsMap, true
}

// HasAsMap returns a boolean if a field has been set.
func (o *Aggregations) HasAsMap() bool {
	if o != nil && o.AsMap != nil {
		return true
	}

	return false
}

// SetAsMap gets a reference to the given map[string]Aggregation and assigns it to the AsMap field.
func (o *Aggregations) SetAsMap(v map[string]Aggregation) {
	o.AsMap = &v
}

// GetFragment returns the Fragment field value if set, zero value otherwise.
func (o *Aggregations) GetFragment() bool {
	if o == nil || o.Fragment == nil {
		var ret bool
		return ret
	}
	return *o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregations) GetFragmentOk() (*bool, bool) {
	if o == nil || o.Fragment == nil {
		return nil, false
	}
	return o.Fragment, true
}

// HasFragment returns a boolean if a field has been set.
func (o *Aggregations) HasFragment() bool {
	if o != nil && o.Fragment != nil {
		return true
	}

	return false
}

// SetFragment gets a reference to the given bool and assigns it to the Fragment field.
func (o *Aggregations) SetFragment(v bool) {
	o.Fragment = &v
}

func (o Aggregations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AsMap != nil {
		toSerialize["asMap"] = o.AsMap
	}
	if o.Fragment != nil {
		toSerialize["fragment"] = o.Fragment
	}
	return json.Marshal(toSerialize)
}

type NullableAggregations struct {
	value *Aggregations
	isSet bool
}

func (v NullableAggregations) Get() *Aggregations {
	return v.value
}

func (v *NullableAggregations) Set(val *Aggregations) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregations) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregations(val *Aggregations) *NullableAggregations {
	return &NullableAggregations{value: val, isSet: true}
}

func (v NullableAggregations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
