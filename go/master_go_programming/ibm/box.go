package main

import "fmt"

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

func main() {
	arr := []int{3, 1, 4, 3}

	var count, count1 int
	var max, secondMax, min int
	for {
		max, secondMax, min = getSecondMax(arr)
		if secondMax == max {
			for i := 0; i < len(arr); i++ {
				if arr[i] > min {
					count1++
					count += arr[i] - min
					arr[i] = min
				}
				// fmt.Printf("i: %d arr: %#v, count: %d, secondMax: %d, count1: %d\n", i, arr, count, secondMax, count1)
			}
			break
		}
		for i := 0; i < len(arr); i++ {
			if arr[i] > secondMax {
				count1++
				count += arr[i] - secondMax
				arr[i] = secondMax
			}
			// fmt.Printf("i: %d arr: %#v, count: %d, secondMax: %d, count1: %d\n", i, arr, count, secondMax, count1)
		}
	}
	fmt.Printf("Number of boxes removed is %d\n", count)
	fmt.Printf("Number of times box removed is %d\n", count1)
}
