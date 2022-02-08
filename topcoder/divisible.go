package main

func change(number int) int{
//check is divisible by 7 or not
	remainder:=(number)%7
	if remainder==0{
		return number
	}else{
		if (7+remainder) <10{
			return number+remainder
		}else{
			return number-remainder
		}
	}
	}