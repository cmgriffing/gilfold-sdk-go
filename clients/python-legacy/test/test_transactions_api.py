# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.0
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import openapi_client
from openapi_client.api.transactions_api import TransactionsApi  # noqa: E501
from openapi_client.rest import ApiException


class TestTransactionsApi(unittest.TestCase):
    """TransactionsApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.transactions_api.TransactionsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_transactions_get(self):
        """Test case for transactions_get

        Get all Transactions for a Business  # noqa: E501
        """
        pass

    def test_transactions_post(self):
        """Test case for transactions_post

        Creates a Transaction for an Business  # noqa: E501
        """
        pass

    def test_transactions_transaction_id_delete(self):
        """Test case for transactions_transaction_id_delete

        Delete Transaction for an Business  # noqa: E501
        """
        pass

    def test_transactions_transaction_id_get(self):
        """Test case for transactions_transaction_id_get

        Get a Transaction for a Business  # noqa: E501
        """
        pass

    def test_transactions_transaction_id_patch(self):
        """Test case for transactions_transaction_id_patch

        Update a Transaction for a Business  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
