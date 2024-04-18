package main

import (
	"fmt"
	"sort"
)

/*

Given an Array - [3,1,4,3] (Represents number of boxes in each stack)

	* Remove the boxes from each stack until all the stacks are of the same size as the smallest stack
	* But we can remove only as many boxes as to reduce to the next smallest stack in each go
	* Count number of times you removed boxes
	* [35, 20, 50, 75] -> [35, 20, 50, 50] -> [35, 20, 35, 35] -> [20, 20, 20, 20]

	For input [3,1,4,3]:

	    *
	*   * *         *   * *           * *              *
	* * * *  --->   *   * *  --->     * *   --->       *  --->
	* * * *         * * * *       * * * *        * * * *       * * * *
	ctr = 0         ctr = 1       ctr = 2        ctr = 3       ctr = 4

	Result is 4

*/

func getSecondMax(arr []int) (int, int, int) {
	var secondMax int
	var max int
	var min int
	min = arr[0]
	max = arr[0]
	secondMax = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			secondMax = max
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	return max, secondMax, min
}
func approach1(arr []int) {
	fmt.Println("approach1")

	var NoOfBoxRemoved, NoOfTimesBoxRemoved int
	var max, secondMax, min int
	for {
		max, secondMax, min = getSecondMax(arr)
		if secondMax == max {
			for i := 0; i < len(arr); i++ {
				if arr[i] > min {
					NoOfTimesBoxRemoved++
					NoOfBoxRemoved += arr[i] - min
					arr[i] = min
				}
				// fmt.Printf("i: %d arr: %#v, NoOfBoxRemoved: %d, secondMax: %d, NoOfTimesBoxRemoved: %d\n", i, arr, NoOfBoxRemoved, secondMax, NoOfTimesBoxRemoved)
			}
			break
		}
		for i := 0; i < len(arr); i++ {
			if arr[i] > secondMax {
				NoOfTimesBoxRemoved++
				NoOfBoxRemoved += arr[i] - secondMax
				arr[i] = secondMax
			}
			// fmt.Printf("i: %d arr: %#v, NoOfBoxRemoved: %d, secondMax: %d, NoOfTimesBoxRemoved: %d\n", i, arr, NoOfBoxRemoved, secondMax, NoOfTimesBoxRemoved)
		}
	}
	fmt.Printf("Number of boxes removed is %d\n", NoOfBoxRemoved)
	fmt.Printf("Number of times box removed is %d\n", NoOfTimesBoxRemoved)
}
func approach2(arr []int) {
	fmt.Println("approach2")
	sortedStacks := make([]int, len(arr))
	copy(sortedStacks, arr)
	sort.Ints(sortedStacks)
	NoOfTimesBoxRemoved := 0
	NoOfBoxRemoved := 0
	length := len(arr)
	i := length - 2
	exit := false
	for {
		// fmt.Println(i, length, sortedStacks)

		for j := 0; j < length; j++ {
			if arr[j] > sortedStacks[i] {
				NoOfTimesBoxRemoved++
				NoOfBoxRemoved += arr[j] - sortedStacks[i]
				arr[j] = sortedStacks[i]
			}
		}
		if exit || i == 0 {
			break
		}

		i -= 2
		if i < 0 {
			i = 0
			exit = true
		}
		// fmt.Println(arr, count)
	}
	fmt.Printf("Number of boxes removed is %d\n", NoOfBoxRemoved)
	fmt.Printf("Number of times box removed is %d\n", NoOfTimesBoxRemoved)
}
func main() {
	arr := []int{35, 20, 50, 75} //{3, 1, 4, 3}
	arr1 := make([]int, len(arr))
	copy(arr1, arr)

	approach1(arr)
	approach2(arr1)

}
