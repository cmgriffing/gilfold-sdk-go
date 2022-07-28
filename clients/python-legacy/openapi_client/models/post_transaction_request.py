# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.11
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


class PostTransactionRequest(object):
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
        'account_id': 'str',
        'amount': 'float',
        'note': 'str',
        'payment_provider': 'PaymentProvider',
        'payment_id': 'str'
    }

    attribute_map = {
        'account_id': 'accountId',
        'amount': 'amount',
        'note': 'note',
        'payment_provider': 'paymentProvider',
        'payment_id': 'paymentId'
    }

    def __init__(self, account_id=None, amount=None, note=None, payment_provider=None, payment_id=None, local_vars_configuration=None):  # noqa: E501
        """PostTransactionRequest - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._account_id = None
        self._amount = None
        self._note = None
        self._payment_provider = None
        self._payment_id = None
        self.discriminator = None

        self.account_id = account_id
        self.amount = amount
        if note is not None:
            self.note = note
        if payment_provider is not None:
            self.payment_provider = payment_provider
        if payment_id is not None:
            self.payment_id = payment_id

    @property
    def account_id(self):
        """Gets the account_id of this PostTransactionRequest.  # noqa: E501


        :return: The account_id of this PostTransactionRequest.  # noqa: E501
        :rtype: str
        """
        return self._account_id

    @account_id.setter
    def account_id(self, account_id):
        """Sets the account_id of this PostTransactionRequest.


        :param account_id: The account_id of this PostTransactionRequest.  # noqa: E501
        :type account_id: str
        """
        if self.local_vars_configuration.client_side_validation and account_id is None:  # noqa: E501
            raise ValueError("Invalid value for `account_id`, must not be `None`")  # noqa: E501

        self._account_id = account_id

    @property
    def amount(self):
        """Gets the amount of this PostTransactionRequest.  # noqa: E501


        :return: The amount of this PostTransactionRequest.  # noqa: E501
        :rtype: float
        """
        return self._amount

    @amount.setter
    def amount(self, amount):
        """Sets the amount of this PostTransactionRequest.


        :param amount: The amount of this PostTransactionRequest.  # noqa: E501
        :type amount: float
        """
        if self.local_vars_configuration.client_side_validation and amount is None:  # noqa: E501
            raise ValueError("Invalid value for `amount`, must not be `None`")  # noqa: E501

        self._amount = amount

    @property
    def note(self):
        """Gets the note of this PostTransactionRequest.  # noqa: E501


        :return: The note of this PostTransactionRequest.  # noqa: E501
        :rtype: str
        """
        return self._note

    @note.setter
    def note(self, note):
        """Sets the note of this PostTransactionRequest.


        :param note: The note of this PostTransactionRequest.  # noqa: E501
        :type note: str
        """

        self._note = note

    @property
    def payment_provider(self):
        """Gets the payment_provider of this PostTransactionRequest.  # noqa: E501


        :return: The payment_provider of this PostTransactionRequest.  # noqa: E501
        :rtype: PaymentProvider
        """
        return self._payment_provider

    @payment_provider.setter
    def payment_provider(self, payment_provider):
        """Sets the payment_provider of this PostTransactionRequest.


        :param payment_provider: The payment_provider of this PostTransactionRequest.  # noqa: E501
        :type payment_provider: PaymentProvider
        """

        self._payment_provider = payment_provider

    @property
    def payment_id(self):
        """Gets the payment_id of this PostTransactionRequest.  # noqa: E501


        :return: The payment_id of this PostTransactionRequest.  # noqa: E501
        :rtype: str
        """
        return self._payment_id

    @payment_id.setter
    def payment_id(self, payment_id):
        """Sets the payment_id of this PostTransactionRequest.


        :param payment_id: The payment_id of this PostTransactionRequest.  # noqa: E501
        :type payment_id: str
        """

        self._payment_id = payment_id

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
        if not isinstance(other, PostTransactionRequest):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, PostTransactionRequest):
            return True

        return self.to_dict() != other.to_dict()
