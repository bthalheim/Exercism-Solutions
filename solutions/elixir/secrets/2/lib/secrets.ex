defmodule Secrets do

  def secret_add(secret),      do: &(&1 + secret) 
  def secret_subtract(secret), do: &(&1 - secret)
  def secret_multiply(secret), do: fn param -> param * secret end
  def secret_divide(secret),   do: fn param -> div(param, secret) end

  def secret_and(secret), do: fn param -> Bitwise.band(param, secret) end
  def secret_xor(secret), do: fn param -> Bitwise.bxor(param, secret) end

  def secret_combine(secret_function1, secret_function2) do
    fn param ->
      next = secret_function1.(param)
      secret_function2.(next)
    end
  end


end
