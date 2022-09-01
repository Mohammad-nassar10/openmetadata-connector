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

// TopicSampleData struct for TopicSampleData
type TopicSampleData struct {
	Messages []string `json:"messages,omitempty"`
}

// NewTopicSampleData instantiates a new TopicSampleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicSampleData() *TopicSampleData {
	this := TopicSampleData{}
	return &this
}

// NewTopicSampleDataWithDefaults instantiates a new TopicSampleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicSampleDataWithDefaults() *TopicSampleData {
	this := TopicSampleData{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *TopicSampleData) GetMessages() []string {
	if o == nil || o.Messages == nil {
		var ret []string
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopicSampleData) GetMessagesOk() ([]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *TopicSampleData) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *TopicSampleData) SetMessages(v []string) {
	o.Messages = v
}

func (o TopicSampleData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableTopicSampleData struct {
	value *TopicSampleData
	isSet bool
}

func (v NullableTopicSampleData) Get() *TopicSampleData {
	return v.value
}

func (v *NullableTopicSampleData) Set(val *TopicSampleData) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicSampleData) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicSampleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicSampleData(val *TopicSampleData) *NullableTopicSampleData {
	return &NullableTopicSampleData{value: val, isSet: true}
}

func (v NullableTopicSampleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicSampleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
