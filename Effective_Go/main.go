package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

func String(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	} else if str, ok := value.(Stringer); ok {
		return str.String()
	}

	return ""
}

func main() {
	var value interface{}
	value = "abe"


	str, ok := value.(string)
	if ok {
		fmt.Printf("string value is: %q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}

	fmt.Println(String(value))
}
