package file


import (
	"strings"
	"fmt" 
	"os" 
)
func Main(filenames [] string) error {
	var errors []string
	for _, name := range filenames {
		file, err := os.Create(name) 
		if err != nil {
			errors = append(errors, err.Error()) 
			continue 
		}
		i := "" 
		for range 45 {
			i += "\n" 
		}
		file.Write(
			[]byte(i), 
		)
		file.Close() 
	}
	if len(errors) == 0 {
		
		return nil 
		
	} else {
return fmt.Errorf("There were errors: \n\t%v", strings.Join(errors, "\n\t")) 

	}
	
}







































