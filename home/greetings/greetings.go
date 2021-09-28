package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) { // 2개의 값을 반환할 수 있다. 반환 값은 ,(컴마)로 구분한다.
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	// in a greeting message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil // nil => 오류 없음을 의미
}


// init sets initial values for variables used in the function.
func init()  {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting message. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}