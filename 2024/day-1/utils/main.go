package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("error opening file: ", err)
	}
	return file
}

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
		}
	}
}

func ScanFile(file *os.File, slice []int) []int {
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Println("error converting string to number: ", err)
			continue
		}

		slice = append(slice, num)
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return slice
}
