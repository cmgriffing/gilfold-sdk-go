# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule .Api.Accounts do
  @moduledoc """
  API calls for all endpoints tagged `Accounts`.
  """

  alias .Connection
  import .RequestBuilder


  @doc """
  Delete an Account of a Business
  Delete an Account of a Business

  ## Parameters

  - connection (.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, map()} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec accounts_account_id_delete(Tesla.Env.client, keyword()) :: {:ok, Map.t} | {:error, Tesla.Env.t}
  def accounts_account_id_delete(connection, _opts \\ []) do
    %{}
    |> method(:delete)
    |> url("/accounts/:accountId")
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %{}},
      { 401, %{}},
      { 404, %{}}
    ])
  end

  @doc """
  Get an account for a Business
  Get an account for a Business

  ## Parameters

  - connection (.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, .Model.AccountResponse.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec accounts_account_id_get(Tesla.Env.client, keyword()) :: {:ok, Map.t} | {:ok, .Model.AccountResponse.t} | {:error, Tesla.Env.t}
  def accounts_account_id_get(connection, _opts \\ []) do
    %{}
    |> method(:get)
    |> url("/accounts/:accountId")
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %.Model.AccountResponse{}},
      { 401, %{}},
      { 404, %{}}
    ])
  end

  @doc """
  Update Account for Business
  Update Account for Business

  ## Parameters

  - connection (.Connection): Connection to server
  - patch_account_request (PatchAccountRequest): 
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, .Model.AccountResponse.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec accounts_account_id_patch(Tesla.Env.client, .Model.PatchAccountRequest.t, keyword()) :: {:ok, Map.t} | {:ok, .Model.AccountResponse.t} | {:error, Tesla.Env.t}
  def accounts_account_id_patch(connection, patch_account_request, _opts \\ []) do
    %{}
    |> method(:patch)
    |> url("/accounts/:accountId")
    |> add_param(:body, :body, patch_account_request)
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %.Model.AccountResponse{}},
      { 400, %{}},
      { 401, %{}}
    ])
  end

  @doc """
  Get Transactions for an Account
  Get Transactions for an Account

  ## Parameters

  - connection (.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, .Model.AllTransactionsResponse.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec accounts_account_id_transactions_get(Tesla.Env.client, keyword()) :: {:ok, .Model.AllTransactionsResponse.t} | {:ok, Map.t} | {:error, Tesla.Env.t}
  def accounts_account_id_transactions_get(connection, _opts \\ []) do
    %{}
    |> method(:get)
    |> url("/accounts/:accountId/transactions")
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %.Model.AllTransactionsResponse{}},
      { 401, %{}},
      { 404, %{}}
    ])
  end

  @doc """
  Get all Accounts for a Business
  Get all Accounts for a Business

  ## Parameters

  - connection (.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, .Model.AllAccountsResponse.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec accounts_get(Tesla.Env.client, keyword()) :: {:ok, .Model.AllAccountsResponse.t} | {:ok, Map.t} | {:error, Tesla.Env.t}
  def accounts_get(connection, _opts \\ []) do
    %{}
    |> method(:get)
    |> url("/accounts")
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %.Model.AllAccountsResponse{}},
      { 401, %{}}
    ])
  end

  @doc """
  Create an Account for a Business
  Create an Account for a Business

  ## Parameters

  - connection (.Connection): Connection to server
  - post_account_request (PostAccountRequest): 
  - opts (KeywordList): [optional] Optional parameters
  ## Returns

  {:ok, .Model.AccountResponse.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec accounts_post(Tesla.Env.client, .Model.PostAccountRequest.t, keyword()) :: {:ok, Map.t} | {:ok, .Model.AccountResponse.t} | {:error, Tesla.Env.t}
  def accounts_post(connection, post_account_request, _opts \\ []) do
    %{}
    |> method(:post)
    |> url("/accounts")
    |> add_param(:body, :body, post_account_request)
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 200, %.Model.AccountResponse{}},
      { 400, %{}},
      { 401, %{}},
      { 409, %{}}
    ])
  end
end
