package main

import "fmt"



func main(){
	var input string
	var amount,result int
	fmt.Scanln(&amount)
	for i:=1;i<=amount;i++{
		fmt.Scanln(&input)
		switch{
		case rune(input[1]) == '+':
			result+=1
		default:
			result-=1
		}
	}
	fmt.Println(result)
}