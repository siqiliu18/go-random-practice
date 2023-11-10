package rotate_matrix

func RotateMatrixRight90(matrix [][]int64) [][]int64 {
	rowLen := len(matrix)
	colLen := len(matrix[0])

	res := make([][]int64, colLen)

	for j := 0; j < colLen; j++ {
		/*

			res[j] = make([]int64, rowLen)

			Here clearly shows the diff between new and make in Golang.
				- new will create a pointer for that datastruct, that is it
				- make will create a pointer and also init (provide default values within) the data structure

			if we do res[j] = make([]int64, rowLen) here, it will create an arry of int64 with rowLen of 0 values in it
		*/
		for i := rowLen - 1; i >= 0; i-- {
			res[j] = append(res[j], matrix[i][j])
		}
	}

	return res
}

func RotateMatrixLeft90(matrix [][]int64) [][]int64 {
	rowLen := len(matrix)
	colLen := len(matrix[0])

	res := make([][]int64, colLen)

	for j := colLen - 1; j >= 0; j-- {
		for i := 0; i < rowLen; i++ {
			res[colLen-1-j] = append(res[colLen-1-j], matrix[i][j])
		}
	}

	return res
}
