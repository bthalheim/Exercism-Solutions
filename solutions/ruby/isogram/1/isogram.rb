




  # def test_empty_string
  #   # skip
  #   input = ''
  #   assert Isogram.isogram?(input), "Expected true, '#{input}' is an isogram"
  # end



module Isogram

  def self.isogram?(input)


    letters = Hash.new do |hash, key|
      hash[key] = 0
    end

    input.downcase.each_char do |c|
      if c.match(/[[:alpha:]]/)
        letters[c] += 1
      end
    end

    letters.values.each do |v|
      if v > 1
        return false
      end  
    end

    true    

  end
  
end




