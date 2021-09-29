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

// Hellos returns a map that associates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string) // map은 해시테이블(Hash table)을 구현한 자료구조, 형식은 map[Key타입]Value타입
	// Logo through the received slice of names, calling
	// the Hollo function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message;
	}
	return messages, nil
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