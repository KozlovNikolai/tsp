package methods

import "testing"

func TestPrintMatrix(t *testing.T) {
	type args struct {
		mx [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				mx: [][]int{
					{0, 1, 2, 3, 4, 5},
					{1, -1, 1, 2, 3, 4},
					{2, 14, -1, 15, 16, 5},
					{3, 13, 20, -1, 17, 6},
					{4, 12, 19, 18, -1, 7},
					{5, 11, 10, 9, 8, -1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintMatrix(tt.args.mx)
		})
	}
}
