package dir


import (
	"strings"
	"fmt" 
	"os" 
)
func Main(dirnames [] string) error {
	var errors []string
	for _, name := range dirnames {
		err := os.Mkdir(name, 0755) 
		if err != nil {
			errors = append(errors, err.Error()) 
			continue 
		}
		
	}
	if len(errors) == 0 {
		
		return nil 
		
	} else {
return fmt.Errorf("There were errors: \n\t%v", strings.Join(errors, "\n\t")) 

	}
	
}







































