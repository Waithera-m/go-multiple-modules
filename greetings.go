package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Please provide a name")
	}
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func randomFormat() string {
	formats := []string {
		"Bonjour, %v",
		"Hallo, %v",
		"Uhiana atia %v",
	}

	return formats[rand.Intn(len(formats))]
}