package main

import (
	"1tap/mathutils"
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Eki san engiz: ")
	fmt.Scan(&a, &b)

	fmt.Println(mathutils.Multiply(a, b))

}
