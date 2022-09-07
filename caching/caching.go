package main

func main() {
	matrix := [4][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}
	P := 4
	result := make([]int, 10)
	dim := 4
	for p := 0; p < P; p++ {
		go func(p int) {
			result[p] = 0
			size := dim/P + 1
			start := p * size
			end := 0
			if (start + size) > dim {
				end = dim
			} else {
				end = start + size
			}
			for i := start; i < end; i++ {
				for j := 0; j < dim; j++ {
					if matrix[p][i*dim+j]%2 != 0 {
						result[p]++
					}
				}
			}
		}(p)
	}
}

func execute(p int) {

}
