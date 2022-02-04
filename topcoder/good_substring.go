package main

import (
	"fmt"
	"log"
)

func MaxGoodSubstring(statement string){
	chars := []rune(statement)
	var subString string=""
	var count=0
	for _,char:=range chars{
		subString+=string(char)
		findCombination(subString,&count)
	}
	log.Println("count",count)

	//-all single letters are identical
	//-multiple same letters join and form good strings
	//-join . elements to formed good strings and count

}

func findCombination(subString string,sum *int){
	fmt.Println("Received string for combination making:",subString)
	chars := []rune(subString)

	for i:=0;i<len(chars);i++{
		var combination string=""
		combination=string(chars[i])+string(chars[i:])+string(chars[:i])
		fmt.Println("Combinations:-",i,combination)
		*sum=*sum+1
	}

}
func main(){
	MaxGoodSubstring("abacd")
}