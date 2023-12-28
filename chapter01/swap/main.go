package main

import "fmt"

func main() {
	// cách sử dụng con trỏ để thực hiện hoán đổi giá trị giữa hai biến, giúp thay đổi trực tiếp giá trị của biến mà con trỏ trỏ đến.
	a, b := 5, 10
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
