//package main

import "fmt"

func cambioValore(ptr *int) {
	*ptr = 100
}
func main() {
	x := 10
	fmt.Println(x)
	cambioValore(&x)
	fmt.Println(x)

}
