package main

import "fmt"


func main() {



	HandleFunc("Home", f)
}

func HandleFunc(pattern string, handler func (req string, res string)) {

	var req = "salom"
	var res = "qalaaysan"

	handler(req, res)

	fmt.Println("Endpoint:" ,pattern)
}

func f(req string, res string) {
	fmt.Println(req, "->" ,res)
}