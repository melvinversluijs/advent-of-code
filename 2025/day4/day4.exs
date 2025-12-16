defmodule Day4 do
	def part1 do
    liftable_rolls = get_grid()
      |> lift_one_round()

    IO.puts("Result Part 1: #{liftable_rolls}")
	end

	def part2 do
		{_grid, liftable_rolls} = get_grid()
		  |> lift_all_rolls()

		IO.puts("Result Part 2: #{liftable_rolls}")
	end

	defp get_grid do
	  File.read!("input.txt")
			|> String.split("\n", trim: true)
			|> Enum.with_index()
			|> Enum.flat_map(fn {line, y} ->
        String.graphemes(line)
          |> Enum.with_index()
          |> Enum.map(fn {char, x} -> {{x, y}, char} end)
			  end)
			|> Map.new()
	end

	defp lift_one_round(grid) do
	  grid
      |> Enum.count(fn {{x, y}, char} -> can_lift_roll(grid, x, y, char) end)
	end

	defp lift_all_rolls(grid, total_lifted \\ 0) do
    {new_grid, lifted_this_round} = lift_one_round_and_remove_from_grid(grid)

		if lifted_this_round > 0,
		  do: lift_all_rolls(new_grid, total_lifted + lifted_this_round),
			else: {new_grid, total_lifted}
	end

	defp lift_one_round_and_remove_from_grid(grid) do
	  grid
      |> Enum.reduce({grid, 0}, fn {{x, y}, char}, {new_grid, liftable_rolls} ->
          if can_lift_roll(new_grid, x, y, char),
            do: {remove_roll_from_grid(new_grid, {x, y}), liftable_rolls + 1},
            else: {new_grid, liftable_rolls}
      end)
	end

	defp can_lift_roll(_grid, _x, _y, char) when char != "@", do: false

	defp can_lift_roll(grid, x, y, "@") do
	  neighbour_rolls = get_neighbours(grid, x, y)
			|> Enum.count(&(&1 == "@"))

		neighbour_rolls < 4
	end

	defp get_neighbours(grid, x, y) do
  	[
      {x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}, # top row
      {x - 1, y},                 {x + 1, y},     # middle row
      {x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1}  # bottom row
    ]
      |> Enum.map(fn pos -> Map.get(grid, pos) end) # Get character at position
	    |> Enum.reject(&is_nil/1) # Remove non-existent positions
	end

	defp remove_roll_from_grid(grid, pos), do: Map.put(grid, pos, ".")
end

Day4.part1() # 1569
Day4.part2() # 9280
