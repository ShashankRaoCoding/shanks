package main

import (
	"os" 
	"fmt" 
	"yoru/methods" 
)

func main() {

	process := os.Args[1] 
	
	m, ok := methods.Methods[process]
	if ok == false {
		fmt.Println("Function does not exist")
		os.Exit(1) 
	}

	err := m(os.Args[2:]) 

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) 
	} 

	os.Exit(0) 
}














































