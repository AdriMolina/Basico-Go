package main

import (
	"fmt"
)

func main() {
	x := 25
	fmt.Println(x)
	cambiarValor(x)
	//ver direccion fisica de la variable
	fmt.Println(&x)
	y := &x
	fmt.Println(*y)
}

func cambiarValor(a int) {
	a = 36
}