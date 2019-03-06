#!/usr/bin/env python3
# Python program for implementation of Bubble Sort 
  
def bubbleSort(arr): 
    n = len(arr) 
  
    # Traverse through all array elements 
    for i in range(n): 
        if i >= k:
            return
        # Last i elements are already in place 
        for j in range(0, n-i-1): 
  
            # traverse the array from 0 to n-i-1 
            # Swap if the element found is greater 
            # than the next element 
            if arr[j] > arr[j+1] : 
                arr[j], arr[j+1] = arr[j+1], arr[j] 
  

if "__main__" == __name__:
    # Driver code to test above 
    arr = [64, 34, 25, 12, 22, 11, 90] 
    k = 3
    
    print(arr)
    
    bubbleSort(arr) 
    
    print ("Sorted array is:", arr) 
    print("K largest number: k= {}, number={}".format(k, arr[-k:]))
