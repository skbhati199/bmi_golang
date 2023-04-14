package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/skbhati199/bmi/info"
)



func main() {

	// Output information
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separatpr)
	// Prompt for user input ( weight  + weight)
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
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
