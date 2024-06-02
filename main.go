package main

import "fmt"

func main(){
	fmt.Println("Hello, World!")
	var x int = 5;
	var y int = 7;
	var sum int = x+y;
	fmt.Println(sum);
	fmt.Println(x+y);
	var cond bool = true;
	if cond {
		fmt.Println("True");
	} else if !cond {
		fmt.Println("False");
	}
}