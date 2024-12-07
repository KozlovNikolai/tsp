package methods

func SetNaming(mx [][]int) [][]int {
	lenRows := len(mx)
	lenCols := len(mx[0])

	mmx := make([][]int, lenRows+1)
	for i := range mmx {
		mmx[i] = make([]int, lenCols+1)
		// заполняем заголовки столбцов:
		if i == 0 {
			for j := range mmx[i] {
				mmx[i][j] = j
			}
		} else {
			mmx[i][0] = i
			for j := 1; j < len(mmx[i]); j++ {
				mmx[i][j] = mx[i-1][j-1]
			}
		}

	}
	return mmx
}
