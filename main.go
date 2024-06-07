package main
import (
  "fmt"
  "greetings/functions"
  "os"
)

func main() {
  if len(os.Args) != 2 {
		fmt.Println("Usage: hello <name>")
		return
	}

	name := os.Args[1]
  message := greetings.Hello(name)
  fmt.Println(message);
}