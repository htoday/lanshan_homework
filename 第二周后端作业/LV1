package main

import "fmt"

type operation func(int, int) int

func add(num1 int, num2 int) int {
	return num1 + num2
}
func subtract(num1 int, num2 int) int {
	return num1 - num2
}
func multiply(num1 int, num2 int) int {
	return num1 * num2
}
func divide(num1 int, num2 int) int {
	return num1 / num2
}
func Calculator(num1 int, num2 int, CMD func(int, int) int) int {
	return CMD(num1, num2)
}
func main() {
	operators := map[string]operation{
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}
	var num1 int
	var num2 int
	var operator byte
	fmt.Println("请输入表达式（例如a+b）")
	for {
		_, err := fmt.Scanf("%d%c%d", &num1, &operator, &num2)
		if err != nil {
			fmt.Println("输入错误，请重新输入", err)
			continue
		} else {
			break
		}
	}
	s := string(operator)
	value, _ := operators[s]
	fmt.Println("结果为", Calculator(num1, num2, value))
}
