package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entrada string, operador string) int {
	cleanImput := strings.Split(entrada, operador)

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
	case "-":
		fmt.Println(operador1 - operador2)
	case "*":
		fmt.Println(operador1 * operador2)
	case "/":
		fmt.Println(operador1 / operador2)
	default:
		fmt.Println(operador1 / operador2)
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println(operacion)
	operador := "-"
	valores := strings.Split(operacion, operador)
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])
	//Conertir a valores enteros (Atoi)
	operador1, err1 := strconv.Atoi(valores[0])
	//Valor nulo (nill)
	if err1 != nil {
		fmt.Println(err1)
	} else {

		fmt.Println(operador1)
	}
	operador2, _ := strconv.Atoi(valores[1])
	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
	case "-":
		fmt.Println(operador1 - operador2)
	case "*":
		fmt.Println(operador1 * operador2)
	case "/":
		fmt.Println(operador1 / operador2)
	default:
		fmt.Println(operador1 / operador2)
	}

}
