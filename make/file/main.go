package file

import (
	// "strings"
	"fmt" 
	"os/exec" 
	"os" 
	// "yoru/globals" 
)

func Main(filenames [] string) error {
	var err error 
	for _, name := range filenames {
		cmd := exec.Command("sh", "-c", fmt.Sprintf("python -c 'for _ in range(45): print()' > %s", name)) 
		cmd.Stdout = os.Stdout 
		cmd.Stderr = os.Stderr 
		err = cmd.Start() 
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Could not create %s due to the error %s\n", name, err) 
			continue 
		}
		err = cmd.Wait() 
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Could not create %s due to the error %s\n", name, err) 
			continue 
		}
	}

	return nil 
}






































