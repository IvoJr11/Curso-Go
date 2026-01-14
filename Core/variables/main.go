package variables

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

func main() {
	// se puede declarar e inicializar por separado
	var name string
	name = "String"

	// o se puede inicializar, infiriendo el tipo
	name2 := "String"

	// fmt se utiliza para operaciones de E/S
	fmt.Println(name, name2)
}
