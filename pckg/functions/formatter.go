package functions

import "fmt"

//Format dfd
func Format(name string) string {
	message := fmt.Sprintf("yo, %v. yo!", name)
	return message
}
