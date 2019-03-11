package main

import (
        "bufio"
        "fmt"
        "motd/message"
        "os"
        "strings"
        "flag"
)

func main() {
  // Define Flags 
  var name string
  var greeting string 
  var prompt bool
  var preview bool
  
  // Parse Flags 
  flag.Stringvar(&name, "name", "", "name to use within the message")
  flag.Stringvar(&greeting, "greeting", "", "phrase to be used within the message")
  flag.BoolVar(&prompt, "prompt", false, "use prompt to input name and greeting")
  flag.BoolVar(&preview, "preview", false, "use preview to output message without writing to /etc/motd")
  

  flag.Parse()

   // show usage if flags are invalid 
   if prompt == false && (name == "" || greeting == "") {
     flag.Usage()
     os.Exit
   }

   // optionally print flags and exit based on DEBUG env variable 

   if os.Getenv("DEBUG") != "" {
    fmt.Println("Name:", name)
    fmt.Println("Greeting:", greeting)
    fmt.Println("Prompt:", prompt)
    fmt.Println("Preview:", preview)

    os.Exit(0)

   // conditionally read from  stdin 
    if prompt {
      name, greeting = renderPrompt()
    }


  // Generate the message 
    m := message.Greeting(name, greeting)

  // Either preview the message or write it to file 
    if preview {
    	fmt.Println(m)
    } else {
      // Write content 

    f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)

        if err != nil {
            fmt.Println("Error: unable to open /etc/motd")
            os.Exit(1)

    defer f.close()

        
    _, err = f.Write([]byte(m))

       if err != nil {
           fmt.Print("unable to write /etc/motd") 
           os.Exit(1)    
    }

    }
}


func renderPrompt() (name, greeting string) {
  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Your Greeting: ")
  greeting, _ = reader.ReadString('\n')
  greeting = strings.TrimSpace(greeting)

  fmt.Print("Your Name: ")
  name, _ = reader.ReadString('\n')
  name = strings.TrimSpace(name)

  return
}

