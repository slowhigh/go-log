package main

import (
	"fmt"
	"log"
	//"os"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Gladys")
	//If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)

		// "log.Fatal(err)"와 같은 결과를 
		//log.Println(err)
		//os.Exit(1)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}

//go run      (X)
//go run .    (O)
