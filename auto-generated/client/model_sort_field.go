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

// SortField struct for SortField
type SortField struct {
	BytesComparator  map[string]interface{} `json:"bytesComparator,omitempty"`
	CanUsePoints     *bool                  `json:"canUsePoints,omitempty"`
	ComparatorSource map[string]interface{} `json:"comparatorSource,omitempty"`
	Field            *string                `json:"field,omitempty"`
	IndexSorter      *IndexSorter           `json:"indexSorter,omitempty"`
	MissingValue     map[string]interface{} `json:"missingValue,omitempty"`
	Reverse          *bool                  `json:"reverse,omitempty"`
	Type             *string                `json:"type,omitempty"`
}

// NewSortField instantiates a new SortField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortField() *SortField {
	this := SortField{}
	return &this
}

// NewSortFieldWithDefaults instantiates a new SortField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortFieldWithDefaults() *SortField {
	this := SortField{}
	return &this
}

// GetBytesComparator returns the BytesComparator field value if set, zero value otherwise.
func (o *SortField) GetBytesComparator() map[string]interface{} {
	if o == nil || o.BytesComparator == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.BytesComparator
}

// GetBytesComparatorOk returns a tuple with the BytesComparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortField) GetBytesComparatorOk() (map[string]interface{}, bool) {
	if o == nil || o.BytesComparator == nil {
		return nil, false
	}
	return o.BytesComparator, true
}

// HasBytesComparator returns a boolean if a field has been set.
func (o *SortField) HasBytesComparator() bool {
	if o != nil && o.BytesComparator != nil {
		return true
	}

	return false
}

// SetBytesComparator gets a reference to the given map[string]interface{} and assigns it to the BytesComparator field.
func (o *SortField) SetBytesComparator(v map[string]interface{}) {
	o.BytesComparator = v
}

// GetCanUsePoints returns the CanUsePoints field value if set, zero value otherwise.
func (o *SortField) GetCanUsePoints() bool {
	if o == nil || o.CanUsePoints == nil {
		var ret bool
		return ret
	}
	return *o.CanUsePoints
}

// GetCanUsePointsOk returns a tuple with the CanUsePoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortField) GetCanUsePointsOk() (*bool, bool) {
	if o == nil || o.CanUsePoints == nil {
		return nil, false
	}
	return o.CanUsePoints, true
}

// HasCanUsePoints returns a boolean if a field has been set.
func (o *SortField) HasCanUsePoints() bool {
	if o != nil && o.CanUsePoints != nil {
		return true
	}

	return false
}

// SetCanUsePoints gets a reference to the given bool and assigns it to the CanUsePoints field.
func (o *SortField) SetCanUsePoints(v bool) {
	o.CanUsePoints = &v
}

// GetComparatorSource returns the ComparatorSource field value if set, zero value otherwise.
func (o *SortField) GetComparatorSource() map[string]interface{} {
	if o == nil || o.ComparatorSource == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ComparatorSource
}

// GetComparatorSourceOk returns a tuple with the ComparatorSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortField) GetComparatorSourceOk() (map[string]interface{}, bool) {
	if o == nil || o.ComparatorSource == nil {
		return nil, false
	}
	return o.ComparatorSource, true
}

// HasComparatorSource returns a boolean if a field has been set.
func (o *SortField) HasComparatorSource() bool {
	if o != nil && o.ComparatorSource != nil {
		return true
	}

	return false
}

// SetComparatorSource gets a reference to the given map[string]interface{} and assigns it to the ComparatorSource field.
func (o *SortField) SetComparatorSource(v map[string]interface{}) {
	o.ComparatorSource = v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *SortField) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortField) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *SortField) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *SortField) SetField(v string) {
	o.Field = &v
}

// GetIndexSorter returns the IndexSorter field value if set, zero value otherwise.
func (o *SortField) GetIndexSorter() IndexSorter {
	if o == nil || o.IndexSorter == nil {
		var ret IndexSorter
		return ret
	}
	return *o.IndexSorter
}

// GetIndexSorterOk returns a tuple with the IndexSorter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortField) GetIndexSorterOk() (*IndexSorter, bool) {
	if o == nil || o.IndexSorter == nil {
		return nil, false
	}
	return o.IndexSorter, true
}

// HasIndexSorter returns a boolean if a field has been set.
func (o *SortField) HasIndexSorter() bool {
	if o != nil && o.IndexSorter != nil {
		return true
	}

	return false
}

// SetIndexSorter gets a reference to the given IndexSorter and assigns it to the IndexSorter field.
func (o *SortField) SetIndexSorter(v IndexSorter) {
	o.IndexSorter = &v
}

// GetMissingValue returns the MissingValue field value if set, zero value otherwise.
func (o *SortField) GetMissingValue() map[string]interface{} {
	if o == nil || o.MissingValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MissingValue
}

// GetMissingValueOk returns a tuple with the MissingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortField) GetMissingValueOk() (map[string]interface{}, bool) {
	if o == nil || o.MissingValue == nil {
		return nil, false
	}
	return o.MissingValue, true
}

// HasMissingValue returns a boolean if a field has been set.
func (o *SortField) HasMissingValue() bool {
	if o != nil && o.MissingValue != nil {
		return true
	}

	return false
}

// SetMissingValue gets a reference to the given map[string]interface{} and assigns it to the MissingValue field.
func (o *SortField) SetMissingValue(v map[string]interface{}) {
	o.MissingValue = v
}

// GetReverse returns the Reverse field value if set, zero value otherwise.
func (o *SortField) GetReverse() bool {
	if o == nil || o.Reverse == nil {
		var ret bool
		return ret
	}
	return *o.Reverse
}

// GetReverseOk returns a tuple with the Reverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortField) GetReverseOk() (*bool, bool) {
	if o == nil || o.Reverse == nil {
		return nil, false
	}
	return o.Reverse, true
}

// HasReverse returns a boolean if a field has been set.
func (o *SortField) HasReverse() bool {
	if o != nil && o.Reverse != nil {
		return true
	}

	return false
}

// SetReverse gets a reference to the given bool and assigns it to the Reverse field.
func (o *SortField) SetReverse(v bool) {
	o.Reverse = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SortField) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortField) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SortField) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SortField) SetType(v string) {
	o.Type = &v
}

func (o SortField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BytesComparator != nil {
		toSerialize["bytesComparator"] = o.BytesComparator
	}
	if o.CanUsePoints != nil {
		toSerialize["canUsePoints"] = o.CanUsePoints
	}
	if o.ComparatorSource != nil {
		toSerialize["comparatorSource"] = o.ComparatorSource
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.IndexSorter != nil {
		toSerialize["indexSorter"] = o.IndexSorter
	}
	if o.MissingValue != nil {
		toSerialize["missingValue"] = o.MissingValue
	}
	if o.Reverse != nil {
		toSerialize["reverse"] = o.Reverse
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSortField struct {
	value *SortField
	isSet bool
}

func (v NullableSortField) Get() *SortField {
	return v.value
}

func (v *NullableSortField) Set(val *SortField) {
	v.value = val
	v.isSet = true
}

func (v NullableSortField) IsSet() bool {
	return v.isSet
}

func (v *NullableSortField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortField(val *SortField) *NullableSortField {
	return &NullableSortField{value: val, isSet: true}
}

func (v NullableSortField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
