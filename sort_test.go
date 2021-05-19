/*
 add by jinchaozhang. 2021/4/22
*/
package alg_test_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func selectSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	arrLength := len(arr)
	for start := 0; start < arrLength-1; start++ {
		minIndex := start
		for j := start + 1; j < arrLength; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, start, minIndex)
	}
}

func insertSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	arrLength := len(arr)
	for index := 1; index < arrLength; index++ {
		for preIndex := index - 1; preIndex >= 0 && arr[preIndex] > arr[preIndex+1]; preIndex-- {
			swap(arr, preIndex, preIndex+1)
		}
	}
}

func bubbleSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for end := len(arr) - 1; end > 0; end-- {
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
		}
	}
}

func swap(arr []int, num1Index, num2Index int) {
	temp := arr[num2Index]
	arr[num2Index] = arr[num1Index]
	arr[num1Index] = temp
}

func randomArray(maxLen, maxValue int) []int {
	length := rand.Intn(maxLen)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxValue)
	}
	return arr
}

func checkSort(arr []int) bool {
	pre := -1
	for i := 0; i < len(arr); i++ {
		if pre > arr[i] {
			return false
		}
		pre = arr[i]
	}
	return true
}

func TestRandomCompareEquals(t *testing.T) {

	fmt.Println("start")
	for i := 0; i < 1000000; i++ {
		arr := randomArray(15, 200)
		selectSort(arr)
		ok := checkSort(arr)
		if !ok {
			fmt.Println(arr)
			break
		}
	}
	fmt.Println("done")
}
