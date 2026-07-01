package methods

import (
	
	"yoru/methods/csv" 
	mk "yoru/methods/make" 
)

var Methods = make(map[string]func([]string)error)

func Init() {
	Methods["mk"] = mk.Main
	Methods["csv"] = csv.Main 
}














































