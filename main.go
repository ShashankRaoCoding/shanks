package main

import (
	"os" 
	"fmt" 
	mk "yoru/make" 
	"yoru/csv"
)

var Methods = make(map[string]func([]string)error) 

func init() {
	Methods["make"] = mk.Main
	Methods["csv"] = csv.Main
}

func main() {

	args := os.Args[1:] 

	method, args := args[0], args[1:] 
	
	m, ok := Methods[method]
	if ok == false {
		fmt.Printf("Function %s does not exist \n", method)
		os.Exit(1) 
	}

	err := m(os.Args[2:]) 

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) 
	} 

	os.Exit(0) 
}














































