package main

import "fmt"

func modifyPointerValue(num *int) {
	if num == nil {
		fmt.Println("The number is nil")
		return
	}
	*num = *num * 10
}

func main() {
	/* los punteros son variables que no almacenan valores en memoria, sino direcciones
	es importante tener en cuenta que al implementar punteros, pueden haber "efectos secundarios"
	cuando se llaman a funciones que puedan cambiar las direcciones */
	num := 20
	modifyPointerValue(&num)
	fmt.Println(num)
}
