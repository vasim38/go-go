package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set props for the predefined logger, including the log entry prefix
	log.SetPrefix("greetings: ")
	// Get a greeting message and print it.
	message, err := greetings.Hello("")
	// If error occurs, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(message)
}
