defmodule GuessingGame do

  def compare(s, g \\ :no_guess )  
  def compare(_, :no_guess), do: "Make a guess"
  def compare(s, g) when s==g, do: "Correct"
  def compare(s, g) when abs(s-g) == 1, do: "So close"
  def compare(s,g) when s<g, do: "Too high"
  def compare(s,g) when s>g, do: "Too low"

 end
