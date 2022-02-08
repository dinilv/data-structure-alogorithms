package main

import (
	"fmt"
)

//logic processor
func rotateArray(nPos int, list []int) {
	length := len(list)
	//scan through list for each rotation
	for j := 0; j < length-1; j++ {
		nIndex := (j+1)
		if nIndex < 0 {
			//reset the initial one to end of list
			nIndex = length + nIndex
		}
		//interchange value at each index
		list[nIndex], list[j] = list[j], list[nIndex]
	}

	fmt.Println("list", list)
}

func main() {
	rotateArray(2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}
