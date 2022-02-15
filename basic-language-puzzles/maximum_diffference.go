package main

import "fmt"

//find maximum & minimum in one scan
func findMinMaxBruteforce(list []int){
	length:=len(list)
	min,max:=list[0],list[0]
	for i:=1;i<length;i++{
		//compare to current min & max
	}
	fmt.Println("You can buy at:",min)
	fmt.Println("You can sell at:",max)
}

func findMinMaxSorted(list []int){
	//sort in increasing order
	//take out first index and last to find difference
}

func main(){
	//add validation for list minimum 2 entries
}
