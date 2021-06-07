/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RetentionPolicyManifest struct for RetentionPolicyManifest
type RetentionPolicyManifest struct {
	Name               string                 `json:"name"`
	ReplicaN           int32                  `json:"replicaN"`
	Duration           int64                  `json:"duration"`
	ShardGroupDuration int64                  `json:"shardGroupDuration"`
	ShardGroups        []ShardGroupManifest   `json:"shardGroups"`
	Subscriptions      []SubscriptionManifest `json:"subscriptions"`
}

// NewRetentionPolicyManifest instantiates a new RetentionPolicyManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetentionPolicyManifest(name string, replicaN int32, duration int64, shardGroupDuration int64, shardGroups []ShardGroupManifest, subscriptions []SubscriptionManifest) *RetentionPolicyManifest {
	this := RetentionPolicyManifest{}
	this.Name = name
	this.ReplicaN = replicaN
	this.Duration = duration
	this.ShardGroupDuration = shardGroupDuration
	this.ShardGroups = shardGroups
	this.Subscriptions = subscriptions
	return &this
}

// NewRetentionPolicyManifestWithDefaults instantiates a new RetentionPolicyManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionPolicyManifestWithDefaults() *RetentionPolicyManifest {
	this := RetentionPolicyManifest{}
	return &this
}

// GetName returns the Name field value
func (o *RetentionPolicyManifest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RetentionPolicyManifest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RetentionPolicyManifest) SetName(v string) {
	o.Name = v
}

// GetReplicaN returns the ReplicaN field value
func (o *RetentionPolicyManifest) GetReplicaN() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReplicaN
}

// GetReplicaNOk returns a tuple with the ReplicaN field value
// and a boolean to check if the value has been set.
func (o *RetentionPolicyManifest) GetReplicaNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplicaN, true
}

// SetReplicaN sets field value
func (o *RetentionPolicyManifest) SetReplicaN(v int32) {
	o.ReplicaN = v
}

// GetDuration returns the Duration field value
func (o *RetentionPolicyManifest) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *RetentionPolicyManifest) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *RetentionPolicyManifest) SetDuration(v int64) {
	o.Duration = v
}

// GetShardGroupDuration returns the ShardGroupDuration field value
func (o *RetentionPolicyManifest) GetShardGroupDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ShardGroupDuration
}

// GetShardGroupDurationOk returns a tuple with the ShardGroupDuration field value
// and a boolean to check if the value has been set.
func (o *RetentionPolicyManifest) GetShardGroupDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShardGroupDuration, true
}

// SetShardGroupDuration sets field value
func (o *RetentionPolicyManifest) SetShardGroupDuration(v int64) {
	o.ShardGroupDuration = v
}

// GetShardGroups returns the ShardGroups field value
func (o *RetentionPolicyManifest) GetShardGroups() []ShardGroupManifest {
	if o == nil {
		var ret []ShardGroupManifest
		return ret
	}

	return o.ShardGroups
}

// GetShardGroupsOk returns a tuple with the ShardGroups field value
// and a boolean to check if the value has been set.
func (o *RetentionPolicyManifest) GetShardGroupsOk() (*[]ShardGroupManifest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShardGroups, true
}

// SetShardGroups sets field value
func (o *RetentionPolicyManifest) SetShardGroups(v []ShardGroupManifest) {
	o.ShardGroups = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *RetentionPolicyManifest) GetSubscriptions() []SubscriptionManifest {
	if o == nil {
		var ret []SubscriptionManifest
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *RetentionPolicyManifest) GetSubscriptionsOk() (*[]SubscriptionManifest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *RetentionPolicyManifest) SetSubscriptions(v []SubscriptionManifest) {
	o.Subscriptions = v
}

func (o RetentionPolicyManifest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["replicaN"] = o.ReplicaN
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["shardGroupDuration"] = o.ShardGroupDuration
	}
	if true {
		toSerialize["shardGroups"] = o.ShardGroups
	}
	if true {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	return json.Marshal(toSerialize)
}

type NullableRetentionPolicyManifest struct {
	value *RetentionPolicyManifest
	isSet bool
}

func (v NullableRetentionPolicyManifest) Get() *RetentionPolicyManifest {
	return v.value
}

func (v *NullableRetentionPolicyManifest) Set(val *RetentionPolicyManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionPolicyManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionPolicyManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionPolicyManifest(val *RetentionPolicyManifest) *NullableRetentionPolicyManifest {
	return &NullableRetentionPolicyManifest{value: val, isSet: true}
}

func (v NullableRetentionPolicyManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionPolicyManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}