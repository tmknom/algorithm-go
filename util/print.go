package util

import "fmt"

const indentSpace = "  "

func CallPrintf(indent int, format string, a ...interface{}) {
	indentPrintf("call ", indent, format, a...)
}

func ReturnPrintf(indent int, format string, a ...interface{}) {
	indentPrintf("return ", indent, format, a...)
}

func indentPrintf(keyword string, indent int, format string, a ...interface{}) {
	prefix := ""
	for i := 0; i < indent; i++ {
		prefix += indentSpace
	}
	prefix += keyword

	fmt.Printf(prefix+format, a...)
}
