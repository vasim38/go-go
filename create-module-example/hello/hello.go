package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set props for the predefined logger, including the log entry prefix
	log.SetPrefix("greetings: ")

	names := []string{"abc", "deef", "ghi"}

	// Get a greeting message and print it.
	//message, err := greetings.Hello("Vasim Vahora")
	//message, err := greetings.Hello("")
	message, err := greetings.Hellos(names)

	// If error occurs, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(message)
}
