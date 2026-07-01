package filter

import (
	"fmt"
	// "os"
	// "os/exec"
	// "yoru/globals" 
)

var Methods = make(map[string]func([]string) error) 

func Init() {
	Methods["cols"] = filterCols 
	Methods["rows"] = filterRows 
}

func Main(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("yoru csv filter: expected a method (rows|cols)")
	}

	method, rest := args[0], args[1:]

	m, ok := Methods[method]

	if !ok {
		return fmt.Errorf("Error: Method %s does not exist for yoru csv filter", method)
	}

	return m(rest)
}





















