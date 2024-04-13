package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	//If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty Name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string](string))
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err

		}

		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	//A slice of greetings
	formats := []string{
		"Hi %v, welcome !",
		"Great to see you %v !",
		"Hail, %v ! How are you !",
	}

	// Return random message by specifying random index of slice of formats
	return formats[rand.Intn(len(formats))]
}
