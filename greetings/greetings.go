package greetings

import (
    "errors"
    "fmt"
)


func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("name not found 🧐")
    }


    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
