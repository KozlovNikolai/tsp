package models

import (
	"sync"
)

type CellWithMaxMin struct {
	RowName int
	ColName int
	MaxSum  int
}

var LowWeightLimit int
var LowerWeightFound int
var Debug bool
var MxRoot [][]int

type NamesOfIndexes struct {
	mu          sync.Mutex
	NamesOfRows []int
	NamesOfCols []int
}

func (nof *NamesOfIndexes) GetRowIdx(name int) int {
	nof.mu.Lock()
	defer nof.mu.Unlock()
	for i, v := range nof.NamesOfRows {
		if v == name {
			return i
		}
	}
	return -1
}

func (nof *NamesOfIndexes) GetRowName(index int) int {
	nof.mu.Lock()
	defer nof.mu.Unlock()
	return nof.NamesOfRows[index]
}

func (nof *NamesOfIndexes) GetColIdx(name int) int {
	nof.mu.Lock()
	defer nof.mu.Unlock()
	for i, v := range nof.NamesOfCols {
		if v == name {
			return i
		}
	}
	return -1
}

func (nof *NamesOfIndexes) GetColName(index int) int {
	nof.mu.Lock()
	defer nof.mu.Unlock()
	return nof.NamesOfCols[index]
}

func (nof *NamesOfIndexes) GetNames() NamesOfIndexes {
	nof.mu.Lock()
	defer nof.mu.Unlock()
	return *nof
}

func (nof *NamesOfIndexes) RemoveRowByIndex(index int) (int, *NamesOfIndexes) {
	nof.mu.Lock()
	defer nof.mu.Unlock()
	name := nof.NamesOfRows[index]
	lenRow := len(nof.NamesOfRows)
	for i := index; i < lenRow-1; i++ {
		nof.NamesOfRows[i] = nof.NamesOfRows[i+1]
	}
	// fmt.Printf("%+v\n", nof.NamesOfRows)
	nof.NamesOfRows = nof.NamesOfRows[:lenRow-1]
	// fmt.Printf("%+v\n", nof.NamesOfRows)

	return name, nof
}

func (nof *NamesOfIndexes) RemoveColByIndex(index int) (int, *NamesOfIndexes) {
	nof.mu.Lock()
	defer nof.mu.Unlock()
	name := nof.NamesOfCols[index]
	lenCol := len(nof.NamesOfCols)
	for i := index; i < lenCol-1; i++ {
		nof.NamesOfCols[i] = nof.NamesOfCols[i+1]
	}
	// fmt.Printf("%+v\n", nof.NamesOfCols)
	nof.NamesOfCols = nof.NamesOfCols[:lenCol-1]
	// fmt.Printf("%+v\n", nof.NamesOfCols)

	return name, nof
}
