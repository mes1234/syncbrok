package main

import (
	"fmt"

	"github.com/mes1234/syncbrok/internal/functions"
)

//Hello dfd
func Hello(name string) string {
	message := fmt.Sprintf("czesc, %v. ", name)
	return message
}

func main() {
	// Get a greeting message and print it.
	message := Hello(functions.Format("Gladys"))
	fmt.Println(message)
}
