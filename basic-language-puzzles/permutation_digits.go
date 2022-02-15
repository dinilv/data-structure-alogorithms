package main

import "fmt"

func Permutation(data,list []int, cid, pos,length int,) {
	if pos==cid {
		fmt.Println("permuted list",list)
		return
	}
	for j := cid; j < length; j++ {
		list[cid]=data[j]
		swap(data, cid, j)
		Permutation(data,list, cid+1,pos, length)
		swap(data, cid, j)
	}
}


func swap(data []int, x int, y int) {
	//fmt.Println("swap",x,y)
	data[x], data[y] = data[y], data[x]
	}

func main() {
	var data [4]int
	var list []int
	j:=0
	for i := 7; i <11; i++ { data[j] = i
		j++
	}
	Permutation(data[:],list, 0, 3,4)
}