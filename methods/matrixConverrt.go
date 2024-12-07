package methods

import (
	"math"
	"tsp/data"
)

func MatrixConversion(mx [][]int) ([][]int, int) {
	rc := rowsConversion(mx)
	cc, sum := columnsConversion(rc)

	matrix := cutFromConversionMatrix(cc)
	//fmt.Println("FUNC MATRIXCONVERSION IS ENDED.")
	return matrix, sum
}

func cutFromConversionMatrix(mx [][]int) [][]int {
	rows := len(mx)
	cols := len(mx[0])
	for i := range mx {
		mx[i] = mx[i][:cols-1]
	}
	mx = mx[:rows-1]
	return mx
}

func rowsConversion(mx [][]int) [][]int {
	rows := len(mx)
	cols := len(mx[0])

	// создаем результирующую матрицу на одну колонку больше
	resultMx := make([][]int, rows)
	for i := 0; i < rows; i++ {
		resultMx[i] = make([]int, cols+1)
	}
	// заполняем заголовки
	for i := 0; i < rows; i++ {
		resultMx[i][0] = mx[i][0]
	}
	for j := 0; j < cols; j++ {
		resultMx[0][j] = mx[0][j]
	}
	// идем по строкам исключая строку с заголовками и ищем минимумы в каждой строке
	for i := 1; i < rows; i++ {
		min := math.MaxInt
		// идем по ячейкам строки исключая заголовок строки
		for j := 1; j < cols; j++ {
			// находим минимум в строке
			if mx[i][j] < data.Inf {
				if mx[i][j] < min {
					min = mx[i][j]
				}
			}
		}
		// вычитаем найденный минимум из каждого элемента строки, исключая заголовок
		for j := 1; j < cols; j++ {
			if mx[i][j] < data.Inf {
				resultMx[i][j] = mx[i][j] - min
			} else {
				resultMx[i][j] = mx[i][j]
			}
		}
		// записываем результат в конец строки
		resultMx[i][cols] = min
	}
	//PrintMatrix(resultMx)
	return resultMx
}

func columnsConversion(mx [][]int) ([][]int, int) {
	rows := len(mx)
	cols := len(mx[0])

	// создаем результирующую матрицу на одну строку больше
	resultMx := make([][]int, rows+1)
	for i := 0; i < rows+1; i++ {
		resultMx[i] = make([]int, cols)
	}
	// заполняем заголовки
	for i := 0; i < rows; i++ {
		resultMx[i][0] = mx[i][0]
	}
	for j := 0; j < cols; j++ {
		resultMx[0][j] = mx[0][j]
	}
	// идем по колонке, исключая колонку с заголовками
	for j := 1; j < cols-1; j++ {
		min := math.MaxInt
		// идем по ячейкам колонки исключая заголовок колонки
		for i := 1; i < rows; i++ {
			if mx[i][j] < data.Inf {
				if mx[i][j] < min {
					min = mx[i][j]
				}
			}
		}
		// if min == math.MaxInt {
		// 	min = 0
		// }
		// вычитаем найденный минимум из каждого элемента колонки, исключая заголовок
		for i := 1; i < rows; i++ {
			if mx[i][j] < data.Inf {
				resultMx[i][j] = mx[i][j] - min
			} else {
				resultMx[i][j] = mx[i][j]
			}
		}
		// записываем результат в конец колонки
		resultMx[rows][j] = min
	}
	// PrintMatrix(resultMx)
	// Считаем сумму коэффициентов строк и колонок
	var sum int
	for i := 1; i < rows; i++ {
		resultMx[i][cols-1] = mx[i][cols-1]
		sum += mx[i][cols-1]
	}
	for j := 1; j < cols-1; j++ {
		sum += resultMx[rows][j]
	}
	resultMx[rows][cols-1] = sum
	return resultMx, sum
}
