/*
# Writing tests
## Writing a test is just like writing a function, with a
 few rules
- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- In order to use the *testing.T type, you need to import
 "testing", like we did with fmt in the other file
*/

package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Dizendo ola para as pessoas", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		correctMessage(t, got, want)
	})

	t.Run("diz 'Hello, World' quando uma string vazia é inputada", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		correctMessage(t, got, want)
	})
}

func correctMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
