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
	"time"
)

// RunManually struct for RunManually
type RunManually struct {
	// Time used for run's \"now\" option, RFC3339.  Default is the server's now time.
	ScheduledFor NullableTime `json:"scheduledFor,omitempty"`
}

// NewRunManually instantiates a new RunManually object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunManually() *RunManually {
	this := RunManually{}
	return &this
}

// NewRunManuallyWithDefaults instantiates a new RunManually object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunManuallyWithDefaults() *RunManually {
	this := RunManually{}
	return &this
}

// GetScheduledFor returns the ScheduledFor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunManually) GetScheduledFor() time.Time {
	if o == nil || o.ScheduledFor.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledFor.Get()
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunManually) GetScheduledForOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledFor.Get(), o.ScheduledFor.IsSet()
}

// HasScheduledFor returns a boolean if a field has been set.
func (o *RunManually) HasScheduledFor() bool {
	if o != nil && o.ScheduledFor.IsSet() {
		return true
	}

	return false
}

// SetScheduledFor gets a reference to the given NullableTime and assigns it to the ScheduledFor field.
func (o *RunManually) SetScheduledFor(v time.Time) {
	o.ScheduledFor.Set(&v)
}

// SetScheduledForNil sets the value for ScheduledFor to be an explicit nil
func (o *RunManually) SetScheduledForNil() {
	o.ScheduledFor.Set(nil)
}

// UnsetScheduledFor ensures that no value is present for ScheduledFor, not even an explicit nil
func (o *RunManually) UnsetScheduledFor() {
	o.ScheduledFor.Unset()
}

func (o RunManually) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScheduledFor.IsSet() {
		toSerialize["scheduledFor"] = o.ScheduledFor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRunManually struct {
	value *RunManually
	isSet bool
}

func (v NullableRunManually) Get() *RunManually {
	return v.value
}

func (v *NullableRunManually) Set(val *RunManually) {
	v.value = val
	v.isSet = true
}

func (v NullableRunManually) IsSet() bool {
	return v.isSet
}

func (v *NullableRunManually) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunManually(val *RunManually) *NullableRunManually {
	return &NullableRunManually{value: val, isSet: true}
}

func (v NullableRunManually) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunManually) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}