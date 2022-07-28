# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.9
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


class AllTransactionsResponse(object):
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
        'transactions': 'AllTransactionsResponseTransactions',
        'total': 'float'
    }

    attribute_map = {
        'transactions': 'transactions',
        'total': 'total'
    }

    def __init__(self, transactions=None, total=None, local_vars_configuration=None):  # noqa: E501
        """AllTransactionsResponse - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._transactions = None
        self._total = None
        self.discriminator = None

        self.transactions = transactions
        self.total = total

    @property
    def transactions(self):
        """Gets the transactions of this AllTransactionsResponse.  # noqa: E501


        :return: The transactions of this AllTransactionsResponse.  # noqa: E501
        :rtype: AllTransactionsResponseTransactions
        """
        return self._transactions

    @transactions.setter
    def transactions(self, transactions):
        """Sets the transactions of this AllTransactionsResponse.


        :param transactions: The transactions of this AllTransactionsResponse.  # noqa: E501
        :type transactions: AllTransactionsResponseTransactions
        """
        if self.local_vars_configuration.client_side_validation and transactions is None:  # noqa: E501
            raise ValueError("Invalid value for `transactions`, must not be `None`")  # noqa: E501

        self._transactions = transactions

    @property
    def total(self):
        """Gets the total of this AllTransactionsResponse.  # noqa: E501


        :return: The total of this AllTransactionsResponse.  # noqa: E501
        :rtype: float
        """
        return self._total

    @total.setter
    def total(self, total):
        """Sets the total of this AllTransactionsResponse.


        :param total: The total of this AllTransactionsResponse.  # noqa: E501
        :type total: float
        """
        if self.local_vars_configuration.client_side_validation and total is None:  # noqa: E501
            raise ValueError("Invalid value for `total`, must not be `None`")  # noqa: E501

        self._total = total

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
        if not isinstance(other, AllTransactionsResponse):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, AllTransactionsResponse):
            return True

        return self.to_dict() != other.to_dict()
