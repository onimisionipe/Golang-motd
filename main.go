package main 

import (
         "fmt"
         "motd/message"
)

func main() {
	message := message.Greeting("Kelvin", "hello")
	fmt.Println(message)
}


 