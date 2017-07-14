package example

import (
	"fmt"
	"os"
	"strings"
)

func main(){

	s, sep := "",""
	for index, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
		fmt.Println(index)
		fmt.Println(arg)
	}
	fmt.Println("Command" + os.Args[0])
	fmt.Println(s)

 fmt.Println(strings.Join(os.Args[1:], " "))
}
