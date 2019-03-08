package message

import (
         "fmt"
)

func Greeting(name, message string) string {
  return fmt.Sprintf("%s, %s", message, name)
  }