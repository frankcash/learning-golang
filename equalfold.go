package main

import(
	"fmt"
	"strings"
)


func main(){
	confirmedSeed := "foo"
	if strings.EqualFold("foobar", confirmedSeed) == true{
		fmt.Println("true")
	}else if strings.EqualFold("foo", confirmedSeed) == true{
		fmt.Println("true also")
	}


}
