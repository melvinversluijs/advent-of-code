defmodule Day2 do
  def part1 do
    result =
      get_input()
      |> Enum.flat_map(fn id_range ->
        Enum.filter(get_range_for_id_range(id_range), fn id ->
          # This regex checks if the second part is the same as the first part.
          Regex.run(~r/(^\d+)\1$/, Integer.to_string(id))
        end)
      end)
      |> Enum.sum()

    IO.puts("Part 1: #{result}")
  end

  def part2 do
    result =
      get_input()
      |> Enum.flat_map(fn id_range ->
        Enum.filter(get_range_for_id_range(id_range), fn id ->
          # This regex checks if the first match exists more than once.
          Regex.run(~r/^(\d+)\1+$/, Integer.to_string(id))
        end)
      end)
      |> Enum.sum()

    IO.puts("Part 2: #{result}")
  end

  defp get_input() do
    File.read!("./input.txt")
    |> String.trim()
    |> String.split(",", trim: true)
  end

  defp get_range_for_id_range(id) when is_binary(id) do
    [first, last] = String.split(id, "-", trim: true, parts: 2)
    first = String.to_integer(first)
    last = String.to_integer(last)

    first..last
  end
end

Day2.part1() # 41294979841
Day2.part2() # 66500947346
