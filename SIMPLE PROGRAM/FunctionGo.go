package main

import (
	"errors"
	"fmt"
)

func somma(a int, b int) int {
	return a + b
}

func divisione(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("divisione per zero")
	}
	return a / b, nil
}

func main() {

	risultato := somma(5, 3)

	fmt.Println(risultato)

	risultatoDiv, err := divisione(10, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(risultatoDiv)
	}

	_, err = divisione(10, 0)

	if err != nil {
		fmt.Println(err)
	}
}
