package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type report struct {
  levels []int
}

func main() {
  reports := read_input()
  
  safe_count := 0
  for _, report := range reports {
    if is_report_safe(report) {
      safe_count += 1
    }
  }
  fmt.Printf("Safe Count: %d\n", safe_count)

  safe_count_with_dampener := 0
  for _, report := range reports {
    report_safe := is_report_safe(report)
    if report_safe {
      safe_count_with_dampener += 1
    } else {
      // Yuck!
      for index := range report.levels {
        if is_report_safe(remove_level_from_report(report, index)) {
          safe_count_with_dampener += 1
          break
        }
      }
    }
  }
  fmt.Printf("Safe Count with Dampener: %d\n", safe_count_with_dampener)
}

func is_report_safe(report report) bool {
  dirs := []int{}
  for index, level := range report.levels {
    if (index == 0) { continue }

    diff := level - report.levels[index - 1]
    abs_diff := int(math.Abs(float64(diff)))

    // Case where the difference is too large
    if abs_diff < 1 || abs_diff > 3 {
      return false
    }

    dir := diff / abs_diff

    // Case where the direction changes
    if len(dirs) > 0 && dirs[len(dirs) - 1] != 0 && dir != dirs[len(dirs) - 1] {
        return false
    }

    dirs = append(dirs, dir)
  }

  return true
}

func remove_level_from_report(original_report report, level_index int) report {
  new_report := report{
    levels: []int{},
  }

  for index, level := range original_report.levels {
    if index == level_index { continue }
    new_report.levels = append(new_report.levels, level)
  }

  return new_report
}

func read_input() []report {
  report_list := []report{}

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    report := report{}
    report.levels = get_levels_for_line(scanner.Text())
    report_list = append(report_list, report)
  }

  return report_list
}

func get_levels_for_line(line string) []int {
  fields := strings.Fields(line)
  levels := []int{}

  for _, level_str := range fields {
    level, err := strconv.Atoi(level_str)
    if err != nil {
    }
    levels = append(levels, level)
  }

  return levels
}

