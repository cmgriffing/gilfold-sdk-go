# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.2
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import openapi_client
from openapi_client.api.usage_api import UsageApi  # noqa: E501
from openapi_client.rest import ApiException


class TestUsageApi(unittest.TestCase):
    """UsageApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.usage_api.UsageApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_accounts_account_id_usage_get(self):
        """Test case for accounts_account_id_usage_get

        Get an Account's current Usage Billing Rate Adjustments  # noqa: E501
        """
        pass

    def test_accounts_account_id_usage_post(self):
        """Test case for accounts_account_id_usage_post

        Set an Account's Usage Billing Rate  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
