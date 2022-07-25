/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PatchBusinessSettingsRequest struct for PatchBusinessSettingsRequest
type PatchBusinessSettingsRequest struct {
	CurrencyLabel *string `json:"currencyLabel,omitempty"`
	CurrencyLabelSuffixed *bool `json:"currencyLabelSuffixed,omitempty"`
	AccountsOverdraftable *bool `json:"accountsOverdraftable,omitempty"`
	BillingType *BillingType `json:"billingType,omitempty"`
	PaymentProviders []string `json:"paymentProviders,omitempty"`
	StripeSandboxPublishableKey *string `json:"stripeSandboxPublishableKey,omitempty"`
	StripeSandboxSecretKey *string `json:"stripeSandboxSecretKey,omitempty"`
	StripePublishableKey *string `json:"stripePublishableKey,omitempty"`
	StripeSecretKey *string `json:"stripeSecretKey,omitempty"`
	PaypalClientId *string `json:"paypalClientId,omitempty"`
	PaypalClientSecret *string `json:"paypalClientSecret,omitempty"`
}

// NewPatchBusinessSettingsRequest instantiates a new PatchBusinessSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchBusinessSettingsRequest() *PatchBusinessSettingsRequest {
	this := PatchBusinessSettingsRequest{}
	return &this
}

// NewPatchBusinessSettingsRequestWithDefaults instantiates a new PatchBusinessSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchBusinessSettingsRequestWithDefaults() *PatchBusinessSettingsRequest {
	this := PatchBusinessSettingsRequest{}
	return &this
}

// GetCurrencyLabel returns the CurrencyLabel field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetCurrencyLabel() string {
	if o == nil || o.CurrencyLabel == nil {
		var ret string
		return ret
	}
	return *o.CurrencyLabel
}

// GetCurrencyLabelOk returns a tuple with the CurrencyLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetCurrencyLabelOk() (*string, bool) {
	if o == nil || o.CurrencyLabel == nil {
		return nil, false
	}
	return o.CurrencyLabel, true
}

// HasCurrencyLabel returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasCurrencyLabel() bool {
	if o != nil && o.CurrencyLabel != nil {
		return true
	}

	return false
}

// SetCurrencyLabel gets a reference to the given string and assigns it to the CurrencyLabel field.
func (o *PatchBusinessSettingsRequest) SetCurrencyLabel(v string) {
	o.CurrencyLabel = &v
}

// GetCurrencyLabelSuffixed returns the CurrencyLabelSuffixed field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetCurrencyLabelSuffixed() bool {
	if o == nil || o.CurrencyLabelSuffixed == nil {
		var ret bool
		return ret
	}
	return *o.CurrencyLabelSuffixed
}

// GetCurrencyLabelSuffixedOk returns a tuple with the CurrencyLabelSuffixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetCurrencyLabelSuffixedOk() (*bool, bool) {
	if o == nil || o.CurrencyLabelSuffixed == nil {
		return nil, false
	}
	return o.CurrencyLabelSuffixed, true
}

// HasCurrencyLabelSuffixed returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasCurrencyLabelSuffixed() bool {
	if o != nil && o.CurrencyLabelSuffixed != nil {
		return true
	}

	return false
}

// SetCurrencyLabelSuffixed gets a reference to the given bool and assigns it to the CurrencyLabelSuffixed field.
func (o *PatchBusinessSettingsRequest) SetCurrencyLabelSuffixed(v bool) {
	o.CurrencyLabelSuffixed = &v
}

// GetAccountsOverdraftable returns the AccountsOverdraftable field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetAccountsOverdraftable() bool {
	if o == nil || o.AccountsOverdraftable == nil {
		var ret bool
		return ret
	}
	return *o.AccountsOverdraftable
}

// GetAccountsOverdraftableOk returns a tuple with the AccountsOverdraftable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetAccountsOverdraftableOk() (*bool, bool) {
	if o == nil || o.AccountsOverdraftable == nil {
		return nil, false
	}
	return o.AccountsOverdraftable, true
}

// HasAccountsOverdraftable returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasAccountsOverdraftable() bool {
	if o != nil && o.AccountsOverdraftable != nil {
		return true
	}

	return false
}

// SetAccountsOverdraftable gets a reference to the given bool and assigns it to the AccountsOverdraftable field.
func (o *PatchBusinessSettingsRequest) SetAccountsOverdraftable(v bool) {
	o.AccountsOverdraftable = &v
}

// GetBillingType returns the BillingType field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetBillingType() BillingType {
	if o == nil || o.BillingType == nil {
		var ret BillingType
		return ret
	}
	return *o.BillingType
}

// GetBillingTypeOk returns a tuple with the BillingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetBillingTypeOk() (*BillingType, bool) {
	if o == nil || o.BillingType == nil {
		return nil, false
	}
	return o.BillingType, true
}

// HasBillingType returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasBillingType() bool {
	if o != nil && o.BillingType != nil {
		return true
	}

	return false
}

// SetBillingType gets a reference to the given BillingType and assigns it to the BillingType field.
func (o *PatchBusinessSettingsRequest) SetBillingType(v BillingType) {
	o.BillingType = &v
}

// GetPaymentProviders returns the PaymentProviders field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetPaymentProviders() []string {
	if o == nil || o.PaymentProviders == nil {
		var ret []string
		return ret
	}
	return o.PaymentProviders
}

// GetPaymentProvidersOk returns a tuple with the PaymentProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetPaymentProvidersOk() ([]string, bool) {
	if o == nil || o.PaymentProviders == nil {
		return nil, false
	}
	return o.PaymentProviders, true
}

// HasPaymentProviders returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasPaymentProviders() bool {
	if o != nil && o.PaymentProviders != nil {
		return true
	}

	return false
}

// SetPaymentProviders gets a reference to the given []string and assigns it to the PaymentProviders field.
func (o *PatchBusinessSettingsRequest) SetPaymentProviders(v []string) {
	o.PaymentProviders = v
}

// GetStripeSandboxPublishableKey returns the StripeSandboxPublishableKey field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetStripeSandboxPublishableKey() string {
	if o == nil || o.StripeSandboxPublishableKey == nil {
		var ret string
		return ret
	}
	return *o.StripeSandboxPublishableKey
}

// GetStripeSandboxPublishableKeyOk returns a tuple with the StripeSandboxPublishableKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetStripeSandboxPublishableKeyOk() (*string, bool) {
	if o == nil || o.StripeSandboxPublishableKey == nil {
		return nil, false
	}
	return o.StripeSandboxPublishableKey, true
}

// HasStripeSandboxPublishableKey returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasStripeSandboxPublishableKey() bool {
	if o != nil && o.StripeSandboxPublishableKey != nil {
		return true
	}

	return false
}

// SetStripeSandboxPublishableKey gets a reference to the given string and assigns it to the StripeSandboxPublishableKey field.
func (o *PatchBusinessSettingsRequest) SetStripeSandboxPublishableKey(v string) {
	o.StripeSandboxPublishableKey = &v
}

// GetStripeSandboxSecretKey returns the StripeSandboxSecretKey field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetStripeSandboxSecretKey() string {
	if o == nil || o.StripeSandboxSecretKey == nil {
		var ret string
		return ret
	}
	return *o.StripeSandboxSecretKey
}

// GetStripeSandboxSecretKeyOk returns a tuple with the StripeSandboxSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetStripeSandboxSecretKeyOk() (*string, bool) {
	if o == nil || o.StripeSandboxSecretKey == nil {
		return nil, false
	}
	return o.StripeSandboxSecretKey, true
}

// HasStripeSandboxSecretKey returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasStripeSandboxSecretKey() bool {
	if o != nil && o.StripeSandboxSecretKey != nil {
		return true
	}

	return false
}

// SetStripeSandboxSecretKey gets a reference to the given string and assigns it to the StripeSandboxSecretKey field.
func (o *PatchBusinessSettingsRequest) SetStripeSandboxSecretKey(v string) {
	o.StripeSandboxSecretKey = &v
}

// GetStripePublishableKey returns the StripePublishableKey field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetStripePublishableKey() string {
	if o == nil || o.StripePublishableKey == nil {
		var ret string
		return ret
	}
	return *o.StripePublishableKey
}

// GetStripePublishableKeyOk returns a tuple with the StripePublishableKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetStripePublishableKeyOk() (*string, bool) {
	if o == nil || o.StripePublishableKey == nil {
		return nil, false
	}
	return o.StripePublishableKey, true
}

// HasStripePublishableKey returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasStripePublishableKey() bool {
	if o != nil && o.StripePublishableKey != nil {
		return true
	}

	return false
}

// SetStripePublishableKey gets a reference to the given string and assigns it to the StripePublishableKey field.
func (o *PatchBusinessSettingsRequest) SetStripePublishableKey(v string) {
	o.StripePublishableKey = &v
}

// GetStripeSecretKey returns the StripeSecretKey field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetStripeSecretKey() string {
	if o == nil || o.StripeSecretKey == nil {
		var ret string
		return ret
	}
	return *o.StripeSecretKey
}

// GetStripeSecretKeyOk returns a tuple with the StripeSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetStripeSecretKeyOk() (*string, bool) {
	if o == nil || o.StripeSecretKey == nil {
		return nil, false
	}
	return o.StripeSecretKey, true
}

// HasStripeSecretKey returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasStripeSecretKey() bool {
	if o != nil && o.StripeSecretKey != nil {
		return true
	}

	return false
}

// SetStripeSecretKey gets a reference to the given string and assigns it to the StripeSecretKey field.
func (o *PatchBusinessSettingsRequest) SetStripeSecretKey(v string) {
	o.StripeSecretKey = &v
}

// GetPaypalClientId returns the PaypalClientId field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetPaypalClientId() string {
	if o == nil || o.PaypalClientId == nil {
		var ret string
		return ret
	}
	return *o.PaypalClientId
}

// GetPaypalClientIdOk returns a tuple with the PaypalClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetPaypalClientIdOk() (*string, bool) {
	if o == nil || o.PaypalClientId == nil {
		return nil, false
	}
	return o.PaypalClientId, true
}

// HasPaypalClientId returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasPaypalClientId() bool {
	if o != nil && o.PaypalClientId != nil {
		return true
	}

	return false
}

// SetPaypalClientId gets a reference to the given string and assigns it to the PaypalClientId field.
func (o *PatchBusinessSettingsRequest) SetPaypalClientId(v string) {
	o.PaypalClientId = &v
}

// GetPaypalClientSecret returns the PaypalClientSecret field value if set, zero value otherwise.
func (o *PatchBusinessSettingsRequest) GetPaypalClientSecret() string {
	if o == nil || o.PaypalClientSecret == nil {
		var ret string
		return ret
	}
	return *o.PaypalClientSecret
}

// GetPaypalClientSecretOk returns a tuple with the PaypalClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessSettingsRequest) GetPaypalClientSecretOk() (*string, bool) {
	if o == nil || o.PaypalClientSecret == nil {
		return nil, false
	}
	return o.PaypalClientSecret, true
}

// HasPaypalClientSecret returns a boolean if a field has been set.
func (o *PatchBusinessSettingsRequest) HasPaypalClientSecret() bool {
	if o != nil && o.PaypalClientSecret != nil {
		return true
	}

	return false
}

// SetPaypalClientSecret gets a reference to the given string and assigns it to the PaypalClientSecret field.
func (o *PatchBusinessSettingsRequest) SetPaypalClientSecret(v string) {
	o.PaypalClientSecret = &v
}

func (o PatchBusinessSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencyLabel != nil {
		toSerialize["currencyLabel"] = o.CurrencyLabel
	}
	if o.CurrencyLabelSuffixed != nil {
		toSerialize["currencyLabelSuffixed"] = o.CurrencyLabelSuffixed
	}
	if o.AccountsOverdraftable != nil {
		toSerialize["accountsOverdraftable"] = o.AccountsOverdraftable
	}
	if o.BillingType != nil {
		toSerialize["billingType"] = o.BillingType
	}
	if o.PaymentProviders != nil {
		toSerialize["paymentProviders"] = o.PaymentProviders
	}
	if o.StripeSandboxPublishableKey != nil {
		toSerialize["stripeSandboxPublishableKey"] = o.StripeSandboxPublishableKey
	}
	if o.StripeSandboxSecretKey != nil {
		toSerialize["stripeSandboxSecretKey"] = o.StripeSandboxSecretKey
	}
	if o.StripePublishableKey != nil {
		toSerialize["stripePublishableKey"] = o.StripePublishableKey
	}
	if o.StripeSecretKey != nil {
		toSerialize["stripeSecretKey"] = o.StripeSecretKey
	}
	if o.PaypalClientId != nil {
		toSerialize["paypalClientId"] = o.PaypalClientId
	}
	if o.PaypalClientSecret != nil {
		toSerialize["paypalClientSecret"] = o.PaypalClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullablePatchBusinessSettingsRequest struct {
	value *PatchBusinessSettingsRequest
	isSet bool
}

func (v NullablePatchBusinessSettingsRequest) Get() *PatchBusinessSettingsRequest {
	return v.value
}

func (v *NullablePatchBusinessSettingsRequest) Set(val *PatchBusinessSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBusinessSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBusinessSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBusinessSettingsRequest(val *PatchBusinessSettingsRequest) *NullablePatchBusinessSettingsRequest {
	return &NullablePatchBusinessSettingsRequest{value: val, isSet: true}
}

func (v NullablePatchBusinessSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBusinessSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


