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
  fmt.Print(">>> ")

  scanner.Scan()

  text := scanner.Text()

  cleanedInput := clearInput(text)
   
  fmt.Println("Echoing: ", cleanedInput)
  
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
