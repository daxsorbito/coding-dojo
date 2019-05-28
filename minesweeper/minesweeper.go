package minesweeper

import (
	"strconv"
	"strings"
)

func solve(gridSize []int, mines []string) []string {
	cols, rows := gridSize[0], gridSize[1]
	if rows == 0 && cols == 0 {
		return []string{}
	}

	result := make([]string, rows)

	if mines == nil || len(mines) == 0 {
		for x := 0; x < rows; x++ {
			result[x] = strings.Repeat(".", cols)
		}

		return result
	}

	for r := 0; r < rows; r++ {
		colResult := make([]string, cols)
		for c := 0; c < cols; c++ {
			currentValue := string(mines[r][c])
			switch currentValue {
			case "*":
				colResult[c] = "*"
			default:
				colResult[c] = checkSurroundingValues(mines, r, c)
			}

		}
		result[r] = strings.Join(colResult, "")
	}

	return result
}

// checkSurroundingValues is called when currentValue (at caller) is "."
// Hence to find out how many mines are surrounding the current position.
func checkSurroundingValues(mines []string, currentRow, currentCol int) string {

	// figure out the row (index) to start with and end with
	// figure out the col (index) to start with and end with
	rowStart, rowEnd := currentRow-1, currentRow+1
	colStart, colEnd := currentCol-1, currentCol+1

	var minesCount int
	// technically, rowStart might be -1 or above
	for r := rowStart; r <= rowEnd && r < len(mines); r++ {
		if r < 0 {
			continue
		}
		for c := colStart; c <= colEnd && c < len(mines[r]); c++ {
			if c < 0 || (r == currentRow && c == currentCol) {
				continue
			}

			value := string(mines[r][c])
			if value == "*" {
				minesCount = minesCount + 1
			}
		}
	}

	if minesCount > 0 {
		return strconv.Itoa(minesCount)
	}
	return "."
}
