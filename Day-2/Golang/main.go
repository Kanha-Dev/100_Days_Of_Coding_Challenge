package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func roundOff(number float64, decimal int) float64 {
	shift := math.Pow(10, float64(decimal))
	return math.Round(number*shift) / shift
}
func main() {
	fmt.Scanln()
	fmt.Println("Welcome to the Tip Calculator.")
	//Prompt the User for the Total Bill
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Printf("What was the total bill:$ ")
	bill, _ := reader1.ReadString('\n')
	num_bill, _ := strconv.ParseFloat(strings.TrimSpace(bill), 64)
	//Prompt the User for Percentage of Tip
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Printf("What Percentage tip would you like to give? 10,12 or 15? ")
	percentage, _ := reader2.ReadString('\n')
	num_percentage, _ := strconv.ParseFloat(strings.TrimSpace(percentage), 64)
	//Prompt the User for Number of People
	reader3 := bufio.NewReader(os.Stdin)
	fmt.Printf("How many people to spilt the bill? ")
	people, _ := reader3.ReadString('\n')
	num_people, _ := strconv.ParseFloat(strings.TrimSpace(people), 64)
	//Calculate the Final Bill
	var final_bill float64 = (num_bill / num_people) * (1 + (num_percentage / 100))
	final := roundOff(final_bill, 3)
	fmt.Println(final)

}
