package example

import "fmt"

func main(){
	fmt.Println("Hello, world")
	var(
		name = "r00ty"
		age = 23
	)

	location := "Blahastan"

	fmt.Printf("%s (%d) of %s", name, age, location)
}
