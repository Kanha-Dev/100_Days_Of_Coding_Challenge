package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Welcome Message
	welcome := "This is a Band Name Generator"
	fmt.Println(welcome)
	//Prompt User for City Name
	fmt.Printf("Enter your City Name: ")
	city := bufio.NewReader(os.Stdin)
	cityinput, _ := city.ReadString('\n')
	//Prompt User for Pet Name
	fmt.Printf("Enter your Pet's Name: ")
	pet := bufio.NewReader(os.Stdin)
	petinput, _ := pet.ReadString('\n')
	//Adding the two lines
	BandName := strings.TrimSpace(cityinput) + " " + strings.TrimSpace(petinput)
	//Printing the Band Name
	fmt.Println("This is the name of your Band: ", BandName)

}
