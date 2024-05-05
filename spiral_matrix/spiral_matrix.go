package spiralmatrix

type coord struct {
	x int
	y int
}

func spiralOrder(matrix [][]int) []int {
	xlen := len(matrix)
	ylen := len(matrix[0])

	visited := make([][]bool, len(matrix))

	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]bool, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			visited[i][j] = false
		}
	}

	direct := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	di := 0

	res := []int{}

	q := []coord{}
	q = append(q, coord{0, 0})

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		res = append(res, matrix[curr.x][curr.y])
		visited[curr.x][curr.y] = true

		if curr.x+direct[di][0] < 0 || curr.x+direct[di][0] >= xlen || curr.y+direct[di][1] < 0 || curr.y+direct[di][1] >= ylen || visited[curr.x+direct[di][0]][curr.y+direct[di][1]] {
			if di+1 == 4 {
				di = 0
			} else {
				di += 1
			}
		}
		if curr.x+direct[di][0] >= 0 && curr.x+direct[di][0] < xlen && curr.y+direct[di][1] >= 0 && curr.y+direct[di][1] < ylen && !visited[curr.x+direct[di][0]][curr.y+direct[di][1]] {
			q = append(q, coord{curr.x + direct[di][0], curr.y + direct[di][1]})
		}
	}

	return res
}
