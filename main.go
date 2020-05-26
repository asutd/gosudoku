package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

type Value struct {
	i            int
	v            int
	cell         *Cell
	indexInCell  int
	vline        *Vline
	indexInVline int
	hline        *Hline
	indexInHline int
	s            *Sudoku
}

func (v *Value) set(value int) {
	v.v = value
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
	su.Generate()
	//var t [][]int
	//Perm([]int{2,3,5}, func(a []int) {
	//	var temp []int
	//	temp = append(temp, a...)
	//	t = append(t, temp)
	//})
	//fmt.Println(t)
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
	availableValues []int
}

func NewSudoku() *Sudoku {
	su := Sudoku{}
	su.Init()
	return &su
}

func (s *Sudoku) Init() {
	s.cellsIndexInfo = [6][6]int{
		{0, 1, 2, 6, 7, 8},
		{3, 4, 5, 9, 10, 11},
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

	s.availableValues = []int{1, 2, 3, 4, 5, 6}

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
			if i == 11 {
				_ = &struct {
				}{}
			}
			cellIndex, indexInCell := s.GetCellIndexes(i)
			v.cell = &s.cells[cellIndex]
			v.indexInCell = indexInCell
			v.cell.i = cellIndex
			v.cell.unknownCount = 6
			v.cell.values[v.indexInCell] = &v
		}

		{
			vlineIndex, indexInVline := s.GetVlineIndexes(i)
			v.vline = &s.vlines[vlineIndex]
			v.indexInVline = indexInVline
			v.vline.i = vlineIndex
			v.vline.unknownCount = 6
			v.vline.values[indexInVline] = &v
		}

		{
			hlineIndex, indexInHline := s.GetHlineIndexes(i)
			v.hline = &s.hlines[hlineIndex]
			v.indexInHline = indexInHline
			v.hline.i = hlineIndex
			v.hline.unknownCount = 6
			v.hline.values[v.indexInHline] = &v
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
	for i := range s.cells {
		cell := s.cells[i]
		if i == 0 {
			firstValues := s.shuffleFirst()
			for g, v := range firstValues {
				s.SetValue(cell.values[g].i, v)
			}
			continue
		}
		if (i+1)%2 == 0 {
			upperThreeValues := make([]int, 0, 3)
			valuesForUpper := s.cells[i-1].values[3:]
			for iUpper := range valuesForUpper {
				upperThreeValues = append(upperThreeValues, valuesForUpper[iUpper].v)
			}
			bottomThreeValues := make([]int, 0, 3)
			valuesForBottom := s.cells[i-1].values[:3]
			for iBottom := range valuesForBottom {
				bottomThreeValues = append(bottomThreeValues, valuesForBottom[iBottom].v)
			}
			var upperPermutation [][]int
			Perm(upperThreeValues, func(a []int) {
				var temp []int
				temp = append(temp, a...)
				upperPermutation = append(upperPermutation, temp)
			})

			fmt.Println(upperPermutation)
			var upperRes [][]int

			for g := range upperPermutation {
				if s.checkCellValues(upperPermutation[g], &cell) {
					upperRes = append(upperRes, upperPermutation[g])
				}
			}

			pivot := rand.Intn(len(upperRes) - 1)

			for g, v := range upperRes[pivot] {
				s.SetValue(cell.values[g].i, v)
			}
			continue
		}
		for j := range cell.values {
			v := s.generateValue(cell.values[j].i)
			s.SetValue(cell.values[j].i, v)
		}
	}
}

func (s *Sudoku) checkCellValues(cellValues []int, cell *Cell) bool {
	shift := 0
	if cell.values[0].v != 0 {
		shift = 2
	}
	for k := range cellValues {
		if !s.checkByIndexValue(cell.values[k+shift].i, cellValues[k]) {
			return false
		}
	}

	return true
}

func (s *Sudoku) shuffleFirst() []int {
	a := make([]int, 0, 6)
	a = append(a, s.availableValues...)

	return s.shuffleSlice(a)
}

func (s *Sudoku) shuffleSlice(a []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func (s *Sudoku) generateValue(i int) int {
	values, err := s.getPossibleValues(i)
	if err != nil {
		panic("not determined")
	}
	lenValues := len(values)
	if lenValues == 1 {
		return values[0]
	}
	pivot := rand.Intn(len(values) - 1)

	return values[pivot]
}

func (s *Sudoku) getPossibleValues(i int) ([]int, error) {
	p := [6]*struct{}{}
	values := make([]int, 0, 6)
	value := s.v[i]
	if i == 19 {
		f := &p
		_ = f
	}
	for j := range s.v[i].cell.containValues {
		if value.cell.containValues[j] != nil {
			p[j] = &struct {
			}{}
		}
		if value.hline.containValues[j] != nil {
			p[j] = &struct {
			}{}
		}
		if value.vline.containValues[j] != nil {
			p[j] = &struct {
			}{}
		}
	}

	for z := range p {
		if p[z] == nil {
			values = append(values, z+1)
		}
	}
	len := len(values)
	if len == 0 {
		return nil, errors.New("there is no possible values")
	}

	return values, nil
}

func (s *Sudoku) SetValue(i int, v int) bool {
	if !s.checkByIndexValue(i, v) {
		return false
	}
	s.v[i].hline.containValues[v-1] = &struct{}{}
	s.v[i].vline.containValues[v-1] = &struct{}{}
	s.v[i].cell.containValues[v-1] = &struct{}{}
	s.v[i].set(v)

	return true
}

func (s *Sudoku) checkByIndexValue(i int, v int) bool {
	if s.v[i].hline.containValues[v-1] != nil || s.v[i].vline.containValues[v-1] != nil || s.v[i].cell.containValues[v-1] != nil {
		return false
	}

	return true
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

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
