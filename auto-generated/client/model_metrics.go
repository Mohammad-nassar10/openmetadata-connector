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

// Metrics struct for Metrics
type Metrics struct {
	ChangeDescription  *ChangeDescription     `json:"changeDescription,omitempty"`
	Deleted            *bool                  `json:"deleted,omitempty"`
	Description        *string                `json:"description,omitempty"`
	DisplayName        *string                `json:"displayName,omitempty"`
	Extension          map[string]interface{} `json:"extension,omitempty"`
	Followers          []EntityReference      `json:"followers,omitempty"`
	FullyQualifiedName *string                `json:"fullyQualifiedName,omitempty"`
	Href               *string                `json:"href,omitempty"`
	Id                 string                 `json:"id"`
	Name               string                 `json:"name"`
	Owner              *EntityReference       `json:"owner,omitempty"`
	Service            EntityReference        `json:"service"`
	Tags               []TagLabel             `json:"tags,omitempty"`
	UpdatedAt          *int64                 `json:"updatedAt,omitempty"`
	UpdatedBy          *string                `json:"updatedBy,omitempty"`
	UsageSummary       *UsageDetails          `json:"usageSummary,omitempty"`
	Version            *float64               `json:"version,omitempty"`
}

// NewMetrics instantiates a new Metrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetrics(id string, name string, service EntityReference) *Metrics {
	this := Metrics{}
	this.Id = id
	this.Name = name
	this.Service = service
	return &this
}

// NewMetricsWithDefaults instantiates a new Metrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsWithDefaults() *Metrics {
	this := Metrics{}
	return &this
}

// GetChangeDescription returns the ChangeDescription field value if set, zero value otherwise.
func (o *Metrics) GetChangeDescription() ChangeDescription {
	if o == nil || o.ChangeDescription == nil {
		var ret ChangeDescription
		return ret
	}
	return *o.ChangeDescription
}

// GetChangeDescriptionOk returns a tuple with the ChangeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetChangeDescriptionOk() (*ChangeDescription, bool) {
	if o == nil || o.ChangeDescription == nil {
		return nil, false
	}
	return o.ChangeDescription, true
}

// HasChangeDescription returns a boolean if a field has been set.
func (o *Metrics) HasChangeDescription() bool {
	if o != nil && o.ChangeDescription != nil {
		return true
	}

	return false
}

// SetChangeDescription gets a reference to the given ChangeDescription and assigns it to the ChangeDescription field.
func (o *Metrics) SetChangeDescription(v ChangeDescription) {
	o.ChangeDescription = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Metrics) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Metrics) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Metrics) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Metrics) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Metrics) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Metrics) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Metrics) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Metrics) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Metrics) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Metrics) GetExtension() map[string]interface{} {
	if o == nil || o.Extension == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetExtensionOk() (map[string]interface{}, bool) {
	if o == nil || o.Extension == nil {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Metrics) HasExtension() bool {
	if o != nil && o.Extension != nil {
		return true
	}

	return false
}

// SetExtension gets a reference to the given map[string]interface{} and assigns it to the Extension field.
func (o *Metrics) SetExtension(v map[string]interface{}) {
	o.Extension = v
}

// GetFollowers returns the Followers field value if set, zero value otherwise.
func (o *Metrics) GetFollowers() []EntityReference {
	if o == nil || o.Followers == nil {
		var ret []EntityReference
		return ret
	}
	return o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetFollowersOk() ([]EntityReference, bool) {
	if o == nil || o.Followers == nil {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *Metrics) HasFollowers() bool {
	if o != nil && o.Followers != nil {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given []EntityReference and assigns it to the Followers field.
func (o *Metrics) SetFollowers(v []EntityReference) {
	o.Followers = v
}

// GetFullyQualifiedName returns the FullyQualifiedName field value if set, zero value otherwise.
func (o *Metrics) GetFullyQualifiedName() string {
	if o == nil || o.FullyQualifiedName == nil {
		var ret string
		return ret
	}
	return *o.FullyQualifiedName
}

// GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetFullyQualifiedNameOk() (*string, bool) {
	if o == nil || o.FullyQualifiedName == nil {
		return nil, false
	}
	return o.FullyQualifiedName, true
}

// HasFullyQualifiedName returns a boolean if a field has been set.
func (o *Metrics) HasFullyQualifiedName() bool {
	if o != nil && o.FullyQualifiedName != nil {
		return true
	}

	return false
}

// SetFullyQualifiedName gets a reference to the given string and assigns it to the FullyQualifiedName field.
func (o *Metrics) SetFullyQualifiedName(v string) {
	o.FullyQualifiedName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Metrics) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Metrics) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Metrics) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *Metrics) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Metrics) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Metrics) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Metrics) SetName(v string) {
	o.Name = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Metrics) GetOwner() EntityReference {
	if o == nil || o.Owner == nil {
		var ret EntityReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetOwnerOk() (*EntityReference, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Metrics) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given EntityReference and assigns it to the Owner field.
func (o *Metrics) SetOwner(v EntityReference) {
	o.Owner = &v
}

// GetService returns the Service field value
func (o *Metrics) GetService() EntityReference {
	if o == nil {
		var ret EntityReference
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *Metrics) GetServiceOk() (*EntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *Metrics) SetService(v EntityReference) {
	o.Service = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Metrics) GetTags() []TagLabel {
	if o == nil || o.Tags == nil {
		var ret []TagLabel
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetTagsOk() ([]TagLabel, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Metrics) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagLabel and assigns it to the Tags field.
func (o *Metrics) SetTags(v []TagLabel) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Metrics) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Metrics) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Metrics) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Metrics) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Metrics) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *Metrics) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUsageSummary returns the UsageSummary field value if set, zero value otherwise.
func (o *Metrics) GetUsageSummary() UsageDetails {
	if o == nil || o.UsageSummary == nil {
		var ret UsageDetails
		return ret
	}
	return *o.UsageSummary
}

// GetUsageSummaryOk returns a tuple with the UsageSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetUsageSummaryOk() (*UsageDetails, bool) {
	if o == nil || o.UsageSummary == nil {
		return nil, false
	}
	return o.UsageSummary, true
}

// HasUsageSummary returns a boolean if a field has been set.
func (o *Metrics) HasUsageSummary() bool {
	if o != nil && o.UsageSummary != nil {
		return true
	}

	return false
}

// SetUsageSummary gets a reference to the given UsageDetails and assigns it to the UsageSummary field.
func (o *Metrics) SetUsageSummary(v UsageDetails) {
	o.UsageSummary = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Metrics) GetVersion() float64 {
	if o == nil || o.Version == nil {
		var ret float64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metrics) GetVersionOk() (*float64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Metrics) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float64 and assigns it to the Version field.
func (o *Metrics) SetVersion(v float64) {
	o.Version = &v
}

func (o Metrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeDescription != nil {
		toSerialize["changeDescription"] = o.ChangeDescription
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Extension != nil {
		toSerialize["extension"] = o.Extension
	}
	if o.Followers != nil {
		toSerialize["followers"] = o.Followers
	}
	if o.FullyQualifiedName != nil {
		toSerialize["fullyQualifiedName"] = o.FullyQualifiedName
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if o.UsageSummary != nil {
		toSerialize["usageSummary"] = o.UsageSummary
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableMetrics struct {
	value *Metrics
	isSet bool
}

func (v NullableMetrics) Get() *Metrics {
	return v.value
}

func (v *NullableMetrics) Set(val *Metrics) {
	v.value = val
	v.isSet = true
}

func (v NullableMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetrics(val *Metrics) *NullableMetrics {
	return &NullableMetrics{value: val, isSet: true}
}

func (v NullableMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
