package main

import (
	"fmt"
	"reflect"
)

func inspectVariable(variable interface{}) {
	t := reflect.TypeOf(variable)
	v := reflect.ValueOf(variable)
	fmt.Println("Type:", t)
	fmt.Println("Value:", v)
}
func main() {
	myVar := 42
	inspectVariable(myVar)
}