package main

import (
	"fmt"
	"strings"
)

func main() {
	//Welcome the User
	fmt.Println("Welcome to Treasure Island.")
	//Tell the user what his mission is
	fmt.Println("Your mission is to find the treasure.")
	var choice string
	//Give the user his first choice, also lower case the input so that it can match with the if statement
	fmt.Println("You're at a crossroad. Where do you want to go? Type 'left' or 'right'")
	_, _ = fmt.Scanf("%s", &choice)
	choice = strings.ToLower(choice)
	//#Use "if" command for making a choice
	if choice == "left" {
		fmt.Println("You've come to a lake. There is an island in the middle of the lake. Type 'wait' to wait for a boat. Type 'swim' to swim across.")
		var choice2 string
		_, _ = fmt.Scanf("%s", &choice2)
		choice2 = strings.ToLower(choice2)
		//First Nested Choice
		if choice2 == "swim" {
			fmt.Println("You arrive at the island unharmed. There is a house with 3 doors. One red, one yellow, and one blue. Which color do you choose?")
			var choice3 string
			_, _ = fmt.Scanf("%s", &choice3)
			choice3 = strings.ToLower(choice3)
			//Now ask the user for the last nested three choices
			if choice3 == "red" {
				fmt.Println("It's a room full of fire. Game Over.")
			} else if choice3 == "yellow" {
				fmt.Println("You found the treasure! You Win!")
			} else if choice3 == "blue" {
				fmt.Println("You enter a room of beasts. Game Over.")
			} else {
				fmt.Println("You chose a door that doesn't exist. Game Over.")
			}
		} else {
			//This is the else choice for the First nested "if" statement
			fmt.Println("You got attacked by an angry trout. Game Over.")
		}
	} else {
		//This is the else choice for the Parent "if" statement
		fmt.Println("You fell into a hole. Game Over.")
	}
}
