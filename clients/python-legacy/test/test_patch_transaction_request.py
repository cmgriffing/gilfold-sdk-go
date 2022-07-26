# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.3
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import openapi_client
from openapi_client.models.patch_transaction_request import PatchTransactionRequest  # noqa: E501
from openapi_client.rest import ApiException

class TestPatchTransactionRequest(unittest.TestCase):
    """PatchTransactionRequest unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test PatchTransactionRequest
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.patch_transaction_request.PatchTransactionRequest()  # noqa: E501
        if include_optional :
            return PatchTransactionRequest(
                account_id = '', 
                amount = 1.337, 
                note = '', 
                payment_provider = None, 
                payment_id = ''
            )
        else :
            return PatchTransactionRequest(
        )

    def testPatchTransactionRequest(self):
        """Test PatchTransactionRequest"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
