/*
 add by jinchaozhang. 2021/4/22
*/
package alg_test_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func find(arr []int, num int) bool {
	len := len(arr)
	if len == 0 {
		return false
	}
	l := 0
	r := len - 1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == num {
			return true
		}
		if arr[mid] < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func findNearLeftIndex(arr []int, num int) int {
	len := len(arr)
	if len == 0 {
		return -1
	}
	l := 0
	r := len - 1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] >= num {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func testFind(arr []int, num int) bool {
	for _, val := range arr {
		if val == num {
			return true
		}
	}
	return false
}

func testFindNearLeftIndex(arr []int, num int) int {
	for i, val := range arr {
		if val >= num {
			return i
		}
	}
	return -1
}

func TestBinSearch(t *testing.T) {
	fmt.Println("start")
	maxLen := 15
	maxVal := 200
	for i := 0; i < 10000000; i++ {
		arr := randomArray(maxLen, maxVal)
		insertSort(arr)
		num := rand.Intn(maxVal)
		if findNearLeftIndex(arr, num) != findNearLeftIndex(arr, num) {
			fmt.Println(arr)
			break
		}
	}
	fmt.Println("done")
}

func findPartMinIndex(arr []int) int {
	len := len(arr)
	if len == 0 {
		return -1
	}
	if len == 1 {
		return 0
	}
	if arr[0] < arr[1] {
		return 0
	}
	if arr[len-1] < arr[len-2] {
		return len - 1
	}
	l := 0
	r := len - 1
	for l < r-1 {
		mid := (l + r) / 2
		if arr[mid] < arr[mid-1] && arr[mid] < arr[mid+1] {
			return mid
		}
		if arr[mid] > arr[mid-1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if arr[l] < arr[r] {
		return l
	}
	return r
}

func randomPartArray(maxLen, maxValue int) []int {
	length := rand.Intn(maxLen)
	arr := make([]int, length)
	if length > 0 {
		arr[0] = rand.Intn(maxValue)
		i := 1
		for i < length {
			arr[i] = rand.Intn(maxValue)
			if arr[i] != arr[i-1] {
				i++
			}
		}

	}
	return arr
}

func checkPartArry(arr []int, index int) bool {
	if len(arr) == 0 {
		return index == -1
	}
	leftIndex := index - 1
	rightIndex := index + 1
	leftBiger := true
	rightBiger := true
	if leftIndex >= 0 {
		leftBiger = arr[index] < arr[leftIndex]
	}
	if rightIndex < len(arr) {
		rightBiger = arr[rightIndex] > arr[index]
	}
	return leftBiger && rightBiger
}

func TestBinPartSearch(t *testing.T) {
	fmt.Println("start")
	maxLen := 15
	maxVal := 200
	for i := 0; i < 10000000; i++ {
		arr := randomPartArray(maxLen, maxVal)
		index := findPartMinIndex(arr)
		if !checkPartArry(arr, index) {
			fmt.Println(index)
			fmt.Println(arr)
			break
		}
	}
	fmt.Println("done")
}
