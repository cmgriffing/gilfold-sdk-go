/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BusinessResponse struct for BusinessResponse
type BusinessResponse struct {
	BusinessId string `json:"businessId"`
	Name string `json:"name"`
}

// NewBusinessResponse instantiates a new BusinessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessResponse(businessId string, name string) *BusinessResponse {
	this := BusinessResponse{}
	this.BusinessId = businessId
	this.Name = name
	return &this
}

// NewBusinessResponseWithDefaults instantiates a new BusinessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessResponseWithDefaults() *BusinessResponse {
	this := BusinessResponse{}
	return &this
}

// GetBusinessId returns the BusinessId field value
func (o *BusinessResponse) GetBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *BusinessResponse) GetBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *BusinessResponse) SetBusinessId(v string) {
	o.BusinessId = v
}

// GetName returns the Name field value
func (o *BusinessResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BusinessResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BusinessResponse) SetName(v string) {
	o.Name = v
}

func (o BusinessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["businessId"] = o.BusinessId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBusinessResponse struct {
	value *BusinessResponse
	isSet bool
}

func (v NullableBusinessResponse) Get() *BusinessResponse {
	return v.value
}

func (v *NullableBusinessResponse) Set(val *BusinessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessResponse(val *BusinessResponse) *NullableBusinessResponse {
	return &NullableBusinessResponse{value: val, isSet: true}
}

func (v NullableBusinessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


