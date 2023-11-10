package main

import (
	"fmt"

	rm "rotate_matrix/rotate_matrix"
)

func main() {
	fmt.Println("ready")

	inputMatrix := [][]int64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(rm.RotateMatrixRight90(inputMatrix))
}
