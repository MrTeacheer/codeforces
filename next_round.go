package main


import "fmt"


func main(){
	var n,k,ni int
	var array []int
	fmt.Scanln(&n,&k)
	k--

	for i:=0;i<n;i++{
		fmt.Scan(&ni)
		array = append(array,ni)
	}
	arrayl:= array[:k]
	arrayr := array[k:]
	var right int
	for ind := range arrayr{
		switch{
		case array[k]!=0 && arrayr[ind]==array[k]:
			right++
		default:
			break
		}
	}
	var left int
	for ind := range arrayl{
		switch{
		case arrayl[ind]>=array[k] && arrayl[ind] != 0:
			left++
		}
	}
	fmt.Println(left+right)

}