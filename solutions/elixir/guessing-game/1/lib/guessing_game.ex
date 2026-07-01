defmodule GuessingGame do

  def compare(_), do: "Make a guess" 
  def compare(_, :no_guess), do: "Make a guess"
  def compare(s, g) when s==g, do: "Correct"
  def compare(s, g) when s-g==1, do: "So close"
  def compare(s, g) when g-s==1, do: "So close"
  def compare(s,g) when s<g, do: "Too high"
  def compare(s,g) when s>g, do: "Too low"

 end
