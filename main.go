package main

import (
	"fmt"
)

type Value struct {
	i            int
	v            int
	cellIndex    int
	indexInCell  int
	vlineIndex   int
	indexInVline int
	hlineIndex   int
	indexInHline int
	s            *Sudoku
}

// Cell structure
type Cell struct {
	i             int
	unknownCount  int
	values        [6]*Value
	containValues [6]*struct{}
}

type Hline struct {
	i             int
	unknownCount  int
	values        [6]*Value
	containValues [6]*struct{}
}

type Vline struct {
	i             int
	unknownCount  int
	values        [6]*Value
	containValues [6]*struct{}
}

func main() {
	su := NewSudoku()
	fmt.Println(su)
}

// Sudoku main game structure
type Sudoku struct {
	init            bool
	v               [36]*Value
	cells           [6]Cell
	cellsIndexInfo  [6][6]int
	hlines          [6]Hline
	hlinesIndexInfo [6][6]int
	vlines          [6]Vline
	vlinesIndexInfo [6][6]int
}

func NewSudoku() *Sudoku {
	su := Sudoku{}
	su.Generate()
	return &su
}

func (s *Sudoku) Init() {
	s.cellsIndexInfo = [6][6]int{
		{0, 1, 2, 6, 7, 8},
		{3, 4, 5, 9, 10},
		{12, 13, 14, 18, 19, 20},
		{15, 16, 17, 21, 22, 23},
		{24, 25, 26, 30, 31, 32},
		{27, 28, 29, 33, 34, 35},
	}
	s.hlinesIndexInfo = [6][6]int{
		{0, 1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10, 11},
		{12, 13, 14, 15, 16, 17},
		{18, 19, 20, 21, 22, 23},
		{24, 25, 26, 27, 28, 29},
		{30, 31, 32, 33, 34, 35},
	}
	s.vlinesIndexInfo = [6][6]int{
		{0, 6, 12, 18, 24, 30},
		{1, 7, 13, 19, 25, 31},
		{2, 8, 14, 20, 26, 32},
		{3, 9, 15, 21, 27, 33},
		{4, 10, 16, 22, 28, 34},
		{5, 11, 17, 23, 29, 35},
	}

	s.Fill()
	s.init = true
}

func (s *Sudoku) Fill() {
	for i := 0; i < 36; i++ {
		v := Value{
			i: i,
			s: s,
		}
		s.v[i] = &v

		{
			v.cellIndex, v.indexInCell = s.GetCellIndexes(i)
			s.cells[v.cellIndex].i = v.cellIndex
			s.cells[v.cellIndex].unknownCount = 6
			s.cells[v.cellIndex].values[v.indexInCell] = &v
		}

		{
			v.vlineIndex, v.indexInVline = s.GetVlineIndexes(i)
			s.vlines[v.vlineIndex].i = v.vlineIndex
			s.vlines[v.vlineIndex].unknownCount = 6
			s.vlines[v.vlineIndex].values[v.indexInVline] = &v
		}

		{
			v.hlineIndex, v.indexInHline = s.GetHlineIndexes(i)
			s.hlines[v.hlineIndex].i = v.hlineIndex
			s.hlines[v.hlineIndex].unknownCount = 6
			s.hlines[v.hlineIndex].values[v.indexInHline] = &v
		}
	}
}

// GetCellIndexes returns cellIndexes
func (s *Sudoku) GetCellIndexes(i int) (objectIndex int, valueIndex int) {
	return s.SearchIndex(i, s.cellsIndexInfo)
}

// GetVlineIndexes gets vlineIndex index by common slice of integers
func (s *Sudoku) GetVlineIndexes(i int) (objectIndex int, valueIndex int) {
	return s.SearchIndex(i, s.vlinesIndexInfo)
}

// GetHlineIndexes gets hlineIndex index by common slice of integers
func (s *Sudoku) GetHlineIndexes(i int) (objectIndex int, valueIndex int) {
	return s.SearchIndex(i, s.hlinesIndexInfo)
}

func (s *Sudoku) SearchIndex(index int, arr [6][6]int) (objectIndex int, valueIndex int) {
	for i := range arr {
		for j := range arr[i] {
			if index == arr[i][j] {
				return i, j
			}
		}
	}

	return 0, 0
}

// just for testing
func (s *Sudoku) debug() {
	for i := 0; i < 36; i++ {
		s.printSudoku()
		fmt.Println()
	}
}

// Generate generates new 6x6 sudoku
func (s *Sudoku) Generate() {
	if s.init == false {
		s.Init()
	}
	for i := 0; i < 36; i++ {

	}
}

// printSudoku just for debug
func (s *Sudoku) printSudoku() {
	for i := range s.v {
		fmt.Print(i, ":", s.v[i], " ")
		if (i+1)%3 == 0 {
			fmt.Print("|")
		}
		if (i+1)%6 == 0 {
			fmt.Print("\n")
		}
		if (i+1)%12 == 0 {
			fmt.Print("\n")
		}
	}
}
