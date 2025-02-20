package slice

import "fmt"

func PrintSlice(slice []int) {
	for i, val := range slice {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d", val)
	}
	fmt.Println()
}

func PrintSlice2d(slice [][]int) {
	for _, rows := range slice {
		for i, val := range rows {
			if i > 0 {
				fmt.Print("\t")
			}
			fmt.Printf("%d\t", val)
		}
		fmt.Println("\n")
	}
}
