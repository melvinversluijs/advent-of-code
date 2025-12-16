defmodule Day4 do
	def part1 do
    grid = get_grid()

    liftable_rolls = grid
      |> Enum.count(fn {{x, y}, char} -> can_lift_roll(grid, x, y, char) end)

    IO.puts("Result Part 1: #{liftable_rolls}")
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
end

Day4.part1() # 1569
