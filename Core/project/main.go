package project

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

func printLevel(level LogLevel) {
	if level < 0 || level > LogLevel(len(levelNames)-1) {
		panic("Invalid log level")
	}
	fmt.Println(levelNames[level])
}

func main() {
	printLevel(LogDebug)
	printLevel(LogError)
	printLevel(LogWarning)
	printLevel(LogInfo)
}
