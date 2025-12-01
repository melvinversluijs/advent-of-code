defmodule Day1 do
    @starting_point 50
    @minimal_point 0
    @maximum_point 100

    def exercise1 do
      {_, zero_hits} = get_input()
        |> Enum.reduce({@starting_point, 0}, fn line, {current_point, zero_hits} ->
          <<direction::binary-size(1), numbers::binary>> = line
          points = String.to_integer(numbers)

          if points === 0 do
            {current_point, zero_hits}
          else
            # The points can be more than the range permits.
            # But we only need the remainder of the division by the maximum.
            points = rem(points, @maximum_point)

            new_point = move(direction, points, current_point)

            zero_hits = if new_point === 0,
              do: zero_hits + 1,
              else: zero_hits

            {new_point, zero_hits}
          end
        end)

      IO.puts("Part 1: #{zero_hits}")
    end

    def exercise2 do
      {_, zero_hits} = get_input()
        |> Enum.reduce({@starting_point, 0}, fn line, {current_point, zero_hits} ->
          <<direction::binary-size(1), numbers::binary>> = line
          points = String.to_integer(numbers)

          # Calculate how many times it went past 0, and add it to the zero_hits.
          zero_hits = zero_hits + floor(points / @maximum_point)

          # Calculate what the remaining points is.
          points = rem(points, @maximum_point)

          if points === 0 do
            # If the points is zero no need to calculate anything else.
            {current_point, zero_hits}
          else
            new_point = move(direction, points, current_point)

            zero_hits = if has_hit_zero?(direction, current_point, new_point),
              do: zero_hits + 1,
              else: zero_hits

            {new_point, zero_hits}
          end
        end)

      IO.puts("Part 2: #{zero_hits}")
    end

    defp get_input() do
      File.read!("./input.txt")
        |> String.split("\n", trim: true)
    end

    defp move("L", points, current_point) when points <= 100 do
      new_point = current_point - points

      # Are now allowed to go below the minumum. Take all that is left from the maximum.
      if new_point < @minimal_point,
        do: @maximum_point + new_point,
        else: new_point
    end

    defp move("R", points, current_point) when points <= 100 do
      new_point = current_point + points

      # Are not allowed to go over the maximum. Take all that is left from the minumum
      if new_point >= @maximum_point,
        do: @minimal_point + (new_point - @maximum_point),
        else: new_point
    end

    defp has_hit_zero?(_direction, _current_point, 0), do: true

    defp has_hit_zero?(_direction, 0, _new_point), do: false

    defp has_hit_zero?("L", current_point, new_point),
      do: new_point > current_point

    defp has_hit_zero?("R", current_point, new_point),
      do: new_point < current_point
end

Day1.exercise1() # 980
Day1.exercise2() # 5961
