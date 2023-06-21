package main

import "fmt"

func Algorithm(a, b string) int {
	matrix := [][]int{}
	for i := 0; i < len(a); i++ {
		matrix = append(matrix, make([]int, len(b)+1))
	}
	for i := 0; i < len(a); i++ {
		matrix[i][0] = i
	}
	for i := 0; i < len(b); i++ {
		matrix[0][i] = i
	}

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(b); j++ {
			substationCost := 0
			if a[i-1] != b[j-1] {
				substationCost = 1
			} else {
				substationCost = 0
			}
			cost1 := matrix[i-1][j] + 1
			cost2 := matrix[i][j-1] + 1
			cost3 := matrix[i-1][j-1] + substationCost
			minimum := cost1
			if cost2 < minimum {
				minimum = cost2
			}
			if cost3 < minimum {
				minimum = cost3
			}
			matrix[i][j] = minimum
		}
	}
	return matrix[len(a)-1][len(b)-1]
}

func main() {
	fmt.Println(Algorithm("kitten", "sitting"))
}
