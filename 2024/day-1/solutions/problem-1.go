package solutions

import (
	// "fmt"
	"math"

	"github.com/YugandharrPatil/advent-of-code/2024/day-1/utils"
)

func Solution1() int {
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
	// fmt.Printf("sorted array 1: %v\n", arr1)
	// fmt.Printf("sorted array 2: %v\n", arr2)

	var distance float64 = 0
	for i := 0; i < 1000; i++ {
		distance += math.Abs(float64(arr2[i]) - float64(arr1[i]))
	}
	// fmt.Printf("%f", distance)

	return int(distance)
}
