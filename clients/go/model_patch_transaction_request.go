/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PatchTransactionRequest struct for PatchTransactionRequest
type PatchTransactionRequest struct {
	AccountId *string `json:"accountId,omitempty"`
	Amount *float32 `json:"amount,omitempty"`
	Note *string `json:"note,omitempty"`
	PaymentProvider *PaymentProvider `json:"paymentProvider,omitempty"`
	PaymentId *string `json:"paymentId,omitempty"`
}

// NewPatchTransactionRequest instantiates a new PatchTransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchTransactionRequest() *PatchTransactionRequest {
	this := PatchTransactionRequest{}
	return &this
}

// NewPatchTransactionRequestWithDefaults instantiates a new PatchTransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchTransactionRequestWithDefaults() *PatchTransactionRequest {
	this := PatchTransactionRequest{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *PatchTransactionRequest) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTransactionRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *PatchTransactionRequest) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *PatchTransactionRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PatchTransactionRequest) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTransactionRequest) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PatchTransactionRequest) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *PatchTransactionRequest) SetAmount(v float32) {
	o.Amount = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *PatchTransactionRequest) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTransactionRequest) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *PatchTransactionRequest) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *PatchTransactionRequest) SetNote(v string) {
	o.Note = &v
}

// GetPaymentProvider returns the PaymentProvider field value if set, zero value otherwise.
func (o *PatchTransactionRequest) GetPaymentProvider() PaymentProvider {
	if o == nil || o.PaymentProvider == nil {
		var ret PaymentProvider
		return ret
	}
	return *o.PaymentProvider
}

// GetPaymentProviderOk returns a tuple with the PaymentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTransactionRequest) GetPaymentProviderOk() (*PaymentProvider, bool) {
	if o == nil || o.PaymentProvider == nil {
		return nil, false
	}
	return o.PaymentProvider, true
}

// HasPaymentProvider returns a boolean if a field has been set.
func (o *PatchTransactionRequest) HasPaymentProvider() bool {
	if o != nil && o.PaymentProvider != nil {
		return true
	}

	return false
}

// SetPaymentProvider gets a reference to the given PaymentProvider and assigns it to the PaymentProvider field.
func (o *PatchTransactionRequest) SetPaymentProvider(v PaymentProvider) {
	o.PaymentProvider = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *PatchTransactionRequest) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTransactionRequest) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *PatchTransactionRequest) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *PatchTransactionRequest) SetPaymentId(v string) {
	o.PaymentId = &v
}

func (o PatchTransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if o.PaymentProvider != nil {
		toSerialize["paymentProvider"] = o.PaymentProvider
	}
	if o.PaymentId != nil {
		toSerialize["paymentId"] = o.PaymentId
	}
	return json.Marshal(toSerialize)
}

type NullablePatchTransactionRequest struct {
	value *PatchTransactionRequest
	isSet bool
}

func (v NullablePatchTransactionRequest) Get() *PatchTransactionRequest {
	return v.value
}

func (v *NullablePatchTransactionRequest) Set(val *PatchTransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchTransactionRequest(val *PatchTransactionRequest) *NullablePatchTransactionRequest {
	return &NullablePatchTransactionRequest{value: val, isSet: true}
}

func (v NullablePatchTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


