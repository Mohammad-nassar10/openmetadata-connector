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

// Column struct for Column
type Column struct {
	ArrayDataType      *string        `json:"arrayDataType,omitempty"`
	Children           []Column       `json:"children,omitempty"`
	ColumnTests        []ColumnTest   `json:"columnTests,omitempty"`
	Constraint         *string        `json:"constraint,omitempty"`
	CustomMetrics      []CustomMetric `json:"customMetrics,omitempty"`
	DataLength         *int32         `json:"dataLength,omitempty"`
	DataType           string         `json:"dataType"`
	DataTypeDisplay    *string        `json:"dataTypeDisplay,omitempty"`
	Description        *string        `json:"description,omitempty"`
	DisplayName        *string        `json:"displayName,omitempty"`
	FullyQualifiedName *string        `json:"fullyQualifiedName,omitempty"`
	JsonSchema         *string        `json:"jsonSchema,omitempty"`
	Name               string         `json:"name"`
	OrdinalPosition    *int32         `json:"ordinalPosition,omitempty"`
	Precision          *int32         `json:"precision,omitempty"`
	Scale              *int32         `json:"scale,omitempty"`
	Tags               []TagLabel     `json:"tags,omitempty"`
}

// NewColumn instantiates a new Column object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColumn(dataType string, name string) *Column {
	this := Column{}
	this.DataType = dataType
	this.Name = name
	return &this
}

// NewColumnWithDefaults instantiates a new Column object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColumnWithDefaults() *Column {
	this := Column{}
	return &this
}

// GetArrayDataType returns the ArrayDataType field value if set, zero value otherwise.
func (o *Column) GetArrayDataType() string {
	if o == nil || o.ArrayDataType == nil {
		var ret string
		return ret
	}
	return *o.ArrayDataType
}

// GetArrayDataTypeOk returns a tuple with the ArrayDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetArrayDataTypeOk() (*string, bool) {
	if o == nil || o.ArrayDataType == nil {
		return nil, false
	}
	return o.ArrayDataType, true
}

// HasArrayDataType returns a boolean if a field has been set.
func (o *Column) HasArrayDataType() bool {
	if o != nil && o.ArrayDataType != nil {
		return true
	}

	return false
}

// SetArrayDataType gets a reference to the given string and assigns it to the ArrayDataType field.
func (o *Column) SetArrayDataType(v string) {
	o.ArrayDataType = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *Column) GetChildren() []Column {
	if o == nil || o.Children == nil {
		var ret []Column
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetChildrenOk() ([]Column, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Column) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []Column and assigns it to the Children field.
func (o *Column) SetChildren(v []Column) {
	o.Children = v
}

// GetColumnTests returns the ColumnTests field value if set, zero value otherwise.
func (o *Column) GetColumnTests() []ColumnTest {
	if o == nil || o.ColumnTests == nil {
		var ret []ColumnTest
		return ret
	}
	return o.ColumnTests
}

// GetColumnTestsOk returns a tuple with the ColumnTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetColumnTestsOk() ([]ColumnTest, bool) {
	if o == nil || o.ColumnTests == nil {
		return nil, false
	}
	return o.ColumnTests, true
}

// HasColumnTests returns a boolean if a field has been set.
func (o *Column) HasColumnTests() bool {
	if o != nil && o.ColumnTests != nil {
		return true
	}

	return false
}

// SetColumnTests gets a reference to the given []ColumnTest and assigns it to the ColumnTests field.
func (o *Column) SetColumnTests(v []ColumnTest) {
	o.ColumnTests = v
}

// GetConstraint returns the Constraint field value if set, zero value otherwise.
func (o *Column) GetConstraint() string {
	if o == nil || o.Constraint == nil {
		var ret string
		return ret
	}
	return *o.Constraint
}

// GetConstraintOk returns a tuple with the Constraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetConstraintOk() (*string, bool) {
	if o == nil || o.Constraint == nil {
		return nil, false
	}
	return o.Constraint, true
}

// HasConstraint returns a boolean if a field has been set.
func (o *Column) HasConstraint() bool {
	if o != nil && o.Constraint != nil {
		return true
	}

	return false
}

// SetConstraint gets a reference to the given string and assigns it to the Constraint field.
func (o *Column) SetConstraint(v string) {
	o.Constraint = &v
}

// GetCustomMetrics returns the CustomMetrics field value if set, zero value otherwise.
func (o *Column) GetCustomMetrics() []CustomMetric {
	if o == nil || o.CustomMetrics == nil {
		var ret []CustomMetric
		return ret
	}
	return o.CustomMetrics
}

// GetCustomMetricsOk returns a tuple with the CustomMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetCustomMetricsOk() ([]CustomMetric, bool) {
	if o == nil || o.CustomMetrics == nil {
		return nil, false
	}
	return o.CustomMetrics, true
}

// HasCustomMetrics returns a boolean if a field has been set.
func (o *Column) HasCustomMetrics() bool {
	if o != nil && o.CustomMetrics != nil {
		return true
	}

	return false
}

// SetCustomMetrics gets a reference to the given []CustomMetric and assigns it to the CustomMetrics field.
func (o *Column) SetCustomMetrics(v []CustomMetric) {
	o.CustomMetrics = v
}

// GetDataLength returns the DataLength field value if set, zero value otherwise.
func (o *Column) GetDataLength() int32 {
	if o == nil || o.DataLength == nil {
		var ret int32
		return ret
	}
	return *o.DataLength
}

// GetDataLengthOk returns a tuple with the DataLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetDataLengthOk() (*int32, bool) {
	if o == nil || o.DataLength == nil {
		return nil, false
	}
	return o.DataLength, true
}

// HasDataLength returns a boolean if a field has been set.
func (o *Column) HasDataLength() bool {
	if o != nil && o.DataLength != nil {
		return true
	}

	return false
}

// SetDataLength gets a reference to the given int32 and assigns it to the DataLength field.
func (o *Column) SetDataLength(v int32) {
	o.DataLength = &v
}

// GetDataType returns the DataType field value
func (o *Column) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *Column) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *Column) SetDataType(v string) {
	o.DataType = v
}

// GetDataTypeDisplay returns the DataTypeDisplay field value if set, zero value otherwise.
func (o *Column) GetDataTypeDisplay() string {
	if o == nil || o.DataTypeDisplay == nil {
		var ret string
		return ret
	}
	return *o.DataTypeDisplay
}

// GetDataTypeDisplayOk returns a tuple with the DataTypeDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetDataTypeDisplayOk() (*string, bool) {
	if o == nil || o.DataTypeDisplay == nil {
		return nil, false
	}
	return o.DataTypeDisplay, true
}

// HasDataTypeDisplay returns a boolean if a field has been set.
func (o *Column) HasDataTypeDisplay() bool {
	if o != nil && o.DataTypeDisplay != nil {
		return true
	}

	return false
}

// SetDataTypeDisplay gets a reference to the given string and assigns it to the DataTypeDisplay field.
func (o *Column) SetDataTypeDisplay(v string) {
	o.DataTypeDisplay = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Column) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Column) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Column) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Column) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Column) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Column) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFullyQualifiedName returns the FullyQualifiedName field value if set, zero value otherwise.
func (o *Column) GetFullyQualifiedName() string {
	if o == nil || o.FullyQualifiedName == nil {
		var ret string
		return ret
	}
	return *o.FullyQualifiedName
}

// GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetFullyQualifiedNameOk() (*string, bool) {
	if o == nil || o.FullyQualifiedName == nil {
		return nil, false
	}
	return o.FullyQualifiedName, true
}

// HasFullyQualifiedName returns a boolean if a field has been set.
func (o *Column) HasFullyQualifiedName() bool {
	if o != nil && o.FullyQualifiedName != nil {
		return true
	}

	return false
}

// SetFullyQualifiedName gets a reference to the given string and assigns it to the FullyQualifiedName field.
func (o *Column) SetFullyQualifiedName(v string) {
	o.FullyQualifiedName = &v
}

// GetJsonSchema returns the JsonSchema field value if set, zero value otherwise.
func (o *Column) GetJsonSchema() string {
	if o == nil || o.JsonSchema == nil {
		var ret string
		return ret
	}
	return *o.JsonSchema
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetJsonSchemaOk() (*string, bool) {
	if o == nil || o.JsonSchema == nil {
		return nil, false
	}
	return o.JsonSchema, true
}

// HasJsonSchema returns a boolean if a field has been set.
func (o *Column) HasJsonSchema() bool {
	if o != nil && o.JsonSchema != nil {
		return true
	}

	return false
}

// SetJsonSchema gets a reference to the given string and assigns it to the JsonSchema field.
func (o *Column) SetJsonSchema(v string) {
	o.JsonSchema = &v
}

// GetName returns the Name field value
func (o *Column) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Column) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Column) SetName(v string) {
	o.Name = v
}

// GetOrdinalPosition returns the OrdinalPosition field value if set, zero value otherwise.
func (o *Column) GetOrdinalPosition() int32 {
	if o == nil || o.OrdinalPosition == nil {
		var ret int32
		return ret
	}
	return *o.OrdinalPosition
}

// GetOrdinalPositionOk returns a tuple with the OrdinalPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetOrdinalPositionOk() (*int32, bool) {
	if o == nil || o.OrdinalPosition == nil {
		return nil, false
	}
	return o.OrdinalPosition, true
}

// HasOrdinalPosition returns a boolean if a field has been set.
func (o *Column) HasOrdinalPosition() bool {
	if o != nil && o.OrdinalPosition != nil {
		return true
	}

	return false
}

// SetOrdinalPosition gets a reference to the given int32 and assigns it to the OrdinalPosition field.
func (o *Column) SetOrdinalPosition(v int32) {
	o.OrdinalPosition = &v
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *Column) GetPrecision() int32 {
	if o == nil || o.Precision == nil {
		var ret int32
		return ret
	}
	return *o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetPrecisionOk() (*int32, bool) {
	if o == nil || o.Precision == nil {
		return nil, false
	}
	return o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *Column) HasPrecision() bool {
	if o != nil && o.Precision != nil {
		return true
	}

	return false
}

// SetPrecision gets a reference to the given int32 and assigns it to the Precision field.
func (o *Column) SetPrecision(v int32) {
	o.Precision = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *Column) GetScale() int32 {
	if o == nil || o.Scale == nil {
		var ret int32
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetScaleOk() (*int32, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *Column) HasScale() bool {
	if o != nil && o.Scale != nil {
		return true
	}

	return false
}

// SetScale gets a reference to the given int32 and assigns it to the Scale field.
func (o *Column) SetScale(v int32) {
	o.Scale = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Column) GetTags() []TagLabel {
	if o == nil || o.Tags == nil {
		var ret []TagLabel
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetTagsOk() ([]TagLabel, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Column) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagLabel and assigns it to the Tags field.
func (o *Column) SetTags(v []TagLabel) {
	o.Tags = v
}

func (o Column) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArrayDataType != nil {
		toSerialize["arrayDataType"] = o.ArrayDataType
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.ColumnTests != nil {
		toSerialize["columnTests"] = o.ColumnTests
	}
	if o.Constraint != nil {
		toSerialize["constraint"] = o.Constraint
	}
	if o.CustomMetrics != nil {
		toSerialize["customMetrics"] = o.CustomMetrics
	}
	if o.DataLength != nil {
		toSerialize["dataLength"] = o.DataLength
	}
	if true {
		toSerialize["dataType"] = o.DataType
	}
	if o.DataTypeDisplay != nil {
		toSerialize["dataTypeDisplay"] = o.DataTypeDisplay
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.FullyQualifiedName != nil {
		toSerialize["fullyQualifiedName"] = o.FullyQualifiedName
	}
	if o.JsonSchema != nil {
		toSerialize["jsonSchema"] = o.JsonSchema
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OrdinalPosition != nil {
		toSerialize["ordinalPosition"] = o.OrdinalPosition
	}
	if o.Precision != nil {
		toSerialize["precision"] = o.Precision
	}
	if o.Scale != nil {
		toSerialize["scale"] = o.Scale
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableColumn struct {
	value *Column
	isSet bool
}

func (v NullableColumn) Get() *Column {
	return v.value
}

func (v *NullableColumn) Set(val *Column) {
	v.value = val
	v.isSet = true
}

func (v NullableColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColumn(val *Column) *NullableColumn {
	return &NullableColumn{value: val, isSet: true}
}

func (v NullableColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
