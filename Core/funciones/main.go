package main

import (
	"errors"
	"fmt"
	"time"
)

/*
Generalmente una función tiene un tamaño fijo de parametros, pero podemos
apliarlo usando ... para declarar un slice del tipo elegido.
También podemos declarar más de elemento a retornar
*/

/*
Para manejar errores con mayor precisión, podemos usar custom errors.
Declaramos variables globales para reemplazar los error.new a lo largo del código.
*/

var (
	ErrDivByZero      = errors.New("Cannot divide by zero")
	ErrMustBePositive = errors.New("Numbers must be positive")
	ErrEmptyList      = errors.New("Numbers list is empty")
)

/*
Si queremos más información en nuestros errores, podemos crear un tipo
y extenderle la interfaz Error
*/

type OpError struct {
	Operation string
	Code      string
	Message   string
	Time      time.Time
}

func (op OpError) Error() string {
	return op.Message
}

func NewOpError(operation, code, message string, time time.Time) *OpError {
	return &OpError{
		Operation: operation,
		Code:      code,
		Message:   message,
		Time:      time,
	}
}

func sum(numbers ...int) (int, error) {
	result := 0
	error := error(nil)
	if len(numbers) == 0 {
		error = ErrEmptyList
	} else {
		for _, number := range numbers {
			result += number
		}
	}
	return result, error
}

func div(a, b int) (int, error) {
	error := error(nil)
	result := 0
	if b == 0 {
		error = ErrDivByZero
	} else if a < 0 || b < 0 {
		error = ErrMustBePositive
	} else {
		result = a / b
	}
	return result, error
}

func main() {
	val, err := sum(1, 2, 3, 4, 5)
	if err == ErrEmptyList {
		fmt.Println("empty list")
	} else {
		fmt.Println(val)
	}

	val, err = div(10, 2)
	if err == ErrDivByZero {
		fmt.Println("divide by zero")
	} else if err == ErrMustBePositive {
		fmt.Println("numbers must be positive")
	} else if err != nil {
		fmt.Println(val)
	}
}
