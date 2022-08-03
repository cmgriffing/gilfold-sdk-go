/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PaymentProvider the model 'PaymentProvider'
type PaymentProvider string

// List of PaymentProvider
const (
	PAYPAL PaymentProvider = "paypal"
	STRIPE PaymentProvider = "stripe"
)

// All allowed values of PaymentProvider enum
var AllowedPaymentProviderEnumValues = []PaymentProvider{
	"paypal",
	"stripe",
}

func (v *PaymentProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentProvider(value)
	for _, existing := range AllowedPaymentProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentProvider", value)
}

// NewPaymentProviderFromValue returns a pointer to a valid PaymentProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentProviderFromValue(v string) (*PaymentProvider, error) {
	ev := PaymentProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentProvider: valid values are %v", v, AllowedPaymentProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentProvider) IsValid() bool {
	for _, existing := range AllowedPaymentProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentProvider value
func (v PaymentProvider) Ptr() *PaymentProvider {
	return &v
}

type NullablePaymentProvider struct {
	value *PaymentProvider
	isSet bool
}

func (v NullablePaymentProvider) Get() *PaymentProvider {
	return v.value
}

func (v *NullablePaymentProvider) Set(val *PaymentProvider) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentProvider) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentProvider(val *PaymentProvider) *NullablePaymentProvider {
	return &NullablePaymentProvider{value: val, isSet: true}
}

func (v NullablePaymentProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

