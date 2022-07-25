=begin
#

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 0.0.2

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.1.0-SNAPSHOT

=end

require 'spec_helper'
require 'json'

# Unit tests for OpenapiClient::UsageApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'UsageApi' do
  before do
    # run before each test
    @api_instance = OpenapiClient::UsageApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of UsageApi' do
    it 'should create an instance of UsageApi' do
      expect(@api_instance).to be_instance_of(OpenapiClient::UsageApi)
    end
  end

  # unit tests for accounts_account_id_usage_get
  # Get an Account&#39;s current Usage Billing Rate Adjustments
  # Get an Account&#39;s current Usage Billing Rate Adjustments
  # @param [Hash] opts the optional parameters
  # @return [AccountUsageBillingAdjustmentsResponse]
  describe 'accounts_account_id_usage_get test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for accounts_account_id_usage_post
  # Set an Account&#39;s Usage Billing Rate
  # Set an Account&#39;s Usage Billing Rate
  # @param post_account_usage_billing_rate_adjustment_request 
  # @param [Hash] opts the optional parameters
  # @return [Object]
  describe 'accounts_account_id_usage_post test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
