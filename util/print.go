package util

import "fmt"

const indentSpace = "  "

func IndentPrintf(indent int, format string, a ...interface{}) {
	prefix := ""
	for i := 0; i < indent; i++ {
		prefix += indentSpace
	}

	fmt.Printf(prefix+format, a...)
}
