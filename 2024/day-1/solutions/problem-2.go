package solutions

import (
	// "fmt"

	"github.com/YugandharrPatil/advent-of-code/2024/day-1/utils"
)

func Solution2() int {
	file1 := utils.OpenFile("list-1.txt")
	file2 := utils.OpenFile("list-2.txt")

	defer file1.Close()
	defer file2.Close()

	var arr1 []int
	var arr2 []int

	arr1 = utils.ScanFile(file1, arr1)
	arr2 = utils.ScanFile(file2, arr2)

	// fmt.Printf("unsorted array 1: %v\n", arr1)
	// fmt.Printf("unsorted array 2: %v\n", arr2)
	utils.SelectionSort(arr1)
	utils.SelectionSort(arr2)

	similarityScore := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if arr1[i] == arr2[j] {
				similarityScore += arr1[i]
			}
		}
	}

	// fmt.Printf("%v", similarityScore)

	return similarityScore
}
