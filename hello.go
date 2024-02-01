package main

import "fmt"

const englishPrefix = "Hello, "

func Hello(value string, language string) string {
	if value == "" {
		value = "World"
	}
	return englishPrefix + value
}

func main() {
	fmt.Println(Hello("", ""))
}
