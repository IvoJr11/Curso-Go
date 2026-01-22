package main

import "fmt"

type LogLevel int

const (
	LogError LogLevel = iota
	LogWarning
	LogInfo
	LogDebug
)

var levelNames = []string{
	"ERROR",
	"WARNING",
	"INFO",
	"DEBUG",
}

// el metodo implementa la interfaz String, cuando se llame a algo.String() este es ejecutado
func (level LogLevel) String() string {
	if level < 0 || level > LogLevel(len(levelNames)-1) {
		panic("Invalid log level")
	}
	return levelNames[level]
}

func printLevel(level LogLevel) {
	// imprimir por pantalla definiendo un formato de salida
	fmt.Printf("Log level: %d %s\n", level, level.String())
}

func main() {
	printLevel(LogDebug)
	printLevel(LogError)
	printLevel(LogWarning)
	printLevel(LogInfo)
}
