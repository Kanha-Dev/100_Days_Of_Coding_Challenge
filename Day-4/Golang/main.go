package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Welcome the User
	fmt.Println("Welcome to the Rock Paper Scissor Game!!")
	//Initialise three variables with rock,paper and scissor images
	rock := `
    _______
---'   ____)
      (_____)
      (_____)
      (____)
---.__(___)
`
	paper := `
	    _______
	---'   ____)____
			  ______)
			  _______)
			 _______)
	---.__________)
	`
	scissors := `
	    _______
	---'   ____)____
			  ______)
		   __________)
		  (____)
	---.__(___)
	`
	//Inirialise choice variable with integer type
	var choice int
	//Displat rock,paper and scissors
	fmt.Println(rock)
	fmt.Println(paper)
	fmt.Println(scissors)
	//Prompt the user for choice
	fmt.Println("What do you choose? Type 0 for Rock, 1 for Paper or 2 for Scissors.")
	_, _ = fmt.Scanf("%d", &choice)
	//Generate random number between the range 0 to 2
	computer := rand.Intn(2)
	//Display the user's choice
	fmt.Println("The user Chooses:")
	if choice == 0 {
		fmt.Println(rock)
	} else if choice == 1 {
		fmt.Println(paper)
	} else {
		fmt.Println(scissors)
	}
	//Display Computer choice
	fmt.Println("The Computer Chooses:")
	if computer == 0 {
		fmt.Println(rock)
	} else if computer == 1 {
		fmt.Println(paper)
	} else {
		fmt.Println(scissors)
	}
	//Build if else ladder
	if choice > 0 && computer == 2 {
		fmt.Println("You win!")
	} else if choice == 2 && computer > 0 {
		fmt.Println("You lose")
	} else if choice > computer {
		fmt.Println("You win!")
	} else if computer > choice {
		fmt.Println("You lose")
	} else if computer == choice {
		fmt.Println("Its a draw")
	} else {
		fmt.Println("Invalid choice")
	}

}
