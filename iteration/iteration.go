package iteration

import (
	"fmt"
	"tsp/bitree"
	"tsp/data"
	"tsp/methods"
	"tsp/models"
)

// var cnt = 0
var prevFoundWeight int = 0

func IterationBranch(start int) []bitree.Node {
	var toursArray []bitree.Node
	// prevFoundWeight := data.Inf
	// prevFoundWeight := 0
	// устанавливаем начальную лучшую стоимость по случайному маршруту:
	t := 1
	for {
		prevFoundWeight += models.MxOriginal[t][t+1]
		t++
		if t == len(models.MxOriginal)-1 {
			prevFoundWeight += models.MxOriginal[t][1]
			break
		}
	}
	fmt.Printf("Начальная стоимость: %d\n", prevFoundWeight)
	weight := 0
	fmt.Println("Строим ветвь............................................................................................")
	// начинаем итерации создания ветвей:
	//matrix := bitree.CloneMx(models.MxRoot)
	//IterationNode(matrix)
	//toursArray = append(toursArray, bitree.BT.Result.Tour...)
	for {
		// 	// начинаем итерации создания узлов:
		// 	if models.Debug {
		// 		// fmt.Println("Матрица на входе в итератор:")
		// 		// methods.PrintMatrix(models.MxRoot)
		// 	}

		matrix := bitree.CloneMx(models.MxRoot)

		tourIsValid := IterationNode(matrix, start)

		if tourIsValid {
			if bitree.BT.CurWeight < prevFoundWeight {
				prevFoundWeight = bitree.BT.CurWeight
				toursArray = nil
				toursArray = append(toursArray, bitree.BT.Result.Tour...)
			} else {
				break
			}
			var row, col, id int
			bitree.BT.CurrentNode, weight, row, col, id = findInBack()

			if bitree.BT.CurrentNode == nil {
				fmt.Printf("!!! current node is NIL !!!\n")
				break
			}
			if weight == 0 {
				fmt.Printf("!!! weight is Null !!!\n")
				break
			}
			if row == 0 || col == 0 {
				fmt.Printf("!!! Row or Col is Null !!!\n")
				break
			}
			bitree.BT.CurrentID = id
			// models.MxRoot = bitree.CloneMx(bitree.BT.AllNodes[id].Mxs)

			mtr := bitree.CloneMx(bitree.BT.AllNodes[0].Mxs)
			mtr[row][col] = data.Inf
			models.MxRoot, models.LowWeightLimit = methods.MatrixConversion(mtr)
			if bitree.BT.AllNodes[id].Sign == "-" {

				//models.LowWeightLimit = bitree.BT.AllNodes[id].W
			} else {
				//models.LowWeightLimit = bitree.BT.AllNodes[id].W
				bitree.BT.Result.Tour = nil
				bitree.BT.Result.Tour = append(bitree.BT.Result.Tour, bitree.BT.AllNodes[id])
			}
		} else {
			break
		}
	}
	bitree.PrintTree(bitree.BT.RootNode)
	return toursArray
}

var step int

func IterationNode(matrix [][]int, start int) bool {
	// создаем узлы ветви:

	for {
		step++
		if models.Debug {
			fmt.Println("________________________ Начало создания узла _______________________")
			fmt.Printf("Шаг: %d\n", step)
		}
		mx, isRight := Step(matrix)
		if isRight {
			if len(mx) == 3 {
				// fmt.Printf("\nBreak, размер матрицы достиг: [%dx%d]\n", len(mx), len(mx[0]))
				if models.Debug {
					fmt.Printf("\nBreak, размер матрицы достиг: [%dx%d]\n", len(mx)-1, len(mx[0])-1)
				}
				EndingBranch(mx)
				// сохраняем найденный лучший вес и выходим
				ok := checkTour(bitree.BT.Result.Tour, start)
				if ok {
					fmt.Println("Chektour - OK !!!")
					bitree.BT.CurWeight = models.LowWeightLimit
				} else {
					fmt.Println("Chektour - FAILURE !!!")
				}
				for _, v := range bitree.BT.Result.Tour {
					fmt.Printf("id:%d, %s(%d,%d), W:%d\n", v.ID, v.Sign, v.Out, v.In, v.W)
				}
				return true
			}

		}
		matrix = bitree.CloneMx(mx)
	}
}

func Step(mc [][]int) ([][]int, bool) {
	//tempWeight := models.LowWeightLimit
	var isRight bool
	//saveMx := bitree.CloneMx(mc)
	if models.Debug {
		fmt.Println("Матрица на вход в STEP:")
		methods.PrintMatrix(mc)
	}
	// ищем ячейку по максимальной сумме минимумов строк и столбцов нулевых ячеек:
	nextNode := methods.FindCellWithMaxMin(mc)
	if models.Debug {
		fmt.Printf("(%d,%d) MaxSum:%d\n", nextNode.RowName, nextNode.ColName, nextNode.MaxSum)
	}

	// получаем матрицу для левого узла
	leftMx := bitree.CloneMx(mc)
	// исключаем ребро
	leftMxRow, leftMxCol, ok := methods.IdxByName(leftMx, nextNode.RowName, nextNode.ColName)
	if !ok {
		fmt.Printf("Не могу получить индексы для левой матрицы на шаге: %d\n", step)
	}
	leftMx[leftMxRow][leftMxCol] = data.Inf
	// приводим матрицу и получаем вес
	conversionLeftMatrix, leftMatrixLowWeightLimit := methods.MatrixConversion(leftMx)
	if models.Debug {
		fmt.Printf("исключаем ребро (%d,%d) и получаем нижнюю границу целевой функции (НГЦФ):\n", nextNode.RowName, nextNode.ColName)
		methods.PrintMatrix(conversionLeftMatrix)
		fmt.Printf("нижняя весовая граница левой матрицы: %d + %d = %d\n", models.LowWeightLimit, leftMatrixLowWeightLimit, models.LowWeightLimit+leftMatrixLowWeightLimit)
	}
	leftMatrixLowWeightLimit += models.LowWeightLimit
	// if leftMatrixLowWeightLimit < 0 {
	// 	leftMatrixLowWeightLimit = data.Inf
	// }
	// получаем матрицу для правого узла
	rightMx := bitree.CloneMx(mc)
	// удаляем найденную ячейку с ее строкой и столбцом:
	reductionRightMatrix := methods.RemoveCellFromMatrixByName(rightMx, nextNode.RowName, nextNode.ColName)
	// маркируем зеркальную ячейку бесконечностью, (если она есть)
	rightMxRow, rightMxCol, ok := methods.IdxByName(reductionRightMatrix, nextNode.ColName, nextNode.RowName)
	if !ok {
		fmt.Printf("Не могу получить индексы для правой матрицы на шаге: %d\n", step)
	} else {
		reductionRightMatrix[rightMxRow][rightMxCol] = data.Inf
	}
	// получаем приведённую правую матрицу и нижнюю границу целевой функции (НГЦФ)
	conversionRightMatrix, rightMatrixLowWeightLimit := methods.MatrixConversion(reductionRightMatrix)
	if models.Debug {
		fmt.Printf("\nвключаем ребро, удаляя строку: %d и колонку: %d\n", nextNode.RowName, nextNode.ColName)
		fmt.Printf("приводим правую матрицу и получаем нижнюю границу целевой функции (НГЦФ):\n")
		methods.PrintMatrix(conversionRightMatrix)
		fmt.Printf("нижняя весовая граница правой матрицы: %d + %d = %d\n", models.LowWeightLimit, rightMatrixLowWeightLimit, models.LowWeightLimit+rightMatrixLowWeightLimit)
	}

	rightMatrixLowWeightLimit += models.LowWeightLimit
	// if rightMatrixLowWeightLimit < 0 {
	// 	rightMatrixLowWeightLimit = data.Inf
	// }
	// помечаем ячейки для предотвращения внутренних циклов

	markInfinityCells(conversionRightMatrix, nextNode.RowName, nextNode.ColName)
	if models.Debug {
		fmt.Printf("Помечаем ячейки для предотвращения внутренних циклов:\n")
		methods.PrintMatrix(conversionRightMatrix)
	}

	// определяем два новых узла и выбираем из них следующий
	var setCurrentRightNode bool
	if leftMatrixLowWeightLimit >= rightMatrixLowWeightLimit {
		setCurrentRightNode = true
	}

	parentID := bitree.BT.CurrentID
	bitree.BT.CreateLeftNode(parentID, conversionLeftMatrix, leftMatrixLowWeightLimit, nextNode.RowName, nextNode.ColName, !setCurrentRightNode)
	bitree.BT.CreateRightNode(parentID, conversionRightMatrix, rightMatrixLowWeightLimit, nextNode.RowName, nextNode.ColName, setCurrentRightNode)
	if models.Debug {
		for _, v := range bitree.BT.Result.Tour {
			fmt.Printf("id:%d, %s(%d,%d), W:%d\n", v.ID, v.Sign, v.Out, v.In, v.W)
		}
	}
	var mx [][]int
	if rightMatrixLowWeightLimit > prevFoundWeight {
		if len(bitree.BT.Result.Tour) > 0 {
			bitree.BT.Result.Tour = bitree.BT.Result.Tour[:len(bitree.BT.Result.Tour)-1]
		}
		grandParentID := bitree.BT.AllNodes[parentID].ParentID
		bitree.BT.CurrentNode = bitree.BT.AllNodes[grandParentID].Node.Left
		id := bitree.BT.AllNodes[grandParentID].LeftID
		models.LowWeightLimit = bitree.BT.AllNodes[grandParentID].W
		mx = bitree.CloneMx(bitree.BT.AllNodes[id].Mxs)
		bitree.BT.CurrentID = id
		if models.Debug {
			fmt.Println("Выбран верхний альтернативный узел.")
			fmt.Printf("id:%d,%s(%d,%d),W:%d\n",
				bitree.BT.AllNodes[id].ID,
				bitree.BT.AllNodes[id].Sign,
				bitree.BT.AllNodes[id].Out,
				bitree.BT.AllNodes[id].In,
				bitree.BT.AllNodes[id].W)
			methods.PrintMatrix(mx)
		}
		if models.Debug {
			for _, v := range bitree.BT.Result.Tour {
				fmt.Printf("id:%d, %s(%d,%d), W:%d\n", v.ID, v.Sign, v.Out, v.In, v.W)
			}
		}
		return mx, isRight
	}

	if setCurrentRightNode {
		//	fmt.Println("Выбран правый узел.")
		if models.Debug {
			fmt.Println("Выбран правый узел.")
		}
		isRight = true
		bitree.BT.CurrentNode = bitree.BT.CurrentNode.Right
		models.LowWeightLimit = rightMatrixLowWeightLimit
		mx = bitree.CloneMx(conversionRightMatrix)
	} else {
		//fmt.Println("Выбран левый узел.")
		if models.Debug {
			fmt.Println("Выбран левый узел.")
			// fmt.Println("Пропускаем правый узел.")
		}
		// bitree.BT.CurrentNode = bitree.BT.CurrentNode.Right
		bitree.BT.CurrentNode = bitree.BT.CurrentNode.Left
		models.LowWeightLimit = leftMatrixLowWeightLimit
		// models.LowWeightLimit = tempWeight
		mx = bitree.CloneMx(conversionLeftMatrix)
	}

	return mx, isRight
}

func markInfinityCells(mx [][]int, rowName, colName int) {
	infCellArr := infArr(rowName, colName)

	for i := 1; i < len(mx); i++ {
		for j := 1; j < len(mx[0]); j++ {
			if _, ok := infCellArr[struct {
				i int
				j int
			}{
				i: mx[i][0],
				j: mx[0][j],
			}]; ok {
				mx[i][j] = data.Inf
			}
		}

	}

	// // идем по колонкам
	// for j := 1; j < len(mx[0]); j++ {
	// 	name := mx[0][j] // берем имя колонки из матрицы
	// 	count := 0       // счетчик итераций мапе
	// 	for {
	// 		count++
	// 		value, ok := list[name] // ищем по имени колонки матрицы (как имя строки в списке) колонку из записей в построенного пути
	// 		fmt.Printf("rowName:%d, colName:%d\n", rowName, colName)
	// 		for key, val := range list {
	// 			fmt.Printf("key:%d, value:%d\n", key, val)
	// 		}
	// 		fmt.Println("__________")
	// 		if !ok { // если такой пары в построенном пути нет - прерываемся и берем следующую колонку из матрицы
	// 			break
	// 		}
	// 		// если такую пару в построенном пути нашли, то идем по именам строк матрицы и ищем
	// 		// совпадение имени строки матрицы  с именем колонки из найденой пары
	// 		for i := 1; i < len(mx); i++ {
	// 			if mx[i][0] == value { // если совпадение нашли - маркируем ячейку матрицы
	// 				if count == (len(models.MxRoot) - 2) { //если
	// 					//list[mx[1][0]] = mx[0][1]
	// 					log.Println("---  Point ----")
	// 					//return
	// 					break
	// 				}
	// 				mx[i][j] = data.Inf
	// 			}
	// 		}
	// 		n := list[value]
	// 		if name != n {
	// 			name = value
	// 		}
	// 	}
	// }
}

// func markInfinityCells(mx [][]int, rowName, colName int) {
// 	list := map[int]int{
// 		rowName: colName,
// 	}

// 	for _, node := range bitree.BT.Result.Tour {
// 		list[node.Out] = node.In
// 	}

// 	for j := 1; j < len(mx[0]); j++ {
// 		name := mx[0][j]
// 		count := 0
// 		for {
// 			count++
// 			value, ok := list[name]
// 			if !ok {
// 				break
// 			}
// 			for i := 1; i < len(mx); i++ {
// 				if mx[i][0] == value {
// 					if count == (len(models.MxRoot) - 2) {
// 						list[mx[1][0]] = mx[0][1]
// 						log.Println("---  Point ----")
// 						return
// 					}
// 					mx[i][j] = data.Inf
// 				}
// 			}
// 			name = value
// 		}
// 	}
// }

//	func infArr(data map[int]int, rowName, colName int) (map[struct {
//		i int
//		j int
//	}]struct{}, map[int]int) {
func infArr(rowName, colName int) map[struct {
	i int
	j int
}]struct{} {
	// создаем мапу для пар из уже добавленных узлов тура и добавляем туда текущий создаваемый узел
	tour := make(map[int]int)

	// добавляем туда остальные существующие в результирующем туре узлы
	for _, node := range bitree.BT.Result.Tour {
		tour[node.Out] = node.In
	}

	tour[rowName] = colName

	// for key, value := range data {
	// 	tour[key] = value
	// }
	// создаем список бесконечных ячеек
	list := make(map[struct {
		i int
		j int
	}]struct{})

	for row, col := range tour {
		list[struct {
			i int
			j int
		}{
			i: col,
			j: row,
		}] = struct{}{}

		for {
			if val, ok := tour[col]; ok {
				// fmt.Printf("input row:%d, col:%d\n", rowName, colName)
				// fmt.Printf("***************************************Found key:%d, val:%d\n", col, val)
				// for k, v := range tour {
				// 	fmt.Printf("Tour key:%d, val:%d\n", k, v)
				// }
				list[struct {
					i int
					j int
				}{i: val, j: row}] = struct{}{}
				col = val
				continue
			}
			break
		}
	}
	// return list, tour
	return list
}

func EndingBranch(mx [][]int) {
	if models.Debug {
		fmt.Println("Ending branch matrix:")
		methods.PrintMatrix(mx)
	}
	for i := 1; i < len(mx); i++ {
		for j := 1; j < len(mx[0]); j++ {
			if mx[i][j] == 0 {
				bitree.BT.CreateRightNode(bitree.BT.CurrentID, mx, models.LowWeightLimit, mx[i][0], mx[0][j], true)
				bitree.BT.CurrentNode = bitree.BT.CurrentNode.Right
			}
		}
	}
}

// rowIdx, colIdx, ok := methods.IdxByName(models.MxRoot, mx[1][0], mx[0][1])
// if !ok {
// 	log.Println("Ending branch: не могу получить индексы из имени !!!")
// }

func findInBack() (*bitree.TreeNode, int, int, int, int) {
	if models.Debug {

		fmt.Printf("Поиск в отложенных узлах: %d штук\n", len(bitree.BT.Result.Back))
	}
	minWeight := data.Inf
	var n int
	for i := 1; i < len(bitree.BT.Result.Back); i++ {

		if bitree.BT.Result.Back[i].W < minWeight {
			minWeight = bitree.BT.Result.Back[i].W
			n = i
		}
	}

	if bitree.BT.CurWeight > bitree.BT.Result.Back[n].W {
		if models.Debug {
			fmt.Printf("Найдено в отложенных:  W:%d, %s(%d,%d), id: %d\n",
				bitree.BT.Result.Back[n].W,
				bitree.BT.Result.Back[n].Sign,
				bitree.BT.Result.Back[n].Out,
				bitree.BT.Result.Back[n].In,
				bitree.BT.Result.Back[n].ID)
		}

		node := bitree.BT.Result.Back[n].Node
		w := bitree.BT.Result.Back[n].W
		row := bitree.BT.Result.Back[n].Out
		col := bitree.BT.Result.Back[n].In
		id := bitree.BT.Result.Back[n].ID
		bitree.BT.Result.Back[n] = bitree.BT.Result.Back[len(bitree.BT.Result.Back)-1]
		bitree.BT.Result.Back = bitree.BT.Result.Back[:len(bitree.BT.Result.Back)-1]
		bitree.BT.Result.Tour = bitree.BT.Result.Tour[:0]
		return node, w, row, col, id
	}

	return nil, 0, 0, 0, 0
}

func checkTour(tour []bitree.Node, out int) bool {
	rt := make(map[int]int)
	for _, v := range tour {
		rt[v.Out] = v.In
	}
	if out == 0 {
		out = 1
	}
	count := len(tour)
	for count > 0 {
		in, ok := rt[out]
		if !ok {
			return false
		}
		out = in
		count--
	}
	return true
}
