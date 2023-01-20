package main

import "fmt"

func main() {
	//#1. Hello World 출력
	fmt.Println("Hello World!!")

	//#2. String형 변수(문자) 출력
	var test1 string
	test1 = "Hello World2!!"
	fmt.Println(test1)

	//#3. uint8형 변수(숫자) 출력
	var number uint8 = 29
	number = number + 100
	fmt.Println(number)

}
