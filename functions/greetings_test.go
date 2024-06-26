package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile("Hi, "+name+". Welcome!")
    msg := Hello("Gladys")
    if !want.MatchString(msg) {
        t.Fatalf(`Hello("Gladys") = %q,want match for %#q, nil`, msg, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg := Hello("")
    if msg != "Hi, . Welcome!" {
        t.Fatalf(`Hello("") = %q, want "", error`, msg)
    }
}