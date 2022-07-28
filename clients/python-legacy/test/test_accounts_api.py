# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.9
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import openapi_client
from openapi_client.api.accounts_api import AccountsApi  # noqa: E501
from openapi_client.rest import ApiException


class TestAccountsApi(unittest.TestCase):
    """AccountsApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.accounts_api.AccountsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_accounts_account_id_delete(self):
        """Test case for accounts_account_id_delete

        Delete an Account of a Business  # noqa: E501
        """
        pass

    def test_accounts_account_id_get(self):
        """Test case for accounts_account_id_get

        Get an account for a Business  # noqa: E501
        """
        pass

    def test_accounts_account_id_patch(self):
        """Test case for accounts_account_id_patch

        Update Account for Business  # noqa: E501
        """
        pass

    def test_accounts_account_id_transactions_get(self):
        """Test case for accounts_account_id_transactions_get

        Get Transactions for an Account  # noqa: E501
        """
        pass

    def test_accounts_get(self):
        """Test case for accounts_get

        Get all Accounts for a Business  # noqa: E501
        """
        pass

    def test_accounts_post(self):
        """Test case for accounts_post

        Create an Account for a Business  # noqa: E501
        """
        pass

    def test_accounts_stats_get(self):
        """Test case for accounts_stats_get

        Get stats for all accounts  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
