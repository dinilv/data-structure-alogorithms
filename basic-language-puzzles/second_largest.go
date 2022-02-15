package main

import "fmt"

func findSecondLargest(list []int)int{
	//validate list length
	if len(list)==0{
		return 0
	}
	secondLargest:=list[0]
	largest:=list[0]
	for i:=1;i<(len(list));i++{
		if list[i]>largest{
			largest,secondLargest=list[i],largest
		}else if list[i]>secondLargest{
			secondLargest=list[i]
		}

	}
	fmt.Println("largest:",largest,secondLargest)
	return secondLargest
}
func main(){
	findSecondLargest([]int{1,8,4,5,2,3})
}
