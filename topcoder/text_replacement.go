package main

import (
	"bytes"
	"fmt"
)

func OxToTigerNonOptimised(message string){
	//convert to chars
	chars :=[]rune(message)
	//split to words for easy replacement
	var word bytes.Buffer
	var words []string
	for _,char:=range chars{
		if char!=' '{
			word.WriteString(string(char))
		} else{
			words=append(words,word.String())
			word=bytes.Buffer{}
		}
	}
	words=append(words,word.String())
	word=bytes.Buffer{}

	for _,eachWord:=range words{
		if eachWord=="ox"{
			word.WriteString("tiger")
		} else{
			word.WriteString(eachWord)
		}
	}
	fmt.Println(word.String())
}

func OxToTiger(message string){
	//convert to chars
	chars :=[]rune(message)
	length :=len(chars)
	textBuffer:=bytes.Buffer{}
	byPass:=false
	fmt.Println("length",length)
	for index,char:=range chars{
		if byPass {
			byPass=false
			continue
		}else{
			//check independent ox with space prefix to the term and suffix to the term
			if  ((index!=0 && chars[index-1]==' ') ||(index==0))&& char=='o'&& (index+1<=(length-1))&&chars[index+1]=='x' &&
				(index+1==(length-1)  || chars[index+2]==' '){
				//verify it doesnt have previous char so to independent
				textBuffer.WriteString(string("tiger"))
				byPass=true
			}else{
				textBuffer.WriteString(string(char))
			}

		}
	}
	fmt.Println(textBuffer.String())
}

func main(){
	OxToTiger("ox i love u ox ettan: ox")
}
