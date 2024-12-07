package methods

import (
	"fmt"
	"tsp/bitree"
	"tsp/data"
	"tsp/models"

	"github.com/fatih/color"
)

func PrintMatrixColor(mx [][]int) {
	red := color.New(color.FgRed).SprintFunc()
	yel := color.New(color.FgHiYellow).SprintFunc()
	black := color.New(color.FgHiBlack).SprintFunc()
	green := color.New(color.FgHiGreen).SprintFunc()
	white := color.New(color.FgHiWhite).SprintFunc()
	blue := color.New(color.FgHiBlue).SprintFunc()
	// выводим заголовки столбцов
	// for _, v := range names.NamesOfCols {
	// 	fmt.Printf("\t%d", v)
	// }
	fmt.Println()
	for i := 0; i < len(mx); i++ {
		// выводим заголовки строк:
		// if i < len(names.NamesOfRows) {
		// 	fmt.Printf("%d", names.GetRowName(i))
		// }

		for j := 0; j < len(mx[i]); j++ {
			if i == 0 || j == 0 {
				fmt.Printf(blue("\t%d"), mx[i][j])
			} else if mx[i][j] == data.Inf {
				// fmt.Printf(black("%-4d"), mx[i][j])
				fmt.Printf(black("\t%d"), mx[i][j])
			} else {
				if i == len(mx)-1 && j == len(mx[i])-1 {
					// fmt.Printf(white("%-4d"), mx[i][j])
					fmt.Printf(white("\t%d"), mx[i][j])
				} else if i == len(mx)-1 || j == len(mx[i])-1 {
					//fmt.Printf(green("%-4d"), mx[i][j])
					fmt.Printf(green("\t%d"), mx[i][j])
				} else if mx[i][j] == 0 {
					// fmt.Printf(red("%-4d"), mx[i][j])
					fmt.Printf(red("\t%d"), mx[i][j])
				} else {
					// fmt.Printf(yel("%-4d"), mx[i][j])
					fmt.Printf(yel("\t%d"), mx[i][j])
				}
			}

			// fmt.Printf("%-4d", mx[i][j])
		}
		fmt.Println()
	}
}
func PrintMatrix(mx [][]int) {
	red := color.New(color.FgRed).SprintFunc()
	yel := color.New(color.FgHiYellow).SprintFunc()
	black := color.New(color.FgHiBlack).SprintFunc()
	green := color.New(color.FgHiGreen).SprintFunc()

	//fmt.Println()
	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx[i]); j++ {
			if i == 0 || j == 0 {
				fmt.Printf(green("\t%d"), mx[i][j])
			} else if mx[i][j] == data.Inf {
				fmt.Printf(black("\t%s"), "inf")
			} else {
				if mx[i][j] == 0 {
					fmt.Printf(red("\t%d"), mx[i][j])
				} else {
					fmt.Printf(yel("\t%d"), mx[i][j])
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintArrayOfNodes(arr []bitree.Node) {
	rt := make(map[int]int)
	for _, v := range arr {
		fmt.Printf("ID:%d, W:%d, %s(%d,%d)\n", v.ID, v.W, v.Sign, v.Out, v.In)
		rt[v.Out] = v.In
	}
}

func PrintResult(toursArray []bitree.Node, matrixNamed [][]int, out int) {
	rt := make(map[int]int)
	if models.Debug {

		fmt.Printf("\nResult tour with Q: %d\n", bitree.BT.CurWeight)
	}
	for _, v := range toursArray {
		if models.Debug {

			fmt.Printf("ID:%d, W:%d, %s(%d,%d)\n", v.ID, v.W, v.Sign, v.Out, v.In)
		}
		rt[v.Out] = v.In
	}
	temp := 1
	if out != 0 {
		temp = out
	}
	fmt.Printf("\nГород отправления: %d\n", temp)
	sum := 0
	for i := 0; i < len(rt); i++ {
		fmt.Printf("(%d,%d),Cost:%d\n", temp, rt[temp], matrixNamed[temp][rt[temp]])
		sum += matrixNamed[temp][rt[temp]]
		temp = rt[temp]
	}
	fmt.Printf("Sum: %d\n", sum)
}
