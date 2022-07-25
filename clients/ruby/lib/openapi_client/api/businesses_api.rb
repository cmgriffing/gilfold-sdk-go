=begin
#

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 0.0.2

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.1.0-SNAPSHOT

=end

require 'cgi'

module OpenapiClient
  class BusinessesApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # Get Settings for a business
    # Get Settings for a business
    # @param [Hash] opts the optional parameters
    # @return [BusinessSettingsResponse]
    def business_business_id_settings_get(opts = {})
      data, _status_code, _headers = business_business_id_settings_get_with_http_info(opts)
      data
    end

    # Get Settings for a business
    # Get Settings for a business
    # @param [Hash] opts the optional parameters
    # @return [Array<(BusinessSettingsResponse, Integer, Hash)>] BusinessSettingsResponse data, response status code and response headers
    def business_business_id_settings_get_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: BusinessesApi.business_business_id_settings_get ...'
      end
      # resource path
      local_var_path = '/business/:businessId/settings'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'BusinessSettingsResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['ApiKeyAuth']

      new_options = opts.merge(
        :operation => :"BusinessesApi.business_business_id_settings_get",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: BusinessesApi#business_business_id_settings_get\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Get Business Details
    #  Use this to get the name and other business details. Here is a link: [to google](https://google.com). Let's see if it works. 
    # @param [Hash] opts the optional parameters
    # @return [BusinessResponse]
    def businesses_business_id_get(opts = {})
      data, _status_code, _headers = businesses_business_id_get_with_http_info(opts)
      data
    end

    # Get Business Details
    #  Use this to get the name and other business details. Here is a link: [to google](https://google.com). Let&#39;s see if it works. 
    # @param [Hash] opts the optional parameters
    # @return [Array<(BusinessResponse, Integer, Hash)>] BusinessResponse data, response status code and response headers
    def businesses_business_id_get_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: BusinessesApi.businesses_business_id_get ...'
      end
      # resource path
      local_var_path = '/businesses/:businessId'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'BusinessResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['ApiKeyAuth']

      new_options = opts.merge(
        :operation => :"BusinessesApi.businesses_business_id_get",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: BusinessesApi#businesses_business_id_get\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # Update Business Details
    # Update Business Details
    # @param patch_business_request [PatchBusinessRequest] 
    # @param [Hash] opts the optional parameters
    # @return [BusinessResponse]
    def businesses_business_id_patch(patch_business_request, opts = {})
      data, _status_code, _headers = businesses_business_id_patch_with_http_info(patch_business_request, opts)
      data
    end

    # Update Business Details
    # Update Business Details
    # @param patch_business_request [PatchBusinessRequest] 
    # @param [Hash] opts the optional parameters
    # @return [Array<(BusinessResponse, Integer, Hash)>] BusinessResponse data, response status code and response headers
    def businesses_business_id_patch_with_http_info(patch_business_request, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: BusinessesApi.businesses_business_id_patch ...'
      end
      # verify the required parameter 'patch_business_request' is set
      if @api_client.config.client_side_validation && patch_business_request.nil?
        fail ArgumentError, "Missing the required parameter 'patch_business_request' when calling BusinessesApi.businesses_business_id_patch"
      end
      # resource path
      local_var_path = '/businesses/:businessId'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(patch_business_request)

      # return_type
      return_type = opts[:debug_return_type] || 'BusinessResponse'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['ApiKeyAuth']

      new_options = opts.merge(
        :operation => :"BusinessesApi.businesses_business_id_patch",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PATCH, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: BusinessesApi#businesses_business_id_patch\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
