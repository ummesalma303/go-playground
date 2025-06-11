package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)


func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("name not found 🧐")
    }


   // message := fmt.Sprintf("Hi, %v. Welcome!", name)
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}


func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v!🐎. Welcome!",
        "Great to see you, %v! 🦙",
        "Hail, %v!🐴. Well met!",
        "Life is a %v!🦋",
        "Everything is %v!.🥚🥚🥚",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
