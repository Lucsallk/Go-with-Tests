package main

import "fmt"

func Hello(value string) string {
	return "Hello, " + value
}
func main() {
	fmt.Println(Hello("Lucas"))
}
