package iteration

import (
	"fmt"
	"log"
	"math"
	"tsp/bitree"
	"tsp/data"
	"tsp/methods"
	"tsp/models"
)

func IterationBranch() []bitree.Node {
	var toursArray []bitree.Node
	prevFoundWeight := math.MaxInt
	weight := 0
	if models.Debug {
		fmt.Println("Строим ветвь")
	}
	// начинаем итерации создания ветвей:
	for {
		// начинаем итерации создания узлов:
		if models.Debug {
			fmt.Println("Матрица на входе в итератор:")
			methods.PrintMatrix(models.MxRoot)
		}
		matrix := bitree.CloneMx(models.MxRoot)
		// ok := IterationNode(models.MxRoot, bitree.BT.RootNode)
		ok := IterationNode(matrix, bitree.BT.RootNode)
		if ok {
			if models.Debug {
				fmt.Printf("Current Weight: %d\n", bitree.BT.CurWeight)
				fmt.Printf("Previous found Weight: %d\n", prevFoundWeight)
			}
			if bitree.BT.CurWeight < prevFoundWeight {
				prevFoundWeight = bitree.BT.CurWeight
				toursArray = toursArray[:0]
				toursArray = append(toursArray, bitree.BT.Result.Tour...)
			}
			var row, col int
			// ищем в отложенных узлах узел с минимальным весом
			bitree.BT.CurrentNode, weight, row, col = findInBack()
			if models.Debug {
				fmt.Printf("findBack - current Node: %v, weight: %d, row: %d, col: %d\n", bitree.BT.CurrentNode, weight, row, col)
			}
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

			models.MxRoot[row][col] = data.Inf
			models.MxRoot, models.LowWeightLimit = methods.MatrixConversion(models.MxRoot)
			models.LowWeightLimit = weight
		} else {
			if models.Debug {
				fmt.Printf("NOT OK !!!\n")
			}
			break
		}

	}

	return toursArray
}

func IterationNode(matrix [][]int, node *bitree.TreeNode) bool {
	// создаем узлы ветви:
	for {
		if models.Debug {

			fmt.Println("________________________ Начало создания узла _______________________")
		}
		mx := Step(matrix)
		if bitree.BT.CurWeight < bitree.BT.Result.Tour[len(bitree.BT.Result.Tour)-1].W {
			if models.Debug {

				fmt.Printf("\nBreak, вес лучшего маршрута:%d - меньше веса создаваемого\n маршрута: %d, дальше идти нет смысла.\n", bitree.BT.CurWeight, bitree.BT.Result.Tour[len(bitree.BT.Result.Tour)-1].W)
			}
			return false
		}
		if len(mx) == 3 {
			if models.Debug {

				fmt.Printf("\nBreak, размер матрицы достиг: [%dx%d]\n", len(mx), len(mx[0]))
			}
			EndingBranch(mx)
			// сохраняем найденный лучший вес и выходим
			bitree.BT.CurWeight = models.LowWeightLimit
			return true
		}
		matrix = bitree.CloneMx(mx)
	}
}

func Step(mc [][]int) [][]int {
	if models.Debug {

		fmt.Println("Матрица на вход в STEP:")
		methods.PrintMatrix(mc)
	}
	// ищем ячейку по максимальной сумме минимумов строк и столбцов нулевых ячеек:
	nextNode := methods.FindCellWithMaxMin(mc)
	if models.Debug {
		fmt.Printf("(%d,%d) MaxSum:%d\n", nextNode.RowName, nextNode.ColName, nextNode.MaxSum)
	}

	// удаляем найденную ячейку с ее строкой и столбцом:
	reductionMatrix := methods.RemoveCellFromMatrixByIndex(mc, nextNode.RowName, nextNode.ColName)

	// помечаем ячейки для предотвращения внутренних циклов
	markInfinityCells(reductionMatrix, nextNode.RowName, nextNode.ColName)
	if models.Debug {
		fmt.Printf("Удаляем найденные строку и столбец и помечаем ячейки для предотвращения внутренних циклов:\n")
		methods.PrintMatrix(reductionMatrix)
	}

	// получаем приведённую матрицу и нижнюю границу целевой функции (НГЦФ)
	conversionMatrix, currentLowWeightLimit := methods.MatrixConversion(reductionMatrix)
	if models.Debug {
		fmt.Printf("получаем приведённую матрицу и нижнюю границу целевой функции (НГЦФ):\n")
		methods.PrintMatrix(conversionMatrix)
		fmt.Printf("текущая нижняя весовая граница: %d\n", currentLowWeightLimit)
	}

	// определяем два новых узла и выбираем из них следующий
	var setCurrentRightNode bool
	if models.LowWeightLimit+nextNode.MaxSum >= models.LowWeightLimit+currentLowWeightLimit {
		setCurrentRightNode = true
	}
	bitree.BT.CreateLeftNode(models.LowWeightLimit+nextNode.MaxSum, nextNode.RowName, nextNode.ColName, !setCurrentRightNode)
	bitree.BT.CreateRightNode(models.LowWeightLimit+currentLowWeightLimit, nextNode.RowName, nextNode.ColName, setCurrentRightNode)
	if setCurrentRightNode {
		bitree.BT.CurrentNode = bitree.BT.CurrentNode.Right
		models.LowWeightLimit = models.LowWeightLimit + currentLowWeightLimit
	} else {
		bitree.BT.CurrentNode = bitree.BT.CurrentNode.Left
		models.LowWeightLimit = models.LowWeightLimit + nextNode.MaxSum
	}
	return conversionMatrix
}

func markInfinityCells(mx [][]int, rowName, colName int) {
	list := map[int]int{
		rowName: colName,
	}

	for _, node := range bitree.BT.Result.Tour {
		list[node.Out] = node.In
	}

	for j := 1; j < len(mx[0]); j++ {
		name := mx[0][j]
		count := 0
		for {
			count++
			value, ok := list[name]
			if !ok {
				break
			}
			for i := 1; i < len(mx); i++ {
				if mx[i][0] == value {
					if count == (len(models.MxRoot) - 2) {
						list[mx[1][0]] = mx[0][1]
						log.Println("---  Point ----")
						return
					}
					mx[i][j] = data.Inf
				}
			}
			name = value
		}
	}
}

func EndingBranch(mx [][]int) {
	if models.Debug {
		fmt.Println("Ending branch matrix:")
		methods.PrintMatrix(mx)

	}
	for i := 1; i < len(mx); i++ {
		for j := 1; j < len(mx[0]); j++ {
			if mx[i][j] == 0 {
				bitree.BT.CreateRightNode(models.LowWeightLimit, mx[i][0], mx[0][j], true)
				bitree.BT.CurrentNode = bitree.BT.CurrentNode.Right
			}
		}
	}
}

// rowIdx, colIdx, ok := methods.IdxByName(models.MxRoot, mx[1][0], mx[0][1])
// if !ok {
// 	log.Println("Ending branch: не могу получить индексы из имени !!!")
// }

func findInBack() (*bitree.TreeNode, int, int, int) {
	if models.Debug {

		fmt.Printf("Поиск в отложенных узлах: %d штук\n", len(bitree.BT.Result.Back))
	}
	minWeight := math.MaxInt
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
		bitree.BT.Result.Back[n] = bitree.BT.Result.Back[len(bitree.BT.Result.Back)-1]
		bitree.BT.Result.Back = bitree.BT.Result.Back[:len(bitree.BT.Result.Back)-1]
		bitree.BT.Result.Tour = bitree.BT.Result.Tour[:0]
		return node, w, row, col
	}

	return nil, 0, 0, 0
}
