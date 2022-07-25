=begin
#

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 0.0.2

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.1.0-SNAPSHOT

=end

require 'spec_helper'
require 'json'

# Unit tests for OpenapiClient::TransactionsApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'TransactionsApi' do
  before do
    # run before each test
    @api_instance = OpenapiClient::TransactionsApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of TransactionsApi' do
    it 'should create an instance of TransactionsApi' do
      expect(@api_instance).to be_instance_of(OpenapiClient::TransactionsApi)
    end
  end

  # unit tests for transactions_get
  # Get all Transactions for a Business
  # Get all Transactions for a Business
  # @param [Hash] opts the optional parameters
  # @return [AllTransactionsResponse]
  describe 'transactions_get test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for transactions_post
  # Creates a Transaction for an Business
  # Creates a Transaction for an Business
  # @param post_transaction_request 
  # @param [Hash] opts the optional parameters
  # @return [TransactionResponse]
  describe 'transactions_post test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for transactions_stats_get
  # Get stats for all Transactions
  # Get stats for all Transactions
  # @param [Hash] opts the optional parameters
  # @return [TransactionStatsResponse]
  describe 'transactions_stats_get test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for transactions_transaction_id_delete
  # Delete Transaction for an Business
  # Delete Transaction for an Business
  # @param [Hash] opts the optional parameters
  # @return [Object]
  describe 'transactions_transaction_id_delete test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for transactions_transaction_id_get
  # Get a Transaction for a Business
  # Get a Transaction for a Business
  # @param [Hash] opts the optional parameters
  # @return [TransactionResponse]
  describe 'transactions_transaction_id_get test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for transactions_transaction_id_patch
  # Update a Transaction for a Business
  # Update a Transaction for a Business
  # @param patch_transaction_request 
  # @param [Hash] opts the optional parameters
  # @return [TransactionResponse]
  describe 'transactions_transaction_id_patch test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end
