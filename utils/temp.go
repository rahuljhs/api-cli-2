package utils

import (
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data := Person{
		Name: "John Doe",
		Age:  30,
	}

	err := PerformPost("https://example.com/api/persons", data)
	if err != nil {
		fmt.Printf("Failed to perform POST request: %v\n", err)
		return
	}

	fmt.Println("POST request completed successfully!")
}
