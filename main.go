package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main() {

  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Basic Pokedex-cli [press exit to terminate]")
  


  for {

    fmt.Print(">>> ")

    scanner.Scan()

    text := scanner.Text()

    cleanedInput := clearInput(text)[0]

    if len(cleanedInput) == 0 {
      continue
    }

    switch cleanedInput {
    case "exit":
       fmt.Println("Teminating the program")

       os.Exit(0)
    
    case "help":
       fmt.Println("Echoing: ", cleanedInput[0])
    default: 
        fmt.Println("Invalid input")
    }
    


  }
  
}



func clearInput(str string) []string {
  str = strings.TrimSpace(str)
  str = strings.ToLower(str)
  words := strings.Fields(str)

  return words
}


// get user input
// validate input
// make a REPL
