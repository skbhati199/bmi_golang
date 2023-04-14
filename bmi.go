package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	// Output information
	fmt.Println("Welcome to BMI Calculator")
	fmt.Println("___________________________")
	// Prompt for user input ( weight  + weight)
	fmt.Print("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')

	fmt.Print("Please enter your height (meters): ")
	heightInput, _ := reader.ReadString('\n')
	// Save that user input in variables

	weightInput = strings.Replace(weightInput, "\r", "", -1)
	heightInput = strings.Replace(heightInput, "\r", "", -1)
	// fmt.Print(weightInput)
	// fmt.Print(heightInput)
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)

	// Convert user input to float
	weight, err := strconv.ParseFloat(weightInput, 64)
	if err != nil {
		fmt.Println(err)
	}
	height, err := strconv.ParseFloat(heightInput, 64)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(weight)
	// fmt.Println(height)
	// Calculate the BMI weight/(weight + weight)
	bmi := float64(weight / (height * height))

	// Output the calculated BMI
	fmt.Printf("BMI: %f\n", bmi)
}
