# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.3
    Generated by: https://openapi-generator.tech
"""


try:
    from inspect import getfullargspec
except ImportError:
    from inspect import getargspec as getfullargspec
import pprint
import re  # noqa: F401
import six

from openapi_client.configuration import Configuration


class PatchBusinessSettingsRequest(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'currency_label': 'str',
        'currency_label_suffixed': 'bool',
        'accounts_overdraftable': 'bool',
        'billing_type': 'BillingType',
        'payment_providers': 'list[str]',
        'stripe_sandbox_publishable_key': 'str',
        'stripe_sandbox_secret_key': 'str',
        'stripe_publishable_key': 'str',
        'stripe_secret_key': 'str',
        'paypal_client_id': 'str',
        'paypal_client_secret': 'str'
    }

    attribute_map = {
        'currency_label': 'currencyLabel',
        'currency_label_suffixed': 'currencyLabelSuffixed',
        'accounts_overdraftable': 'accountsOverdraftable',
        'billing_type': 'billingType',
        'payment_providers': 'paymentProviders',
        'stripe_sandbox_publishable_key': 'stripeSandboxPublishableKey',
        'stripe_sandbox_secret_key': 'stripeSandboxSecretKey',
        'stripe_publishable_key': 'stripePublishableKey',
        'stripe_secret_key': 'stripeSecretKey',
        'paypal_client_id': 'paypalClientId',
        'paypal_client_secret': 'paypalClientSecret'
    }

    def __init__(self, currency_label=None, currency_label_suffixed=None, accounts_overdraftable=None, billing_type=None, payment_providers=None, stripe_sandbox_publishable_key=None, stripe_sandbox_secret_key=None, stripe_publishable_key=None, stripe_secret_key=None, paypal_client_id=None, paypal_client_secret=None, local_vars_configuration=None):  # noqa: E501
        """PatchBusinessSettingsRequest - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._currency_label = None
        self._currency_label_suffixed = None
        self._accounts_overdraftable = None
        self._billing_type = None
        self._payment_providers = None
        self._stripe_sandbox_publishable_key = None
        self._stripe_sandbox_secret_key = None
        self._stripe_publishable_key = None
        self._stripe_secret_key = None
        self._paypal_client_id = None
        self._paypal_client_secret = None
        self.discriminator = None

        if currency_label is not None:
            self.currency_label = currency_label
        if currency_label_suffixed is not None:
            self.currency_label_suffixed = currency_label_suffixed
        if accounts_overdraftable is not None:
            self.accounts_overdraftable = accounts_overdraftable
        if billing_type is not None:
            self.billing_type = billing_type
        if payment_providers is not None:
            self.payment_providers = payment_providers
        if stripe_sandbox_publishable_key is not None:
            self.stripe_sandbox_publishable_key = stripe_sandbox_publishable_key
        if stripe_sandbox_secret_key is not None:
            self.stripe_sandbox_secret_key = stripe_sandbox_secret_key
        if stripe_publishable_key is not None:
            self.stripe_publishable_key = stripe_publishable_key
        if stripe_secret_key is not None:
            self.stripe_secret_key = stripe_secret_key
        if paypal_client_id is not None:
            self.paypal_client_id = paypal_client_id
        if paypal_client_secret is not None:
            self.paypal_client_secret = paypal_client_secret

    @property
    def currency_label(self):
        """Gets the currency_label of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The currency_label of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: str
        """
        return self._currency_label

    @currency_label.setter
    def currency_label(self, currency_label):
        """Sets the currency_label of this PatchBusinessSettingsRequest.


        :param currency_label: The currency_label of this PatchBusinessSettingsRequest.  # noqa: E501
        :type currency_label: str
        """

        self._currency_label = currency_label

    @property
    def currency_label_suffixed(self):
        """Gets the currency_label_suffixed of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The currency_label_suffixed of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: bool
        """
        return self._currency_label_suffixed

    @currency_label_suffixed.setter
    def currency_label_suffixed(self, currency_label_suffixed):
        """Sets the currency_label_suffixed of this PatchBusinessSettingsRequest.


        :param currency_label_suffixed: The currency_label_suffixed of this PatchBusinessSettingsRequest.  # noqa: E501
        :type currency_label_suffixed: bool
        """

        self._currency_label_suffixed = currency_label_suffixed

    @property
    def accounts_overdraftable(self):
        """Gets the accounts_overdraftable of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The accounts_overdraftable of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: bool
        """
        return self._accounts_overdraftable

    @accounts_overdraftable.setter
    def accounts_overdraftable(self, accounts_overdraftable):
        """Sets the accounts_overdraftable of this PatchBusinessSettingsRequest.


        :param accounts_overdraftable: The accounts_overdraftable of this PatchBusinessSettingsRequest.  # noqa: E501
        :type accounts_overdraftable: bool
        """

        self._accounts_overdraftable = accounts_overdraftable

    @property
    def billing_type(self):
        """Gets the billing_type of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The billing_type of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: BillingType
        """
        return self._billing_type

    @billing_type.setter
    def billing_type(self, billing_type):
        """Sets the billing_type of this PatchBusinessSettingsRequest.


        :param billing_type: The billing_type of this PatchBusinessSettingsRequest.  # noqa: E501
        :type billing_type: BillingType
        """

        self._billing_type = billing_type

    @property
    def payment_providers(self):
        """Gets the payment_providers of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The payment_providers of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: list[str]
        """
        return self._payment_providers

    @payment_providers.setter
    def payment_providers(self, payment_providers):
        """Sets the payment_providers of this PatchBusinessSettingsRequest.


        :param payment_providers: The payment_providers of this PatchBusinessSettingsRequest.  # noqa: E501
        :type payment_providers: list[str]
        """
        allowed_values = ["paypal", "stripe"]  # noqa: E501
        if (self.local_vars_configuration.client_side_validation and
                not set(payment_providers).issubset(set(allowed_values))):  # noqa: E501
            raise ValueError(
                "Invalid values for `payment_providers` [{0}], must be a subset of [{1}]"  # noqa: E501
                .format(", ".join(map(str, set(payment_providers) - set(allowed_values))),  # noqa: E501
                        ", ".join(map(str, allowed_values)))
            )

        self._payment_providers = payment_providers

    @property
    def stripe_sandbox_publishable_key(self):
        """Gets the stripe_sandbox_publishable_key of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The stripe_sandbox_publishable_key of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: str
        """
        return self._stripe_sandbox_publishable_key

    @stripe_sandbox_publishable_key.setter
    def stripe_sandbox_publishable_key(self, stripe_sandbox_publishable_key):
        """Sets the stripe_sandbox_publishable_key of this PatchBusinessSettingsRequest.


        :param stripe_sandbox_publishable_key: The stripe_sandbox_publishable_key of this PatchBusinessSettingsRequest.  # noqa: E501
        :type stripe_sandbox_publishable_key: str
        """

        self._stripe_sandbox_publishable_key = stripe_sandbox_publishable_key

    @property
    def stripe_sandbox_secret_key(self):
        """Gets the stripe_sandbox_secret_key of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The stripe_sandbox_secret_key of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: str
        """
        return self._stripe_sandbox_secret_key

    @stripe_sandbox_secret_key.setter
    def stripe_sandbox_secret_key(self, stripe_sandbox_secret_key):
        """Sets the stripe_sandbox_secret_key of this PatchBusinessSettingsRequest.


        :param stripe_sandbox_secret_key: The stripe_sandbox_secret_key of this PatchBusinessSettingsRequest.  # noqa: E501
        :type stripe_sandbox_secret_key: str
        """

        self._stripe_sandbox_secret_key = stripe_sandbox_secret_key

    @property
    def stripe_publishable_key(self):
        """Gets the stripe_publishable_key of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The stripe_publishable_key of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: str
        """
        return self._stripe_publishable_key

    @stripe_publishable_key.setter
    def stripe_publishable_key(self, stripe_publishable_key):
        """Sets the stripe_publishable_key of this PatchBusinessSettingsRequest.


        :param stripe_publishable_key: The stripe_publishable_key of this PatchBusinessSettingsRequest.  # noqa: E501
        :type stripe_publishable_key: str
        """

        self._stripe_publishable_key = stripe_publishable_key

    @property
    def stripe_secret_key(self):
        """Gets the stripe_secret_key of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The stripe_secret_key of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: str
        """
        return self._stripe_secret_key

    @stripe_secret_key.setter
    def stripe_secret_key(self, stripe_secret_key):
        """Sets the stripe_secret_key of this PatchBusinessSettingsRequest.


        :param stripe_secret_key: The stripe_secret_key of this PatchBusinessSettingsRequest.  # noqa: E501
        :type stripe_secret_key: str
        """

        self._stripe_secret_key = stripe_secret_key

    @property
    def paypal_client_id(self):
        """Gets the paypal_client_id of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The paypal_client_id of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: str
        """
        return self._paypal_client_id

    @paypal_client_id.setter
    def paypal_client_id(self, paypal_client_id):
        """Sets the paypal_client_id of this PatchBusinessSettingsRequest.


        :param paypal_client_id: The paypal_client_id of this PatchBusinessSettingsRequest.  # noqa: E501
        :type paypal_client_id: str
        """

        self._paypal_client_id = paypal_client_id

    @property
    def paypal_client_secret(self):
        """Gets the paypal_client_secret of this PatchBusinessSettingsRequest.  # noqa: E501


        :return: The paypal_client_secret of this PatchBusinessSettingsRequest.  # noqa: E501
        :rtype: str
        """
        return self._paypal_client_secret

    @paypal_client_secret.setter
    def paypal_client_secret(self, paypal_client_secret):
        """Sets the paypal_client_secret of this PatchBusinessSettingsRequest.


        :param paypal_client_secret: The paypal_client_secret of this PatchBusinessSettingsRequest.  # noqa: E501
        :type paypal_client_secret: str
        """

        self._paypal_client_secret = paypal_client_secret

    def to_dict(self, serialize=False):
        """Returns the model properties as a dict"""
        result = {}

        def convert(x):
            if hasattr(x, "to_dict"):
                args = getfullargspec(x.to_dict).args
                if len(args) == 1:
                    return x.to_dict()
                else:
                    return x.to_dict(serialize)
            else:
                return x

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            attr = self.attribute_map.get(attr, attr) if serialize else attr
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: convert(x),
                    value
                ))
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], convert(item[1])),
                    value.items()
                ))
            else:
                result[attr] = convert(value)

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, PatchBusinessSettingsRequest):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, PatchBusinessSettingsRequest):
            return True

        return self.to_dict() != other.to_dict()
