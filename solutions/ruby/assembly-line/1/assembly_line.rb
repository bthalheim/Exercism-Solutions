class AssemblyLine

  CAR_RATE = 221


  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour

    eff = 1.0

    case 
    when @speed.between?(5,8)
      eff = 0.9
    when @speed == 9
      eff = 0.8
    when @speed == 10
      eff = 0.77
    end

    (221 * @speed) * eff  

  end

  def working_items_per_minute
    (production_rate_per_hour / 60.0).floor

  end
end
