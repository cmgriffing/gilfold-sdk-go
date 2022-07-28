# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule .Model.AllTransactionsResponseTransactions do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :Items,
    :Count,
    :ScannedCount,
    :LastEvaluatedKey
  ]

  @type t :: %__MODULE__{
    :Items => [.Model.TransactionResponse.t],
    :Count => float(),
    :ScannedCount => float(),
    :LastEvaluatedKey => %{optional(String.t) => String.t}
  }
end

defimpl Poison.Decoder, for: .Model.AllTransactionsResponseTransactions do
  import .Deserializer
  def decode(value, options) do
    value
    |> deserialize(:Items, :list, .Model.TransactionResponse, options)
  end
end

