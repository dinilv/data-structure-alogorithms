package main

import "fmt"
/*
--Solve the puzzle
--Given a list of positive and negative integers, find a contiguous subarray whose sum (sum of elements) is maximum.
*/

//Solution use nested scan and use pivot to split full array into
func maxOfArray(numbers []int)int{
	//validate array is empty or not
	arrayLength:=len(numbers)
	if arrayLength==0{
		fmt.Println("Empty number array found.")
		return 0
	}
	maxSum:=numbers[0]
	maxArray:=[]int{numbers[0]}
	//loop through array
	for i:=1;i<arrayLength;i++{
		sumSoFar:=0
		subArray:=[]int{}
		//sum till the pivot not including till whole array
		for k:=0;k<=(i-1);k++{
			subArray=append(subArray,numbers[k])
			//calculate sum of all sub array
			sumSoFar=numbers[k]+sumSoFar
			//compare to max so far
			if sumSoFar>maxSum{
				maxSum=sumSoFar
				maxArray=subArray
			}
		}
		//reset of pivot:sum & array
		sumSoFar=0
		subArray=[]int{}
		//sum from pivot to rest of the combination
		for j:=i;j<arrayLength;j++{
			subArray=append(subArray,numbers[j])
			//calculate sum of all sub array
			sumSoFar=numbers[j]+sumSoFar
			//compare to max so far
			if sumSoFar>maxSum{
				maxSum=sumSoFar
				maxArray=subArray
			}
		}
		//compare to individual element
		if numbers[i]>maxSum{
			maxSum=numbers[i]
			maxArray=[]int{numbers[i]}
		}
	}
	fmt.Println("Max sum of array:",maxSum,"-:-",maxArray)
	return maxSum
}
func main(){
	//call procedue with array 1
	data := []int{1, -2, 3, 4, -4, 6, 14, 8, 2}
	maxOfArray(data)
	//call procedue with array 2
	data = []int{1, 2,32}
	maxOfArray(data)
	//call procedue with array 3
}
