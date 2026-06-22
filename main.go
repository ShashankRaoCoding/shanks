package main

import (
	"os" 
	"fmt" 
	"shanks/make" 
)

var functions = map[string]func([]string)error{
	"make": make.Main, 
	"csv": csv.Main, 
}

func main() {

	process := os.Args[1] 
	
	f, ok := functions[process]
	if ok == false {
		fmt.Println("Function does not exist")
		os.Exit(1) 
	}

	err := f(
		os.Args[2:], 
	) 

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) 
	} else {
		os.Exit(0) 
	}
}














































