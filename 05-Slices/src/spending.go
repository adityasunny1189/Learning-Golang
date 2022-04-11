package src

import (
	"fmt"
	"strconv"
	"strings"
)

func Fetch() [][]int {
	content := `12 23 44
	34 56 67
	87
	123 90
	78
	`
	lines := strings.Split(content, "\n")
	spendings := make([][]int, len(lines))
	for i, daily := range lines {
		spending := strings.Fields(daily)
		spendings[i] = make([]int, len(spending))
		for j, amt := range spending {
			spendings[i][j], _ = strconv.Atoi(amt)
		}
	}
	return spendings
}

func Spending() {
	myspendings := make([][]int, 0, 5)
	myspendings = append(myspendings, []int{12, 67, 89})
	myspendings = append(myspendings, []int{123}, []int{65, 89}, []int{56})
	myspendings = append(myspendings, []int{34})

	for i, daily := range myspendings {
		var total int

		for _, spending := range daily {
			total += spending
		}
		fmt.Println(i+1, total)
	}
}
