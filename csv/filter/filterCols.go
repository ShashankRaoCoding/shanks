package filter

import (
	"fmt"
	"os"
	"os/exec"
	"strings" 
	// "yoru/globals" 
)


// format: shanks csv filter cols col1,col2,col3
func filterCols(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("yoru csv filter cols: expected a comma-separated list of column names")
	}

	cols := strings.Split(args[0], ",")
	for i, c := range cols {
		cols[i] = fmt.Sprintf("'%s'", strings.TrimSpace(c))
	}
	colList := "[" + strings.Join(cols, ", ") + "]"

	cmd := exec.Command(
		"python",
		"-c",
		fmt.Sprintf(colsScript, colList),
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("yoru csv filter: There was an error: %s", err)
	}

	return nil
}

var colsScript = `
import sys
import pandas

data = pandas.read_csv(sys.stdin)
data = data[%s]
data.to_csv(sys.stdout, index=False)
`
