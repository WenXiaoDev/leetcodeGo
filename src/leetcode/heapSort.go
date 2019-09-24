package leetcode

import (
	"fmt"
	"math/rand"
	"time"
)

func heapify(arr []int, size int) {
	current := 0
	for current*2+1 < size {
		temp := current*2+1
		if temp+1 < size && arr[temp+1] > arr[temp] {
			temp = temp+1
		}
		if arr[current] < arr[temp] {
			arr[current],arr[temp] = arr[temp],arr[current]
		}
		current = temp
	}
}

func buildMaxHeap(arr []int) {
	for size := 1; size <= len(arr); size++ {
		current := size-1
		for current != 0 {
			f := (current-1)/2
			if arr[f] >= arr[current] {
				break
			}
			arr[f],arr[current] = arr[current],arr[f]
			current = f
		}
	}
}

func HeapSort(arr []int) {
	buildMaxHeap(arr)
	for size := len(arr); size > 0; size-- {
		arr[0],arr[size-1] = arr[size-1],arr[0]
		heapify(arr,size-1)
	}
}

func Test() {
	var array []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		array = append(array,rand.Intn(100))
	}
	fmt.Println(array)
	HeapSort(array)
	fmt.Println(array)
}