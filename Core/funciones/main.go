package main

import (
	"errors"
	"fmt"
)

/*
generalmente una función tiene un tamaño fijo de parametros, pero podemos
apliarlo usando ... para declarar un slice del tipo elegido.
También podemos declarar más de elemento a retornar
*/
func sum(numbers ...int) (int, error) {
	result := 0
	error := error(nil)
	if len(numbers) == 0 {
		error = errors.New("Numbers list is empty")
	} else {
		for _, number := range numbers {
			result += number
		}
	}
	return result, error
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
