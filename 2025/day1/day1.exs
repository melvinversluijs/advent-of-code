defmodule Day1 do
    @starting_point 50
    @minimal_point 0
    @maximum_point 100

    def exercise1 do
        input = File.read!("./input.txt")

        {_, zero_hits} = input
          |> String.split("\n", trim: true)
          |> Enum.reduce({@starting_point, 0}, fn line, {current_point, zero_hits} ->
            <<direction::binary-size(1), numbers::binary>> = line
            points = String.to_integer(numbers)

            # The points can be more than the range permits.
            # But we only need the remainder of the division by the maximum.
            points = rem(points, @maximum_point)

            new_point = move(direction, points, current_point)

            zero_hits = if new_point === 0,
              do: zero_hits + 1,
              else: zero_hits

            {new_point, zero_hits}
          end)

          IO.puts(zero_hits)
    end

    defp move("L", points, current_point) do
      new_point = current_point - points

      # Are now allowed to go below the minumum. Take all that is left from the maximum.
      if new_point < @minimal_point,
        do: @maximum_point + new_point,
        else: new_point
    end

    defp move("R", points, current_point) do
      new_point = current_point + points

      # Are not allowed to go over the maximum. Take all that is left from the minumum
      if new_point >= @maximum_point,
        do: @minimal_point + (new_point - @maximum_point),
        else: new_point
    end
end

Day1.exercise1()
