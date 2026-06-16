=begin
Write your code for the 'Tournament' exercise in this file. Make the tests in
`tournament_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/tournament` directory.
=end


# Team                           | MP |  W |  D |  L |  P
#   Allegoric Alaskans             |  1 |  1 |  0 |  0 |  3
#   Blithering Badgers             |  1 |  0 |  0 |  1 |  0

# skip
#   input = <<~INPUT
#     Allegoric Alaskans;Blithering Badgers;win
#   INPUT

#   actual = Tournament.tally(input)


class Team

  attr_reader :name
  attr_accessor :wins, :losses, :draws, :matches


  def initialize(name)
    @name = name
    @matches = 0
    @wins = 0
    @losses = 0
    @draws = 0
  end

  def points
    wins * 3 + draws
  end
end


class Tournament

  def self.tally(input)

    teams = Hash.new do |hash, key|
      hash[key] = Team.new(key)
    end
    

    input.split(/\r?\n/).each do |line|

      team1_name, team2_name, result = line.split(";")

      team1 = teams[team1_name]
      team2 = teams[team2_name]

      team1.matches += 1
      team2.matches += 1 

      case result
      when "win"
        team1.wins   += 1
        team2.losses += 1
      when "loss"
        team1.losses += 1
        team2.wins   += 1
      when "draw"
        team1.draws += 1
        team2.draws += 1 
      end
      

    end

    return self._format(teams)

  end


  def self._format(teams)

    output = ""


# Team                           | MP |  W |  D |  L |  P
#   Allegoric Alaskans             |  1 |  1 |  0 |  0 |  3
#   Blithering Badgers             |  1 |  0 |  0 |  1 |  0

    output += "Team                           | MP |  W |  D |  L |  P\n"
  

    teams.values
          .sort_by { |team| [-team.points, team.name]}
          .each do |team|

          output += "#{team.name.ljust(31)}|  #{team.matches} |  #{team.wins} |  #{team.draws} |  #{team.losses} | #{team.points.to_s.rjust(2)}\n"
  
    end



    return output


  end


end



   
