package main

import "fmt"

func goto_loop() {
	i := 0
Here:
	fmt.Println(i)
	i++
	if i < 10 {
		goto Here
	}

}

// func main() {
// 	fmt.Printf("Hello, world.")

// 	var arr [10]int

// 	for i := 0; i < 10; i++ {
// 		arr[i] = i
// 	}

// 	fmt.Print(arr)

// 	goto_loop()

// }

func square(x *float64) {
	*x = *x * *x

}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp

}

func main() {
	// x := 1.5
	// square(&x)
	// fmt.Print(x)

	x, y := 1, 2
	swap(&x, &y)
	fmt.Print(x, y)
}
