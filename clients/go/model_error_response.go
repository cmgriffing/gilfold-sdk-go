/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	TransactionId string `json:"transactionId"`
	AccountId string `json:"accountId"`
	BusinessId string `json:"businessId"`
	Amount float32 `json:"amount"`
	Note *string `json:"note,omitempty"`
	PaymentProvider *string `json:"paymentProvider,omitempty"`
	PaymentId *string `json:"paymentId,omitempty"`
	CreatedAt float32 `json:"createdAt"`
	ModifiedAt float32 `json:"modifiedAt"`
}

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse(transactionId string, accountId string, businessId string, amount float32, createdAt float32, modifiedAt float32) *ErrorResponse {
	this := ErrorResponse{}
	this.TransactionId = transactionId
	this.AccountId = accountId
	this.BusinessId = businessId
	this.Amount = amount
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *ErrorResponse) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *ErrorResponse) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetAccountId returns the AccountId field value
func (o *ErrorResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ErrorResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetBusinessId returns the BusinessId field value
func (o *ErrorResponse) GetBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *ErrorResponse) SetBusinessId(v string) {
	o.BusinessId = v
}

// GetAmount returns the Amount field value
func (o *ErrorResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ErrorResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *ErrorResponse) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ErrorResponse) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *ErrorResponse) SetNote(v string) {
	o.Note = &v
}

// GetPaymentProvider returns the PaymentProvider field value if set, zero value otherwise.
func (o *ErrorResponse) GetPaymentProvider() string {
	if o == nil || o.PaymentProvider == nil {
		var ret string
		return ret
	}
	return *o.PaymentProvider
}

// GetPaymentProviderOk returns a tuple with the PaymentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetPaymentProviderOk() (*string, bool) {
	if o == nil || o.PaymentProvider == nil {
		return nil, false
	}
	return o.PaymentProvider, true
}

// HasPaymentProvider returns a boolean if a field has been set.
func (o *ErrorResponse) HasPaymentProvider() bool {
	if o != nil && o.PaymentProvider != nil {
		return true
	}

	return false
}

// SetPaymentProvider gets a reference to the given string and assigns it to the PaymentProvider field.
func (o *ErrorResponse) SetPaymentProvider(v string) {
	o.PaymentProvider = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *ErrorResponse) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *ErrorResponse) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *ErrorResponse) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ErrorResponse) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ErrorResponse) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *ErrorResponse) GetModifiedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *ErrorResponse) GetModifiedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *ErrorResponse) SetModifiedAt(v float32) {
	o.ModifiedAt = v
}

func (o ErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transactionId"] = o.TransactionId
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["businessId"] = o.BusinessId
	}
	if true {
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
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


