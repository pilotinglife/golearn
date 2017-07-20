package main

import "fmt"

func main() {

	a := 1

	var ret int

	ret = max(a, 3)

	fmt.Printf("a is:%d, ret is:%d\n", a, ret)

	//variation on declaring variables

	var c, d int

	c = 45
	d = 65

	fmt.Printf("%d is the max value\n", max(c, d))

	//OR allow dynamic typing

	ret2 := max(a, c)
	fmt.Printf("dynamic return type %d,   the type is %T\n", ret2, ret2)


	////What type is it?
	//var x interface{}
	//x = Type(ret2)
	//fmt.Printf("can I get the type this way? type is: %n \n", x)



	//return multiple variables from a function

	a2, b2 := swap("Mahesh", "Kumar")
	fmt.Println(a2, b2)

}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
	/* local variable declaration */
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
