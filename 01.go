package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
  // Get the lists from stdin
  left_list, right_list := read_input()

  // Sort the lists
  slices.SortFunc(left_list, func(a, b int) int {
    return a - b
  })
  slices.SortFunc(right_list, func(a, b int) int {
    return a - b
  })

  // Calculate the differences
  diffs := calc_diffs(left_list, right_list)
  var total_diff int = 0
  for _, diff := range diffs {
    total_diff += diff
  }
  fmt.Printf("Difference: %d\n", total_diff)

  // Calculate the similarity score
  similarity := calc_similarity(left_list, right_list)
  fmt.Printf("Similarity Score: %d\n", similarity)
}

// Read two lists from standard input and return to slices
func read_input() ([]int, []int) {
  left_list := []int{}
  right_list := []int{}

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    var left, right int

    _, err := fmt.Sscanf(line, "%d\t%d", &left, &right)
    if err != nil {
    }

    left_list = append(left_list, left)
    right_list = append(right_list, right)
  }

  return left_list, right_list
}

// Calculate the difference between elements in the two lists
func calc_diffs(left_list []int, right_list []int) []int {
  diffs := []int{}

  for index := range left_list {
    diff := int(math.Abs(float64(left_list[index] - right_list[index])))
    diffs = append(diffs, diff)
  }

  return diffs
}

// Calculate a similarity score
func calc_similarity(left_list []int, right_list []int) int {
  var similarity int = 0

  for _, left := range left_list {
    var occurrences int = 0

    for _, right := range right_list {
      if left == right {
        occurrences += 1
      }
    }

    similarity += (left * occurrences)
  }

  return similarity
}

