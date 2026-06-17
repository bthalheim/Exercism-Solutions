module Acronym
  def self.abbreviate(phrase)
    output = ""

    phrase.gsub("-", " ")
          .gsub(",", " ")
          .gsub("_", "")
          .split(" ").each do |word|
      output += word[0].upcase
    end
    output
  end
end
