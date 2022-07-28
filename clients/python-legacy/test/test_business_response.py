# coding: utf-8

"""
    

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 0.0.11
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import openapi_client
from openapi_client.models.business_response import BusinessResponse  # noqa: E501
from openapi_client.rest import ApiException

class TestBusinessResponse(unittest.TestCase):
    """BusinessResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test BusinessResponse
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openapi_client.models.business_response.BusinessResponse()  # noqa: E501
        if include_optional :
            return BusinessResponse(
                business_id = '', 
                name = ''
            )
        else :
            return BusinessResponse(
                business_id = '',
                name = '',
        )

    def testBusinessResponse(self):
        """Test BusinessResponse"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
