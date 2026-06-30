

defmodule LogLevel do

  def to_label(level, legacy?) do
    cond do
       legacy? -> to_label_legacy(level)
      !legacy? -> to_label_modern(level)
    end
  end

  def to_label_modern(level) do

    cond do
        level == 0 -> :trace
        level == 1 -> :debug
        level == 2 -> :info
        level == 3 -> :warning
        level == 4 -> :error
        level == 5 -> :fatal
        true -> :unknown
    end
      
  end

  def to_label_legacy(level) do

    cond do
       level == 1 -> :debug
       level == 2 -> :info
       level == 3 -> :warning
       level == 4 -> :error
       true -> :unknown
    end

  end



  def alert_recipient(level, legacy?) do

    case to_label(level, legacy?) do
      :error -> :ops
      :fatal -> :ops
      :unknown when legacy? -> :dev1
      :unknown  -> :dev2
      _ -> false
    end


  end


end





