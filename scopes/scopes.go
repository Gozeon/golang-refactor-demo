package main

import "fmt"

// import "reflect"

func f() {}

var g = "g"

// func main() {
// 	f := "f"
// 	fmt.Println(f) // 局部变量覆盖了包级函数
// 	fmt.Println(g) // 包级变量
// 	// fmt.Println(h) // error
// }

// func main() {
// 	x := "hello!"
// 	for i := 0; i < len(x); i++ {
// 		x := x[i]
// 		if x != '!' {
// 			fmt.Println(reflect.TypeOf(x))
// 			x := x + 'A' - 'a'
// 			fmt.Printf("%c", x) // HELLO
// 		}
// 	}
// }

func main() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // HELLO
	}
}
