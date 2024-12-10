package methods

import (
	"fmt"
	"log"
	"tsp/data"
	"tsp/models"
)

func RemoveCellFromMatrixByName(mx [][]int, nameRow int, nameCol int) [][]int {
	idxRow, idxCol, ok := IdxByName(mx, nameRow, nameCol)
	if !ok {
		log.Println("Первый: не могу получить индексы из имени !!!")
	}
	mt := RemoveRowFromMatrixByIndex(mx, idxRow)
	resultMx := RemoveColFromMatrixByIndex(mt, idxCol)
	return resultMx
}

func RemoveRowFromMatrixByIndex(mx [][]int, nameIndex int) [][]int {
	lenRows := len(mx)
	var resultMx [][]int
	for i := 0; i < lenRows-1; i++ {
		if i < nameIndex {
			resultMx = append(resultMx, mx[i])
		} else {
			resultMx = append(resultMx, mx[i+1])
		}
	}
	return resultMx
}

func RemoveColFromMatrixByIndex(mx [][]int, nameIndex int) [][]int {
	lenRows := len(mx)
	var resultMx [][]int
	for i := 0; i < lenRows; i++ {
		resultMx = append(resultMx, mx[i][:nameIndex])
		resultMx[i] = append(resultMx[i], mx[i][nameIndex+1:]...)
	}
	return resultMx
}

// func FindInfinityCellCoords(mx [][]int, rowDel, colDel int) (rowInfName, colInfName int) {
// 	for i := 0; i < len(mx); i++ {
// 		if mx[i][colDel] == -1 {
// 			rowInfName = i
// 			break
// 		}
// 	}
// 	for j := 0; j < len(mx[0]); j++ {
// 		if mx[rowDel][j] == -1 {
// 			colInfName = j
// 		}
// 	}
// 	return
// }

func FindInfinityCellCoords(mx [][]int) (rowInfName, colInfName int) {
	if models.Debug {
		PrintMatrix(mx)
		fmt.Println("удаляем строку и столбец   ^^^")
		fmt.Println("___________________________________________________")
	}

	for i := 1; i < len(mx); i++ {
		for j := 1; j < len(mx[0]); j++ {
			// fmt.Printf("mx[%d,%d]=%d\n", i, j, mx[i][j])
			if mx[i][j] == data.Inf {
				break
			}
			if j == len(mx[0])-1 {
				rowInfName = i
			}
		}
	}
	for j := 0; j < len(mx[0]); j++ {
		for i := 1; i < len(mx); i++ {
			if mx[i][j] == data.Inf {
				break
			}
			if i == len(mx)-1 {
				colInfName = j
			}
		}
	}
	return
}

func FindInfinityCellCoordsNew(mx [][]int) (rowInfName, colInfName int) {
	if models.Debug {
		PrintMatrix(mx)
		fmt.Println("удаляем строку и столбец   ^^^")
		fmt.Println("___________________________________________________")
	}

	for i := 1; i < len(mx); i++ {
		for j := 1; j < len(mx[0]); j++ {
			// fmt.Printf("mx[%d,%d]=%d\n", i, j, mx[i][j])
			if mx[i][j] == data.Inf {
				break
			}
			if j == len(mx[0])-1 {
				rowInfName = i
			}
		}
	}
	for j := 0; j < len(mx[0]); j++ {
		for i := 1; i < len(mx); i++ {
			if mx[i][j] == data.Inf {
				break
			}
			if i == len(mx)-1 {
				colInfName = j
			}
		}
	}
	return
}

func IdxByName(m [][]int, rowName, colName int) (rowIdx, colIdx int, ok bool) {
	for i, v := range m {
		if v[0] == rowName {
			rowIdx = i
			break
		}
	}
	if rowIdx == 0 {
		return 0, 0, false
	}
	for j, v := range m[0] {
		if v == colName {
			colIdx = j
			break
		}
	}
	if colIdx == 0 {
		return 0, 0, false
	}
	return rowIdx, colIdx, true
}
