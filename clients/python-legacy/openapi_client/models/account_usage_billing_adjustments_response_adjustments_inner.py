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


class AccountUsageBillingAdjustmentsResponseAdjustmentsInner(object):
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
        'billing_id': 'str',
        'account_id': 'str',
        'hourly_cost': 'float',
        'created_at': 'float',
        'modified_at': 'float'
    }

    attribute_map = {
        'billing_id': 'billingId',
        'account_id': 'accountId',
        'hourly_cost': 'hourlyCost',
        'created_at': 'createdAt',
        'modified_at': 'modifiedAt'
    }

    def __init__(self, billing_id=None, account_id=None, hourly_cost=None, created_at=None, modified_at=None, local_vars_configuration=None):  # noqa: E501
        """AccountUsageBillingAdjustmentsResponseAdjustmentsInner - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._billing_id = None
        self._account_id = None
        self._hourly_cost = None
        self._created_at = None
        self._modified_at = None
        self.discriminator = None

        self.billing_id = billing_id
        self.account_id = account_id
        self.hourly_cost = hourly_cost
        self.created_at = created_at
        self.modified_at = modified_at

    @property
    def billing_id(self):
        """Gets the billing_id of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501


        :return: The billing_id of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501
        :rtype: str
        """
        return self._billing_id

    @billing_id.setter
    def billing_id(self, billing_id):
        """Sets the billing_id of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.


        :param billing_id: The billing_id of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501
        :type billing_id: str
        """
        if self.local_vars_configuration.client_side_validation and billing_id is None:  # noqa: E501
            raise ValueError("Invalid value for `billing_id`, must not be `None`")  # noqa: E501

        self._billing_id = billing_id

    @property
    def account_id(self):
        """Gets the account_id of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501


        :return: The account_id of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501
        :rtype: str
        """
        return self._account_id

    @account_id.setter
    def account_id(self, account_id):
        """Sets the account_id of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.


        :param account_id: The account_id of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501
        :type account_id: str
        """
        if self.local_vars_configuration.client_side_validation and account_id is None:  # noqa: E501
            raise ValueError("Invalid value for `account_id`, must not be `None`")  # noqa: E501

        self._account_id = account_id

    @property
    def hourly_cost(self):
        """Gets the hourly_cost of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501


        :return: The hourly_cost of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501
        :rtype: float
        """
        return self._hourly_cost

    @hourly_cost.setter
    def hourly_cost(self, hourly_cost):
        """Sets the hourly_cost of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.


        :param hourly_cost: The hourly_cost of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501
        :type hourly_cost: float
        """
        if self.local_vars_configuration.client_side_validation and hourly_cost is None:  # noqa: E501
            raise ValueError("Invalid value for `hourly_cost`, must not be `None`")  # noqa: E501

        self._hourly_cost = hourly_cost

    @property
    def created_at(self):
        """Gets the created_at of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501


        :return: The created_at of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501
        :rtype: float
        """
        return self._created_at

    @created_at.setter
    def created_at(self, created_at):
        """Sets the created_at of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.


        :param created_at: The created_at of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501
        :type created_at: float
        """
        if self.local_vars_configuration.client_side_validation and created_at is None:  # noqa: E501
            raise ValueError("Invalid value for `created_at`, must not be `None`")  # noqa: E501

        self._created_at = created_at

    @property
    def modified_at(self):
        """Gets the modified_at of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501


        :return: The modified_at of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501
        :rtype: float
        """
        return self._modified_at

    @modified_at.setter
    def modified_at(self, modified_at):
        """Sets the modified_at of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.


        :param modified_at: The modified_at of this AccountUsageBillingAdjustmentsResponseAdjustmentsInner.  # noqa: E501
        :type modified_at: float
        """
        if self.local_vars_configuration.client_side_validation and modified_at is None:  # noqa: E501
            raise ValueError("Invalid value for `modified_at`, must not be `None`")  # noqa: E501

        self._modified_at = modified_at

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
        if not isinstance(other, AccountUsageBillingAdjustmentsResponseAdjustmentsInner):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, AccountUsageBillingAdjustmentsResponseAdjustmentsInner):
            return True

        return self.to_dict() != other.to_dict()
