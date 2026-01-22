package main

import "fmt"

// variables globales/constantes
const STRING string = "String"

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

type LogType int

const (
	// iota nos permite heredar el mismo tipo en las variables siguientes
	// además de agregar un valor por defecto
	LogError LogType = iota
	LogWarning
	LogInfo
	LogDebug
)

// array estandar, tamaño y tipo fijo
var numbers [2]int = [2]int{1, 2}

// slice, tamaño variable
var slice = []string{"a", "b", "c"}

// slice con capacidad (10) e inicializaciones (3)
// cuando se excede la capacidad del slide, esta se duplica para poder añadir los nuevos elementos
// haciendo append(items, n) 3 veces, tendríamos 6 elementos y 5*2=10 de capacidad
var items = make([]int, 3, 5)

// un map asocia un par de clave-valor
// std, ok := studentGrades[string] nos indica si el valor retornado es seguro de usar.
// en caso de que la clave no exista, ok es false y std es el valor nulo del retorno
var studentGrades = map[string]int{
	"student1": 90,
	"student2": 70,
	"student3": 65,
	"student4": 40,
}

func main() {
	// se puede declarar e inicializar por separado
	var name string
	name = "String"

	// o se puede inicializar, infiriendo el tipo
	name2 := "String"

	// fmt se utiliza para operaciones de E/S
	fmt.Println(name, name2)
}
