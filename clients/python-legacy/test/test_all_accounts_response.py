# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.9
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import openapi_client
from openapi_client.models.all_accounts_response import AllAccountsResponse  # noqa: E501
from openapi_client.rest import ApiException

class TestAllAccountsResponse(unittest.TestCase):
    """AllAccountsResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test AllAccountsResponse
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.all_accounts_response.AllAccountsResponse()  # noqa: E501
        if include_optional :
            return AllAccountsResponse(
                accounts = openapi_client.models.all_accounts_response_accounts.AllAccountsResponse_accounts(
                    items = openapi_client.models.account_response.AccountResponse(
                        account_id = '', 
                        name = '', 
                        created_at = 1.337, 
                        modified_at = 1.337, 
                        usage_billing_rate = 1.337, ), 
                    count = 1.337, 
                    scanned_count = 1.337, 
                    last_evaluated_key = {
                        'key' : ''
                        }, ), 
                total = 1.337
            )
        else :
            return AllAccountsResponse(
                accounts = openapi_client.models.all_accounts_response_accounts.AllAccountsResponse_accounts(
                    items = openapi_client.models.account_response.AccountResponse(
                        account_id = '', 
                        name = '', 
                        created_at = 1.337, 
                        modified_at = 1.337, 
                        usage_billing_rate = 1.337, ), 
                    count = 1.337, 
                    scanned_count = 1.337, 
                    last_evaluated_key = {
                        'key' : ''
                        }, ),
                total = 1.337,
        )

    def testAllAccountsResponse(self):
        """Test AllAccountsResponse"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
