package main

import (
	"fmt"
)

const s string = "constant"

func plus(a int, b int) int {
	
	return a + b
}

func main() {
	// fmt.Println("go" + "lang")

	// fmt.Println("1+1 = ", 1+1)
	// fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)

	// var a = "initial"
	// fmt.Println(a)

	// var b, c int =1, 2
	// fmt.Println(b, c)

	// var d = true
	// fmt.Println(d)

	// var e int
	// fmt.Println(e)

	// f := "apple"
	// fmt.Println(f)

	// fmt.Println(a + f)

	// fmt.Println(s)
	
	// const n = 500000000
	// const d = 3e20 / n

	// fmt.Println(d)

	// fmt.Println(int64(d))

	// fmt.Println(math.Sin(n))

	res := plus(1, 2)
	fmt.Println("1+2 = ", res)


}

