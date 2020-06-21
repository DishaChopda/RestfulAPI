package main
import "fmt"

func Calculate(a int)(result int){
	result=a+5
	return result
}

func main(){
answer := Calculate(2)
fmt.Println(answer)
}