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

// ColumnProfile struct for ColumnProfile
type ColumnProfile struct {
	CustomMetricsProfile []CustomMetricProfile `json:"customMetricsProfile,omitempty"`
	DistinctCount        *float64              `json:"distinctCount,omitempty"`
	DistinctProportion   *float64              `json:"distinctProportion,omitempty"`
	DuplicateCount       *float64              `json:"duplicateCount,omitempty"`
	Histogram            *Histogram            `json:"histogram,omitempty"`
	Max                  *float64              `json:"max,omitempty"`
	MaxLength            *float64              `json:"maxLength,omitempty"`
	Mean                 *float64              `json:"mean,omitempty"`
	Min                  *float64              `json:"min,omitempty"`
	MinLength            *float64              `json:"minLength,omitempty"`
	MissingCount         *float64              `json:"missingCount,omitempty"`
	MissingPercentage    *float64              `json:"missingPercentage,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	NullCount            *float64              `json:"nullCount,omitempty"`
	NullProportion       *float64              `json:"nullProportion,omitempty"`
	Stddev               *float64              `json:"stddev,omitempty"`
	Sum                  *float64              `json:"sum,omitempty"`
	UniqueCount          *float64              `json:"uniqueCount,omitempty"`
	UniqueProportion     *float64              `json:"uniqueProportion,omitempty"`
	ValidCount           *float64              `json:"validCount,omitempty"`
	ValuesCount          *float64              `json:"valuesCount,omitempty"`
	ValuesPercentage     *float64              `json:"valuesPercentage,omitempty"`
	Variance             *float64              `json:"variance,omitempty"`
}

// NewColumnProfile instantiates a new ColumnProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColumnProfile() *ColumnProfile {
	this := ColumnProfile{}
	return &this
}

// NewColumnProfileWithDefaults instantiates a new ColumnProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColumnProfileWithDefaults() *ColumnProfile {
	this := ColumnProfile{}
	return &this
}

// GetCustomMetricsProfile returns the CustomMetricsProfile field value if set, zero value otherwise.
func (o *ColumnProfile) GetCustomMetricsProfile() []CustomMetricProfile {
	if o == nil || o.CustomMetricsProfile == nil {
		var ret []CustomMetricProfile
		return ret
	}
	return o.CustomMetricsProfile
}

// GetCustomMetricsProfileOk returns a tuple with the CustomMetricsProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetCustomMetricsProfileOk() ([]CustomMetricProfile, bool) {
	if o == nil || o.CustomMetricsProfile == nil {
		return nil, false
	}
	return o.CustomMetricsProfile, true
}

// HasCustomMetricsProfile returns a boolean if a field has been set.
func (o *ColumnProfile) HasCustomMetricsProfile() bool {
	if o != nil && o.CustomMetricsProfile != nil {
		return true
	}

	return false
}

// SetCustomMetricsProfile gets a reference to the given []CustomMetricProfile and assigns it to the CustomMetricsProfile field.
func (o *ColumnProfile) SetCustomMetricsProfile(v []CustomMetricProfile) {
	o.CustomMetricsProfile = v
}

// GetDistinctCount returns the DistinctCount field value if set, zero value otherwise.
func (o *ColumnProfile) GetDistinctCount() float64 {
	if o == nil || o.DistinctCount == nil {
		var ret float64
		return ret
	}
	return *o.DistinctCount
}

// GetDistinctCountOk returns a tuple with the DistinctCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetDistinctCountOk() (*float64, bool) {
	if o == nil || o.DistinctCount == nil {
		return nil, false
	}
	return o.DistinctCount, true
}

// HasDistinctCount returns a boolean if a field has been set.
func (o *ColumnProfile) HasDistinctCount() bool {
	if o != nil && o.DistinctCount != nil {
		return true
	}

	return false
}

// SetDistinctCount gets a reference to the given float64 and assigns it to the DistinctCount field.
func (o *ColumnProfile) SetDistinctCount(v float64) {
	o.DistinctCount = &v
}

// GetDistinctProportion returns the DistinctProportion field value if set, zero value otherwise.
func (o *ColumnProfile) GetDistinctProportion() float64 {
	if o == nil || o.DistinctProportion == nil {
		var ret float64
		return ret
	}
	return *o.DistinctProportion
}

// GetDistinctProportionOk returns a tuple with the DistinctProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetDistinctProportionOk() (*float64, bool) {
	if o == nil || o.DistinctProportion == nil {
		return nil, false
	}
	return o.DistinctProportion, true
}

// HasDistinctProportion returns a boolean if a field has been set.
func (o *ColumnProfile) HasDistinctProportion() bool {
	if o != nil && o.DistinctProportion != nil {
		return true
	}

	return false
}

// SetDistinctProportion gets a reference to the given float64 and assigns it to the DistinctProportion field.
func (o *ColumnProfile) SetDistinctProportion(v float64) {
	o.DistinctProportion = &v
}

// GetDuplicateCount returns the DuplicateCount field value if set, zero value otherwise.
func (o *ColumnProfile) GetDuplicateCount() float64 {
	if o == nil || o.DuplicateCount == nil {
		var ret float64
		return ret
	}
	return *o.DuplicateCount
}

// GetDuplicateCountOk returns a tuple with the DuplicateCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetDuplicateCountOk() (*float64, bool) {
	if o == nil || o.DuplicateCount == nil {
		return nil, false
	}
	return o.DuplicateCount, true
}

// HasDuplicateCount returns a boolean if a field has been set.
func (o *ColumnProfile) HasDuplicateCount() bool {
	if o != nil && o.DuplicateCount != nil {
		return true
	}

	return false
}

// SetDuplicateCount gets a reference to the given float64 and assigns it to the DuplicateCount field.
func (o *ColumnProfile) SetDuplicateCount(v float64) {
	o.DuplicateCount = &v
}

// GetHistogram returns the Histogram field value if set, zero value otherwise.
func (o *ColumnProfile) GetHistogram() Histogram {
	if o == nil || o.Histogram == nil {
		var ret Histogram
		return ret
	}
	return *o.Histogram
}

// GetHistogramOk returns a tuple with the Histogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetHistogramOk() (*Histogram, bool) {
	if o == nil || o.Histogram == nil {
		return nil, false
	}
	return o.Histogram, true
}

// HasHistogram returns a boolean if a field has been set.
func (o *ColumnProfile) HasHistogram() bool {
	if o != nil && o.Histogram != nil {
		return true
	}

	return false
}

// SetHistogram gets a reference to the given Histogram and assigns it to the Histogram field.
func (o *ColumnProfile) SetHistogram(v Histogram) {
	o.Histogram = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ColumnProfile) GetMax() float64 {
	if o == nil || o.Max == nil {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetMaxOk() (*float64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ColumnProfile) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *ColumnProfile) SetMax(v float64) {
	o.Max = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *ColumnProfile) GetMaxLength() float64 {
	if o == nil || o.MaxLength == nil {
		var ret float64
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetMaxLengthOk() (*float64, bool) {
	if o == nil || o.MaxLength == nil {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *ColumnProfile) HasMaxLength() bool {
	if o != nil && o.MaxLength != nil {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given float64 and assigns it to the MaxLength field.
func (o *ColumnProfile) SetMaxLength(v float64) {
	o.MaxLength = &v
}

// GetMean returns the Mean field value if set, zero value otherwise.
func (o *ColumnProfile) GetMean() float64 {
	if o == nil || o.Mean == nil {
		var ret float64
		return ret
	}
	return *o.Mean
}

// GetMeanOk returns a tuple with the Mean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetMeanOk() (*float64, bool) {
	if o == nil || o.Mean == nil {
		return nil, false
	}
	return o.Mean, true
}

// HasMean returns a boolean if a field has been set.
func (o *ColumnProfile) HasMean() bool {
	if o != nil && o.Mean != nil {
		return true
	}

	return false
}

// SetMean gets a reference to the given float64 and assigns it to the Mean field.
func (o *ColumnProfile) SetMean(v float64) {
	o.Mean = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ColumnProfile) GetMin() float64 {
	if o == nil || o.Min == nil {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetMinOk() (*float64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ColumnProfile) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *ColumnProfile) SetMin(v float64) {
	o.Min = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *ColumnProfile) GetMinLength() float64 {
	if o == nil || o.MinLength == nil {
		var ret float64
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetMinLengthOk() (*float64, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *ColumnProfile) HasMinLength() bool {
	if o != nil && o.MinLength != nil {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given float64 and assigns it to the MinLength field.
func (o *ColumnProfile) SetMinLength(v float64) {
	o.MinLength = &v
}

// GetMissingCount returns the MissingCount field value if set, zero value otherwise.
func (o *ColumnProfile) GetMissingCount() float64 {
	if o == nil || o.MissingCount == nil {
		var ret float64
		return ret
	}
	return *o.MissingCount
}

// GetMissingCountOk returns a tuple with the MissingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetMissingCountOk() (*float64, bool) {
	if o == nil || o.MissingCount == nil {
		return nil, false
	}
	return o.MissingCount, true
}

// HasMissingCount returns a boolean if a field has been set.
func (o *ColumnProfile) HasMissingCount() bool {
	if o != nil && o.MissingCount != nil {
		return true
	}

	return false
}

// SetMissingCount gets a reference to the given float64 and assigns it to the MissingCount field.
func (o *ColumnProfile) SetMissingCount(v float64) {
	o.MissingCount = &v
}

// GetMissingPercentage returns the MissingPercentage field value if set, zero value otherwise.
func (o *ColumnProfile) GetMissingPercentage() float64 {
	if o == nil || o.MissingPercentage == nil {
		var ret float64
		return ret
	}
	return *o.MissingPercentage
}

// GetMissingPercentageOk returns a tuple with the MissingPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetMissingPercentageOk() (*float64, bool) {
	if o == nil || o.MissingPercentage == nil {
		return nil, false
	}
	return o.MissingPercentage, true
}

// HasMissingPercentage returns a boolean if a field has been set.
func (o *ColumnProfile) HasMissingPercentage() bool {
	if o != nil && o.MissingPercentage != nil {
		return true
	}

	return false
}

// SetMissingPercentage gets a reference to the given float64 and assigns it to the MissingPercentage field.
func (o *ColumnProfile) SetMissingPercentage(v float64) {
	o.MissingPercentage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ColumnProfile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ColumnProfile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ColumnProfile) SetName(v string) {
	o.Name = &v
}

// GetNullCount returns the NullCount field value if set, zero value otherwise.
func (o *ColumnProfile) GetNullCount() float64 {
	if o == nil || o.NullCount == nil {
		var ret float64
		return ret
	}
	return *o.NullCount
}

// GetNullCountOk returns a tuple with the NullCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetNullCountOk() (*float64, bool) {
	if o == nil || o.NullCount == nil {
		return nil, false
	}
	return o.NullCount, true
}

// HasNullCount returns a boolean if a field has been set.
func (o *ColumnProfile) HasNullCount() bool {
	if o != nil && o.NullCount != nil {
		return true
	}

	return false
}

// SetNullCount gets a reference to the given float64 and assigns it to the NullCount field.
func (o *ColumnProfile) SetNullCount(v float64) {
	o.NullCount = &v
}

// GetNullProportion returns the NullProportion field value if set, zero value otherwise.
func (o *ColumnProfile) GetNullProportion() float64 {
	if o == nil || o.NullProportion == nil {
		var ret float64
		return ret
	}
	return *o.NullProportion
}

// GetNullProportionOk returns a tuple with the NullProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetNullProportionOk() (*float64, bool) {
	if o == nil || o.NullProportion == nil {
		return nil, false
	}
	return o.NullProportion, true
}

// HasNullProportion returns a boolean if a field has been set.
func (o *ColumnProfile) HasNullProportion() bool {
	if o != nil && o.NullProportion != nil {
		return true
	}

	return false
}

// SetNullProportion gets a reference to the given float64 and assigns it to the NullProportion field.
func (o *ColumnProfile) SetNullProportion(v float64) {
	o.NullProportion = &v
}

// GetStddev returns the Stddev field value if set, zero value otherwise.
func (o *ColumnProfile) GetStddev() float64 {
	if o == nil || o.Stddev == nil {
		var ret float64
		return ret
	}
	return *o.Stddev
}

// GetStddevOk returns a tuple with the Stddev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetStddevOk() (*float64, bool) {
	if o == nil || o.Stddev == nil {
		return nil, false
	}
	return o.Stddev, true
}

// HasStddev returns a boolean if a field has been set.
func (o *ColumnProfile) HasStddev() bool {
	if o != nil && o.Stddev != nil {
		return true
	}

	return false
}

// SetStddev gets a reference to the given float64 and assigns it to the Stddev field.
func (o *ColumnProfile) SetStddev(v float64) {
	o.Stddev = &v
}

// GetSum returns the Sum field value if set, zero value otherwise.
func (o *ColumnProfile) GetSum() float64 {
	if o == nil || o.Sum == nil {
		var ret float64
		return ret
	}
	return *o.Sum
}

// GetSumOk returns a tuple with the Sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetSumOk() (*float64, bool) {
	if o == nil || o.Sum == nil {
		return nil, false
	}
	return o.Sum, true
}

// HasSum returns a boolean if a field has been set.
func (o *ColumnProfile) HasSum() bool {
	if o != nil && o.Sum != nil {
		return true
	}

	return false
}

// SetSum gets a reference to the given float64 and assigns it to the Sum field.
func (o *ColumnProfile) SetSum(v float64) {
	o.Sum = &v
}

// GetUniqueCount returns the UniqueCount field value if set, zero value otherwise.
func (o *ColumnProfile) GetUniqueCount() float64 {
	if o == nil || o.UniqueCount == nil {
		var ret float64
		return ret
	}
	return *o.UniqueCount
}

// GetUniqueCountOk returns a tuple with the UniqueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetUniqueCountOk() (*float64, bool) {
	if o == nil || o.UniqueCount == nil {
		return nil, false
	}
	return o.UniqueCount, true
}

// HasUniqueCount returns a boolean if a field has been set.
func (o *ColumnProfile) HasUniqueCount() bool {
	if o != nil && o.UniqueCount != nil {
		return true
	}

	return false
}

// SetUniqueCount gets a reference to the given float64 and assigns it to the UniqueCount field.
func (o *ColumnProfile) SetUniqueCount(v float64) {
	o.UniqueCount = &v
}

// GetUniqueProportion returns the UniqueProportion field value if set, zero value otherwise.
func (o *ColumnProfile) GetUniqueProportion() float64 {
	if o == nil || o.UniqueProportion == nil {
		var ret float64
		return ret
	}
	return *o.UniqueProportion
}

// GetUniqueProportionOk returns a tuple with the UniqueProportion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetUniqueProportionOk() (*float64, bool) {
	if o == nil || o.UniqueProportion == nil {
		return nil, false
	}
	return o.UniqueProportion, true
}

// HasUniqueProportion returns a boolean if a field has been set.
func (o *ColumnProfile) HasUniqueProportion() bool {
	if o != nil && o.UniqueProportion != nil {
		return true
	}

	return false
}

// SetUniqueProportion gets a reference to the given float64 and assigns it to the UniqueProportion field.
func (o *ColumnProfile) SetUniqueProportion(v float64) {
	o.UniqueProportion = &v
}

// GetValidCount returns the ValidCount field value if set, zero value otherwise.
func (o *ColumnProfile) GetValidCount() float64 {
	if o == nil || o.ValidCount == nil {
		var ret float64
		return ret
	}
	return *o.ValidCount
}

// GetValidCountOk returns a tuple with the ValidCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetValidCountOk() (*float64, bool) {
	if o == nil || o.ValidCount == nil {
		return nil, false
	}
	return o.ValidCount, true
}

// HasValidCount returns a boolean if a field has been set.
func (o *ColumnProfile) HasValidCount() bool {
	if o != nil && o.ValidCount != nil {
		return true
	}

	return false
}

// SetValidCount gets a reference to the given float64 and assigns it to the ValidCount field.
func (o *ColumnProfile) SetValidCount(v float64) {
	o.ValidCount = &v
}

// GetValuesCount returns the ValuesCount field value if set, zero value otherwise.
func (o *ColumnProfile) GetValuesCount() float64 {
	if o == nil || o.ValuesCount == nil {
		var ret float64
		return ret
	}
	return *o.ValuesCount
}

// GetValuesCountOk returns a tuple with the ValuesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetValuesCountOk() (*float64, bool) {
	if o == nil || o.ValuesCount == nil {
		return nil, false
	}
	return o.ValuesCount, true
}

// HasValuesCount returns a boolean if a field has been set.
func (o *ColumnProfile) HasValuesCount() bool {
	if o != nil && o.ValuesCount != nil {
		return true
	}

	return false
}

// SetValuesCount gets a reference to the given float64 and assigns it to the ValuesCount field.
func (o *ColumnProfile) SetValuesCount(v float64) {
	o.ValuesCount = &v
}

// GetValuesPercentage returns the ValuesPercentage field value if set, zero value otherwise.
func (o *ColumnProfile) GetValuesPercentage() float64 {
	if o == nil || o.ValuesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ValuesPercentage
}

// GetValuesPercentageOk returns a tuple with the ValuesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetValuesPercentageOk() (*float64, bool) {
	if o == nil || o.ValuesPercentage == nil {
		return nil, false
	}
	return o.ValuesPercentage, true
}

// HasValuesPercentage returns a boolean if a field has been set.
func (o *ColumnProfile) HasValuesPercentage() bool {
	if o != nil && o.ValuesPercentage != nil {
		return true
	}

	return false
}

// SetValuesPercentage gets a reference to the given float64 and assigns it to the ValuesPercentage field.
func (o *ColumnProfile) SetValuesPercentage(v float64) {
	o.ValuesPercentage = &v
}

// GetVariance returns the Variance field value if set, zero value otherwise.
func (o *ColumnProfile) GetVariance() float64 {
	if o == nil || o.Variance == nil {
		var ret float64
		return ret
	}
	return *o.Variance
}

// GetVarianceOk returns a tuple with the Variance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnProfile) GetVarianceOk() (*float64, bool) {
	if o == nil || o.Variance == nil {
		return nil, false
	}
	return o.Variance, true
}

// HasVariance returns a boolean if a field has been set.
func (o *ColumnProfile) HasVariance() bool {
	if o != nil && o.Variance != nil {
		return true
	}

	return false
}

// SetVariance gets a reference to the given float64 and assigns it to the Variance field.
func (o *ColumnProfile) SetVariance(v float64) {
	o.Variance = &v
}

func (o ColumnProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomMetricsProfile != nil {
		toSerialize["customMetricsProfile"] = o.CustomMetricsProfile
	}
	if o.DistinctCount != nil {
		toSerialize["distinctCount"] = o.DistinctCount
	}
	if o.DistinctProportion != nil {
		toSerialize["distinctProportion"] = o.DistinctProportion
	}
	if o.DuplicateCount != nil {
		toSerialize["duplicateCount"] = o.DuplicateCount
	}
	if o.Histogram != nil {
		toSerialize["histogram"] = o.Histogram
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.MaxLength != nil {
		toSerialize["maxLength"] = o.MaxLength
	}
	if o.Mean != nil {
		toSerialize["mean"] = o.Mean
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.MinLength != nil {
		toSerialize["minLength"] = o.MinLength
	}
	if o.MissingCount != nil {
		toSerialize["missingCount"] = o.MissingCount
	}
	if o.MissingPercentage != nil {
		toSerialize["missingPercentage"] = o.MissingPercentage
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NullCount != nil {
		toSerialize["nullCount"] = o.NullCount
	}
	if o.NullProportion != nil {
		toSerialize["nullProportion"] = o.NullProportion
	}
	if o.Stddev != nil {
		toSerialize["stddev"] = o.Stddev
	}
	if o.Sum != nil {
		toSerialize["sum"] = o.Sum
	}
	if o.UniqueCount != nil {
		toSerialize["uniqueCount"] = o.UniqueCount
	}
	if o.UniqueProportion != nil {
		toSerialize["uniqueProportion"] = o.UniqueProportion
	}
	if o.ValidCount != nil {
		toSerialize["validCount"] = o.ValidCount
	}
	if o.ValuesCount != nil {
		toSerialize["valuesCount"] = o.ValuesCount
	}
	if o.ValuesPercentage != nil {
		toSerialize["valuesPercentage"] = o.ValuesPercentage
	}
	if o.Variance != nil {
		toSerialize["variance"] = o.Variance
	}
	return json.Marshal(toSerialize)
}

type NullableColumnProfile struct {
	value *ColumnProfile
	isSet bool
}

func (v NullableColumnProfile) Get() *ColumnProfile {
	return v.value
}

func (v *NullableColumnProfile) Set(val *ColumnProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableColumnProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableColumnProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColumnProfile(val *ColumnProfile) *NullableColumnProfile {
	return &NullableColumnProfile{value: val, isSet: true}
}

func (v NullableColumnProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColumnProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
