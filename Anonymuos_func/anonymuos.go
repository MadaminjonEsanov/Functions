package main

import "fmt"

//simple anonymus function
func main() {

	func () {
		fmt.Println("Hello, world!")
	}()


//parametric anonymous function

res:=func (num1,num2 int) int {
	 return num1+num2
} (23,34)


fmt.Println(res)

//anonymous function with multiple

r1,r2 :=func (num1,num2 int) (int, int) {
	 return num1+num2, num1*num2

}(23,34)

fmt.Println(r1,r2)
}