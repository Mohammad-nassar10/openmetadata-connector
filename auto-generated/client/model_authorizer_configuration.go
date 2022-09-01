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

// AuthorizerConfiguration struct for AuthorizerConfiguration
type AuthorizerConfiguration struct {
	AdminPrincipals              []string `json:"adminPrincipals"`
	BotPrincipals                []string `json:"botPrincipals"`
	ClassName                    string   `json:"className"`
	ContainerRequestFilter       string   `json:"containerRequestFilter"`
	EnableSecureSocketConnection bool     `json:"enableSecureSocketConnection"`
	EnforcePrincipalDomain       bool     `json:"enforcePrincipalDomain"`
	PrincipalDomain              string   `json:"principalDomain"`
}

// NewAuthorizerConfiguration instantiates a new AuthorizerConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizerConfiguration(adminPrincipals []string, botPrincipals []string, className string, containerRequestFilter string, enableSecureSocketConnection bool, enforcePrincipalDomain bool, principalDomain string) *AuthorizerConfiguration {
	this := AuthorizerConfiguration{}
	this.AdminPrincipals = adminPrincipals
	this.BotPrincipals = botPrincipals
	this.ClassName = className
	this.ContainerRequestFilter = containerRequestFilter
	this.EnableSecureSocketConnection = enableSecureSocketConnection
	this.EnforcePrincipalDomain = enforcePrincipalDomain
	this.PrincipalDomain = principalDomain
	return &this
}

// NewAuthorizerConfigurationWithDefaults instantiates a new AuthorizerConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizerConfigurationWithDefaults() *AuthorizerConfiguration {
	this := AuthorizerConfiguration{}
	return &this
}

// GetAdminPrincipals returns the AdminPrincipals field value
func (o *AuthorizerConfiguration) GetAdminPrincipals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AdminPrincipals
}

// GetAdminPrincipalsOk returns a tuple with the AdminPrincipals field value
// and a boolean to check if the value has been set.
func (o *AuthorizerConfiguration) GetAdminPrincipalsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdminPrincipals, true
}

// SetAdminPrincipals sets field value
func (o *AuthorizerConfiguration) SetAdminPrincipals(v []string) {
	o.AdminPrincipals = v
}

// GetBotPrincipals returns the BotPrincipals field value
func (o *AuthorizerConfiguration) GetBotPrincipals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BotPrincipals
}

// GetBotPrincipalsOk returns a tuple with the BotPrincipals field value
// and a boolean to check if the value has been set.
func (o *AuthorizerConfiguration) GetBotPrincipalsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BotPrincipals, true
}

// SetBotPrincipals sets field value
func (o *AuthorizerConfiguration) SetBotPrincipals(v []string) {
	o.BotPrincipals = v
}

// GetClassName returns the ClassName field value
func (o *AuthorizerConfiguration) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value
// and a boolean to check if the value has been set.
func (o *AuthorizerConfiguration) GetClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassName, true
}

// SetClassName sets field value
func (o *AuthorizerConfiguration) SetClassName(v string) {
	o.ClassName = v
}

// GetContainerRequestFilter returns the ContainerRequestFilter field value
func (o *AuthorizerConfiguration) GetContainerRequestFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerRequestFilter
}

// GetContainerRequestFilterOk returns a tuple with the ContainerRequestFilter field value
// and a boolean to check if the value has been set.
func (o *AuthorizerConfiguration) GetContainerRequestFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerRequestFilter, true
}

// SetContainerRequestFilter sets field value
func (o *AuthorizerConfiguration) SetContainerRequestFilter(v string) {
	o.ContainerRequestFilter = v
}

// GetEnableSecureSocketConnection returns the EnableSecureSocketConnection field value
func (o *AuthorizerConfiguration) GetEnableSecureSocketConnection() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableSecureSocketConnection
}

// GetEnableSecureSocketConnectionOk returns a tuple with the EnableSecureSocketConnection field value
// and a boolean to check if the value has been set.
func (o *AuthorizerConfiguration) GetEnableSecureSocketConnectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableSecureSocketConnection, true
}

// SetEnableSecureSocketConnection sets field value
func (o *AuthorizerConfiguration) SetEnableSecureSocketConnection(v bool) {
	o.EnableSecureSocketConnection = v
}

// GetEnforcePrincipalDomain returns the EnforcePrincipalDomain field value
func (o *AuthorizerConfiguration) GetEnforcePrincipalDomain() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnforcePrincipalDomain
}

// GetEnforcePrincipalDomainOk returns a tuple with the EnforcePrincipalDomain field value
// and a boolean to check if the value has been set.
func (o *AuthorizerConfiguration) GetEnforcePrincipalDomainOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnforcePrincipalDomain, true
}

// SetEnforcePrincipalDomain sets field value
func (o *AuthorizerConfiguration) SetEnforcePrincipalDomain(v bool) {
	o.EnforcePrincipalDomain = v
}

// GetPrincipalDomain returns the PrincipalDomain field value
func (o *AuthorizerConfiguration) GetPrincipalDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalDomain
}

// GetPrincipalDomainOk returns a tuple with the PrincipalDomain field value
// and a boolean to check if the value has been set.
func (o *AuthorizerConfiguration) GetPrincipalDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalDomain, true
}

// SetPrincipalDomain sets field value
func (o *AuthorizerConfiguration) SetPrincipalDomain(v string) {
	o.PrincipalDomain = v
}

func (o AuthorizerConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["adminPrincipals"] = o.AdminPrincipals
	}
	if true {
		toSerialize["botPrincipals"] = o.BotPrincipals
	}
	if true {
		toSerialize["className"] = o.ClassName
	}
	if true {
		toSerialize["containerRequestFilter"] = o.ContainerRequestFilter
	}
	if true {
		toSerialize["enableSecureSocketConnection"] = o.EnableSecureSocketConnection
	}
	if true {
		toSerialize["enforcePrincipalDomain"] = o.EnforcePrincipalDomain
	}
	if true {
		toSerialize["principalDomain"] = o.PrincipalDomain
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizerConfiguration struct {
	value *AuthorizerConfiguration
	isSet bool
}

func (v NullableAuthorizerConfiguration) Get() *AuthorizerConfiguration {
	return v.value
}

func (v *NullableAuthorizerConfiguration) Set(val *AuthorizerConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizerConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizerConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizerConfiguration(val *AuthorizerConfiguration) *NullableAuthorizerConfiguration {
	return &NullableAuthorizerConfiguration{value: val, isSet: true}
}

func (v NullableAuthorizerConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizerConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
