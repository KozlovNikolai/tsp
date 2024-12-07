package methods

import (
	"fmt"
	"testing"
)

func TestFindInfinityCellCoords(t *testing.T) {
	type args struct {
		mx     [][]int
		rowDel int
		colDel int
	}
	tests := []struct {
		name           string
		args           args
		wantRowInfName int
		wantColInfName int
	}{
		{
			name: "",
			args: args{
				mx: [][]int{
					{0, 1, 2, 3, 4, 5},
					{1, -1, 0, 0, 2, 3},
					{2, 0, -1, 3, 5, -1},
					{3, 4, 14, -1, 11, 0},
					{4, 2, 12, 10, -1, 0},
					{5, 0, 2, 0, 0, -1},
				},
				rowDel: 3,
				colDel: 5,
			},
			wantRowInfName: 5,
			wantColInfName: 3,
		},
		{
			name: "",
			args: args{
				mx: [][]int{
					{0, 1, 2, 3, 4},
					{1, -1, 0, 0, 2},
					{2, 0, -1, 3, 5},
					{4, 2, 12, 10, -1},
					{5, 0, 2, 0, 0},
				},
				rowDel: 3,
				colDel: 5,
			},
			wantRowInfName: 5,
			wantColInfName: 3,
		},
		{
			name: "",
			args: args{
				mx: [][]int{
					{0, 1, 2, 3, 4, 5},
					{1, -1, 9, 0, 2, 3},
					{2, 0, -1, 3, 5, -1},
					{3, 4, 14, 3, 11, 0},
					{4, 2, 12, 10, -1, 0},
					{5, 0, 2, 0, 0, -1},
				},
				rowDel: 3,
				colDel: 5,
			},
			wantRowInfName: 5,
			wantColInfName: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRowInfName, gotColInfName := FindInfinityCellCoords(tt.args.mx)
			fmt.Printf("Row: %d, Col: %d\n", gotRowInfName, gotColInfName)
			// if gotRowInfName != tt.wantRowInfName {
			// 	t.Errorf("FindInfinityCellCoords() gotRowInfName = %v, want %v", gotRowInfName, tt.wantRowInfName)
			// }
			// if gotColInfName != tt.wantColInfName {
			// 	t.Errorf("FindInfinityCellCoords() gotColInfName = %v, want %v", gotColInfName, tt.wantColInfName)
			// }
		})
	}
}
