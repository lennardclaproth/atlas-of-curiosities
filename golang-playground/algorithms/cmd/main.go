package main

import SortingAlgorithms "github.com/lennardclaproth/golang-playground/algorithms/pkg/sorting_algorithms"

func main() {
	myArr := []int{1, 2, 3, 4, 5}
	SortingAlgorithms.SelectionSort[int](myArr)
}