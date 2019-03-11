package main

import (
        "bufio"
        "fmt"
        "motd/message"
        "os"
        "strings"
)

func main() {
        f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)

        if err != nil {
            fmt.Println("Error: unable to open /etc/motd")
            os.Exit(1)
        }
        // the above opens the file first and checks if the user can write to it 

        defer f.close()
        // the above close the file after it has been opened 

        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Your Greeting: ")
        phrase, _ := reader.ReadString('\n')
        phrase = strings.TrimSpace(phrase)

        fmt.Print("Your Name: ")
        name, _ := reader.ReadString('\n')
        name = strings.TrimSpace(name)

        // if the above works the user will be prompted for some information 

        m := message.Greeting(name, phrase)
        _, err = f.Write([]byte(m))
       
         // then the program will write to the file  
        if err != nil {
           fmt.Print("unable to write /etc/motd") 
           os.Exit(1)    
        }
}
        // then we would check to see if it succeeded 