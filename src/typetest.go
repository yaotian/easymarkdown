package main

import "fmt"

type test_int func (int) bool

//申明了一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

//申明的函数类型在这个地方当做了一个参数
func filter(slice []int, f test_int) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}


func main() {
	fmt.Println("haha")
	slice := []int {1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) //函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)//函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}
