package main

import (
	"fmt"
	"tsp/bitree"
	"tsp/iteration"
	"tsp/methods"
	"tsp/models"

	"time"
	"tsp/data"
)

var Debug = false

func main() {
	for i := range data.Matrixes {
		if models.Debug {
			fmt.Printf("\n#########################\n#\tMatrix: %d\t#\n#########################\n", i)
		}
		out := 1
		setOut(data.Matrixes[i], out)
		t := time.Now()
		calculate(data.Matrixes[i], out)
		ts := time.Since(t)
		fmt.Printf("Time latency: %v\n", ts)
	}
}

func setOut(mx [][]int, out int) {
	// устанавливаем город отправления
	if out != 0 {
		for i := range mx {
			if mx[i][out-1] != data.Inf {
				mx[i][out-1] = 0
			}
		}
	}
}

func calculate(mx [][]int, out int) {
	models.Debug = Debug

	// именуем столбцы и строки
	matrixNamed := methods.SetNaming(mx)

	models.MxRoot, models.LowWeightLimit = methods.MatrixConversion(matrixNamed)
	if models.Debug {
		fmt.Println("Исходная матрица:")
		methods.PrintMatrix(matrixNamed)
		fmt.Println("Приведённая исходная матрица:")
		methods.PrintMatrix(models.MxRoot)
		fmt.Printf("Нижняя весовая граница: %d\n", models.LowWeightLimit)
	}
	/* создаем корневой узел дерева с параметрами:
	Q           критерий кратчайшего пути
	State       мапа с узлами дерева и копиями матриц отложенных узлов
	Count       счетчик узлов дерева
	Result      структура с результатами одной итерации (Маршрут и отложенные узлы с весам, приведенные матрицы узлов)
	CurrentNode текущий узел дерева
	RootNode    корневой узел дерева */
	bitree.BT = bitree.NewBiTree(matrixNamed, models.LowWeightLimit)

	toursArray := iteration.IterationBranch()
	methods.PrintResult(toursArray, matrixNamed, out)
	//bitree.PrintTree(bitree.BT.RootNode)
}

func printAllNodes() {
	fmt.Println("Все узлы Маршрута:")
	for _, v := range bitree.BT.Result.Tour {
		fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
		// methods.PrintMatrix(v.Mxs)
		methods.PrintMatrix(models.MxRoot)
	}
	fmt.Println("Все отложенные узлы:")
	for _, v := range bitree.BT.Result.Back {
		fmt.Printf("W:%d, %s(%d,%d), id: %d\n", v.W, v.Sign, v.Out, v.In, v.ID)
		// methods.PrintMatrix(v.Mxs)
		methods.PrintMatrix(models.MxRoot)
	}
}
