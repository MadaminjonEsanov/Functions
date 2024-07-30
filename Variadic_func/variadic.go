package main
import "fmt"

func main() {

	max := GetMax(1,1,22,3,34,45,3,4,5,)
	fmt.Println(max)
}

func GetMax(numbers ...int) int {
	if numbers == nil {
		return 0
	}

	max := numbers[0]

	for _,e :=range numbers {
		if max<e {
			max =e
		}
		
	}
	return max
}