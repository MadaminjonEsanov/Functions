package main

import "fmt"

func  main() {
	
	var n int =5

	fn :=foo(n)
	res := fn()

	fmt.Println(res)
}

func foo(num int) func() int {
	return func() int {
		return num * num
	}

}