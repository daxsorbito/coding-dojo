package minesweeper

import (
	"reflect"
	"testing"
)

func Test_solve_exceptionRules(t *testing.T) {
	tests := []struct {
		name          string
		givenGridSize []int
		givenMines    []string
		want          []string
	}{
		{
			name:          "[0,0] will return empty response",
			givenGridSize: []int{0, 0},
			givenMines:    []string{""},
			want:          []string{},
		},
		{
			name:          "Empty mine list (nil) returns a grid of \".\"",
			givenGridSize: []int{2, 2},
			givenMines:    nil,
			want:          []string{"..", ".."},
		},
		{
			name:          "Empty mine list (empty array) with 2x2 gridSize returns a grid of \".\"",
			givenGridSize: []int{2, 2},
			givenMines:    []string{},
			want:          []string{"..", ".."},
		},
		{
			name:          "Empty mine list (empty array) with 4x3 gridSize returns a grid of \".\"",
			givenGridSize: []int{4, 3},
			givenMines:    []string{},
			want:          []string{"....", "....", "...."},
		},

		{
			name:          "Empty mine list (array without mines) with 4x3 gridSize returns a grid of \".\"",
			givenGridSize: []int{4, 3},
			givenMines:    []string{"....", "....", "...."},
			want:          []string{"....", "....", "...."},
		},

		{
			name:          "list with all mines with 4x3 gridSize returns a grid of \".\"",
			givenGridSize: []int{4, 3},
			givenMines:    []string{"****", "****", "****"},
			want:          []string{"****", "****", "****"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.givenGridSize, tt.givenMines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_expectedResult(t *testing.T) {
	tests := []struct {
		name          string
		givenGridSize []int
		givenMines    []string
		want          []string
	}{
		{
			name:          "Given gridSize[3,2] and mines ['.*.', '...']",
			givenGridSize: []int{3, 2},
			givenMines: []string{
				".*.",
				"...",
			},
			want: []string{
				"1*1",
				"111",
			},
		},
		{
			name:          "Given gridSize[5,3] and 3 mines returns expected result",
			givenGridSize: []int{5, 3},
			givenMines: []string{
				"**...",
				".....",
				".*...",
			},
			want: []string{
				"**1..",
				"332..",
				"1*1..",
			},
		},
		{
			name:          "Given gridSize[12,6] and 1 mines returns expected result",
			givenGridSize: []int{12, 6},
			givenMines: []string{
				"............",
				"............",
				"............",
				"....*.......",
				"............",
				"............",
			},
			want: []string{
				"............",
				"............",
				"...111......",
				"...1*1......",
				"...111......",
				"............",
			},
		},
		{
			name:          "Given gridSize[12,6] and 3 mines returns expected result",
			givenGridSize: []int{12, 6},
			givenMines: []string{
				"............",
				"............",
				"......*.....",
				"....*.......",
				"......*.....",
				"............"},
			want: []string{
				"............",
				".....111....",
				"...112*1....",
				"...1*322....",
				"...112*1....",
				".....111....",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.givenGridSize, tt.givenMines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
