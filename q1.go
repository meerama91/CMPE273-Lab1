package main 

import "fmt"

func fib(n int) int {

	if n==0 {
		return 0
	} else if n==1  {
		return 1
	} else {
	    return (fib(n-1) + fib(n-2))
	 }
}
func main() {
	var inp int
	fmt.Println("enter value of n : ")
	fmt.Scanf("%d", &inp)
	fmt.Println(fib(inp));
}
