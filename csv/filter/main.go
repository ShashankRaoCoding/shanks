package filter

import (
	"os"
	"bufio"
	"os/exec" 
)

// format: shanks csv filter by columnName where condition
func Main(args []string) error {
	var columnName string
	var condition string
	
	cmd := exec.Command(
		"python",
		"-c",
		fmt.Sprintf(script, columnName, condition),
	)

	
}

var script = `
import sys
import pandas
import io

data = pandas.read_csv(io.StringIO(sys.stdin.read())) 
data = data[data[%v]%v]
`








































