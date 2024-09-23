defmodule Solution do
  @spec max_area(height :: [integer]) :: integer
  def max_area(height) do
    right = length(height) - 1

    height
    |> List.to_tuple()
    |> solve(0, right, 0)
  end

  defp solve(_, left, right, max_area) when left == right, do: max_area

  defp solve(height, left, right, max_area) do
    lh = elem(height, left)
    rh = elem(height, right)

    max_area = max(min(lh, rh) * (right - left), max_area)

    cond do
      lh < rh -> solve(height, left + 1, right, max_area)
      lh >= rh -> solve(height, left, right - 1, max_area)
    end
  end
end

defmodule SolutionTwo do
  @spec max_area(height :: [integer]) :: integer
  def max_area(height),
    do:
      height
      |> List.to_tuple()
      |> then(&max(&1, 0, tuple_size(&1) - 1, 0))

  defp max(heights, l, r, max_area) when l < r do
    lh = elem(heights, l)
    rh = elem(heights, r)

    max_area = max(min(lh, rh) * (r - l), max_area)

    cond do
      lh < rh -> max(heights, l + 1, r, max_area)
      lh >= rh -> max(heights, l, r - 1, max_area)
    end
  end

  defp max(_, _, _, max_area), do: max_area
end

Solution.max_area([1, 8, 6, 2, 5, 4, 8, 3, 7])
|> IO.puts()

SolutionTwo.max_area([1, 8, 6, 2, 5, 4, 8, 3, 7])
|> IO.puts()
