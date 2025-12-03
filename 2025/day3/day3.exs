defmodule Day3 do
  def part1 do
    result = get_banks()
      |> Enum.map(fn bank ->
        String.graphemes(bank)
        |> find_largest_number()
      end)
      |> Enum.sum()

    IO.puts("Part 1: #{result}")
  end

  defp get_banks() do
    File.read!("input.txt")
      |> String.split("\n", trim: true)
  end

  # When the array is empty.
  defp find_largest_number([]), do: 0

  # Case when the rest part is empty. So then last character is by default the largest one.
  defp find_largest_number([current_char]), do: String.to_integer(current_char)

  # Recursively find the largest number.
  defp find_largest_number([current_char | rest]) do
    largest_number = Enum.map(rest, fn character ->
      String.to_integer(current_char <> character)
    end)
    |> Enum.max()

    max(largest_number, find_largest_number(rest))
  end
end

Day3.part1() # 16946
