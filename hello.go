package main

import "fmt"

const french = "French"
const spanish = "Spanish"

const englishPrefix = "Hello, "
const frenchPrefix = "Bonjour, "
const spanishPrefix = "Hola, "

func DefinePrefix(language string) string {
	var prefix string

	switch language {
	default :
		prefix = englishPrefix
	case "French":
		prefix = frenchPrefix
	case "Spanish":
		prefix = spanishPrefix
	}

	return prefix
}

func Hello(value string, language string) string {
	if value == "" {
		value = "World"
	}
	
	return DefinePrefix(language) + value
}

func main() {
	fmt.Println(Hello("", "English"))
}
