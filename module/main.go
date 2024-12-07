package main

import (
	"fmt"

	greeting "module/greeting" // Import the required package
)

func main() {
	// Get a greeting message and print it.
	message := greeting.Hello("Gladys")
	fmt.Println(message)
}
