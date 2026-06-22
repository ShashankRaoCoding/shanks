package main

import (
	"fmt" 
	"os"
)

var functions = map[string]func([]string){
	"mkfile": mkfile, 
}

func main() {

	process := os.Args[1] 
	
	f, ok := functions[process]
	if ok == false {
		fmt.Println("Function does not exist")
		os.Exit(1) 
	}

	f(
		os.Args[2:], 
	) 
	
}


func mkfile(filenames []string) {
		
	// filenames := os.Args[1:]
	for _, i := range filenames {
		file, err := os.Create(i)
		if err != nil {
			fmt.Println(err.Error()) 
		} else {
			var bytes []byte 
			for range 40 {
				bytes = append(bytes, '\n') 
			}
			file.Write(bytes) 
			file.Close() 
		}
	}
	
}














































