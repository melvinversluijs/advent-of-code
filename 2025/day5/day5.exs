defmodule Day5 do
  def part1 do
    {fresh_ingredient_ranges, ingredients} = get_ingredients()

    fresh_ingredient_count = ingredients
      |> Enum.count(&is_fresh?(&1, fresh_ingredient_ranges))

      IO.puts("Part 1: #{fresh_ingredient_count}")
  end

  defp get_ingredients do
    [fresh_ingredients, ingredients] = File.read!("input.txt")
			|> String.split("\n\n", trim: true)

		fresh_ingredient_ranges = parse_fresh_ingredient_ranges(fresh_ingredients)
		ingredients = parse_ingredients(ingredients)

		{fresh_ingredient_ranges, ingredients}
  end

  defp parse_fresh_ingredient_ranges(fresh_ingredients) do
    fresh_ingredients
		  |> String.split("\n", trim: true)
			|> Enum.map(fn range ->
			  range
					|> String.split("-", trim: true)
					|> Enum.map(&String.to_integer/1)
					|> then(fn [first, last] -> first..last end)
			end)
  end

  defp parse_ingredients(ingredients) do
    ingredients
		  |> String.split("\n", trim: true)
			|> Enum.map(&String.to_integer/1)
  end

  defp is_fresh?(ingredient, fresh_ingredient_ranges) do
   Enum.any?(fresh_ingredient_ranges, fn fresh_ingredient_range -> ingredient in fresh_ingredient_range end)
  end
end

Day5.part1() # 529
