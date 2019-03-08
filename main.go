package main 

import (
         "fmt"
)

func main() {
	message := greeting("Kelvin", "hello")
	fmt.Println(message)
}

func greeting(name string, message string) string {
  return fmt.Sprintf("%s, %s", message, name)
  }
