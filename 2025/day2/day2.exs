defmodule Day2 do
  def part1 do
    result = get_input()
      |> Enum.flat_map(fn id ->
        [first, last] = String.split(id, "-", trim: true, parts: 2)
        first = String.to_integer(first)
        last = String.to_integer(last)

        Enum.filter(first..last, &is_invalid_character?/1)
      end)
      |> Enum.sum()

    IO.puts("Part 1: #{result}")
  end

  defp get_input() do
    File.read!("./input.txt")
      |> String.trim()
      |> String.split(",", trim: true)
  end

  defp is_invalid_character?(id) do
    # This regex checks if the second part is the same as the first part.
    Regex.run(~r/(^\d+)\1$/, Integer.to_string(id))
  end
end

Day2.part1()
