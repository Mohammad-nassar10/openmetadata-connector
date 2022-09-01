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

// SearchShardTarget struct for SearchShardTarget
type SearchShardTarget struct {
	ClusterAlias            *string                `json:"clusterAlias,omitempty"`
	FullyQualifiedIndexName *string                `json:"fullyQualifiedIndexName,omitempty"`
	Index                   *string                `json:"index,omitempty"`
	NodeId                  *string                `json:"nodeId,omitempty"`
	NodeIdText              *Text                  `json:"nodeIdText,omitempty"`
	OriginalIndices         map[string]interface{} `json:"originalIndices,omitempty"`
	ShardId                 *ShardId               `json:"shardId,omitempty"`
}

// NewSearchShardTarget instantiates a new SearchShardTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchShardTarget() *SearchShardTarget {
	this := SearchShardTarget{}
	return &this
}

// NewSearchShardTargetWithDefaults instantiates a new SearchShardTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchShardTargetWithDefaults() *SearchShardTarget {
	this := SearchShardTarget{}
	return &this
}

// GetClusterAlias returns the ClusterAlias field value if set, zero value otherwise.
func (o *SearchShardTarget) GetClusterAlias() string {
	if o == nil || o.ClusterAlias == nil {
		var ret string
		return ret
	}
	return *o.ClusterAlias
}

// GetClusterAliasOk returns a tuple with the ClusterAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShardTarget) GetClusterAliasOk() (*string, bool) {
	if o == nil || o.ClusterAlias == nil {
		return nil, false
	}
	return o.ClusterAlias, true
}

// HasClusterAlias returns a boolean if a field has been set.
func (o *SearchShardTarget) HasClusterAlias() bool {
	if o != nil && o.ClusterAlias != nil {
		return true
	}

	return false
}

// SetClusterAlias gets a reference to the given string and assigns it to the ClusterAlias field.
func (o *SearchShardTarget) SetClusterAlias(v string) {
	o.ClusterAlias = &v
}

// GetFullyQualifiedIndexName returns the FullyQualifiedIndexName field value if set, zero value otherwise.
func (o *SearchShardTarget) GetFullyQualifiedIndexName() string {
	if o == nil || o.FullyQualifiedIndexName == nil {
		var ret string
		return ret
	}
	return *o.FullyQualifiedIndexName
}

// GetFullyQualifiedIndexNameOk returns a tuple with the FullyQualifiedIndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShardTarget) GetFullyQualifiedIndexNameOk() (*string, bool) {
	if o == nil || o.FullyQualifiedIndexName == nil {
		return nil, false
	}
	return o.FullyQualifiedIndexName, true
}

// HasFullyQualifiedIndexName returns a boolean if a field has been set.
func (o *SearchShardTarget) HasFullyQualifiedIndexName() bool {
	if o != nil && o.FullyQualifiedIndexName != nil {
		return true
	}

	return false
}

// SetFullyQualifiedIndexName gets a reference to the given string and assigns it to the FullyQualifiedIndexName field.
func (o *SearchShardTarget) SetFullyQualifiedIndexName(v string) {
	o.FullyQualifiedIndexName = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SearchShardTarget) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShardTarget) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SearchShardTarget) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *SearchShardTarget) SetIndex(v string) {
	o.Index = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *SearchShardTarget) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShardTarget) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *SearchShardTarget) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *SearchShardTarget) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeIdText returns the NodeIdText field value if set, zero value otherwise.
func (o *SearchShardTarget) GetNodeIdText() Text {
	if o == nil || o.NodeIdText == nil {
		var ret Text
		return ret
	}
	return *o.NodeIdText
}

// GetNodeIdTextOk returns a tuple with the NodeIdText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShardTarget) GetNodeIdTextOk() (*Text, bool) {
	if o == nil || o.NodeIdText == nil {
		return nil, false
	}
	return o.NodeIdText, true
}

// HasNodeIdText returns a boolean if a field has been set.
func (o *SearchShardTarget) HasNodeIdText() bool {
	if o != nil && o.NodeIdText != nil {
		return true
	}

	return false
}

// SetNodeIdText gets a reference to the given Text and assigns it to the NodeIdText field.
func (o *SearchShardTarget) SetNodeIdText(v Text) {
	o.NodeIdText = &v
}

// GetOriginalIndices returns the OriginalIndices field value if set, zero value otherwise.
func (o *SearchShardTarget) GetOriginalIndices() map[string]interface{} {
	if o == nil || o.OriginalIndices == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.OriginalIndices
}

// GetOriginalIndicesOk returns a tuple with the OriginalIndices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShardTarget) GetOriginalIndicesOk() (map[string]interface{}, bool) {
	if o == nil || o.OriginalIndices == nil {
		return nil, false
	}
	return o.OriginalIndices, true
}

// HasOriginalIndices returns a boolean if a field has been set.
func (o *SearchShardTarget) HasOriginalIndices() bool {
	if o != nil && o.OriginalIndices != nil {
		return true
	}

	return false
}

// SetOriginalIndices gets a reference to the given map[string]interface{} and assigns it to the OriginalIndices field.
func (o *SearchShardTarget) SetOriginalIndices(v map[string]interface{}) {
	o.OriginalIndices = v
}

// GetShardId returns the ShardId field value if set, zero value otherwise.
func (o *SearchShardTarget) GetShardId() ShardId {
	if o == nil || o.ShardId == nil {
		var ret ShardId
		return ret
	}
	return *o.ShardId
}

// GetShardIdOk returns a tuple with the ShardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShardTarget) GetShardIdOk() (*ShardId, bool) {
	if o == nil || o.ShardId == nil {
		return nil, false
	}
	return o.ShardId, true
}

// HasShardId returns a boolean if a field has been set.
func (o *SearchShardTarget) HasShardId() bool {
	if o != nil && o.ShardId != nil {
		return true
	}

	return false
}

// SetShardId gets a reference to the given ShardId and assigns it to the ShardId field.
func (o *SearchShardTarget) SetShardId(v ShardId) {
	o.ShardId = &v
}

func (o SearchShardTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterAlias != nil {
		toSerialize["clusterAlias"] = o.ClusterAlias
	}
	if o.FullyQualifiedIndexName != nil {
		toSerialize["fullyQualifiedIndexName"] = o.FullyQualifiedIndexName
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.NodeIdText != nil {
		toSerialize["nodeIdText"] = o.NodeIdText
	}
	if o.OriginalIndices != nil {
		toSerialize["originalIndices"] = o.OriginalIndices
	}
	if o.ShardId != nil {
		toSerialize["shardId"] = o.ShardId
	}
	return json.Marshal(toSerialize)
}

type NullableSearchShardTarget struct {
	value *SearchShardTarget
	isSet bool
}

func (v NullableSearchShardTarget) Get() *SearchShardTarget {
	return v.value
}

func (v *NullableSearchShardTarget) Set(val *SearchShardTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchShardTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchShardTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchShardTarget(val *SearchShardTarget) *NullableSearchShardTarget {
	return &NullableSearchShardTarget{value: val, isSet: true}
}

func (v NullableSearchShardTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchShardTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
