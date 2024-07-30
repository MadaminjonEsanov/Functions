package main

import "fmt"

func main() {

	//simple expression function
	var fn = func () {
		fmt.Println("Hello,sami")
	}

	fn()

	fn = func () {
		fmt.Println("how are you, sami ")
	}
	fn()

	//with parameters

	var fn1 = func (num1,num2 int) {
		fmt.Println(num1 + num2)
	}
	fn1(3, 5)

	//with multiple parameters

	var fn2 = func (num1,num2 int) (int,int){
		return num1 + num2, num1 * num2
    
}
     r1,r2 :=fn2(3, 5)

	fmt.Println(r1,r2)
}