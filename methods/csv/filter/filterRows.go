package filter

import (
	"fmt"
	"os"
	"os/exec"
	// "yoru/globals" 
)


// format: shanks csv filter rows columnName condition
// e.g.:   shanks csv filter rows age ">30"
func filterRows(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("yoru csv filter rows: expected <columnName> <condition>")
	}

	columnName := args[0]
	operator:= args[1]
	value := args[2] 

	cmd := exec.Command(
		"python",
		"-c",
		fmt.Sprintf(rowsScript, columnName, operator, value),
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("yoru csv filter: There was an error: %s", err)
	}

	return nil
}



var rowsScript = `
import sys
import pandas

data = pandas.read_csv(sys.stdin)
data = data[data['%s'] %s '%s']
data.to_csv(sys.stdout, index=False)
`


