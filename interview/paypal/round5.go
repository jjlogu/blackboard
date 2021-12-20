package main

import "fmt"

/*

    https://codebunk.com/b/4711100379196/

    input array

    output subarray where sum of is largest

    input:
    [1 2 3 -2 5] +> 
    [1 2 3 -4 5 6]

    output:
    [1 2 3]
    [5 6]


    1 2 3 

    index_start = 0
    index_end = len(array)

    max_sum =0

    sum_all_element= x
    max_sum = sum_all_element
    is_negative_present = true

    if is_negative_present = true {
    sum_of_arr=0
    for i=index_start; i<len(arr); i++ {

        if arr[i] >=0 
            sum_of_arr+=arr[i]
        else
            if sum_of_arr > max_sum {
                index_end = i-1
            }


    }
    }
    else {
    return arr
    }



    [1 2 3]
    i=2
    sum=3

    index_start=0
    index_end=3
    max_sum=6

    [1 2 -3]

    i=1
    sum=3

    inde_start=0
    index_end=1
    max_sum=3


    [1 2 3 4 5 7 8 9]

    [9 8 7 5 4 3 2 1]

*/


func max_sub_arr(arr []int) []int {
    index_start := 0
    index_end := len(arr)
    
    max_sum := 0
    j:=0
    sum:=0
    for i:=0;i<len(arr);i++ {
        sum=0
        for j=i; j<len(arr);j++ {
            sum+=arr[j]
        }
        if sum > max_sum {
            index_start=i
            index_end=j-1
            max_sum = sum
        }
    }
    
    for i:=len(arr)-1;i>0;i-- {
        sum =0
        for j=0; j<=i;j++ {
            sum+=arr[j]
        }
        if sum > max_sum {
            index_start=0
            index_end=i
            max_sum = sum
        }
    }
    return arr[index_start:index_end+1]
}


func main() {
    arr := []int{1,2, -4, 10};
    fmt.Printf("max_sub_arr %#v\n", max_sub_arr(arr));
}
