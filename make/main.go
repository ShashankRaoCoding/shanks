package make

import (
	"fmt" 
	"shanks/make/dir" 
	"shanks/make/file" 
)

var makeMap = map[string]func([]string)error{
	"file": file.Main, 
	"files": file.Main, 
	"dir": dir.Main,
	"dirs": dir.Main, 
}

func Main(args []string) error {
	var makeType string
	makeType, args = args[0], args[1:] 
	f, ok := makeMap[makeType]
	if ok == false {
		return fmt.Errorf("Unrecognised Command: %v is not a valid make argument")
	} 

	err := f(args) 
	if err == nil {
		return nil 
	} else {
		return fmt.Errorf("There was an error: %v", err.Error()) 
	}
		
}












































