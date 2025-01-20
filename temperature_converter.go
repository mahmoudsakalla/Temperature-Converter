package main

import "fmt"

func main() {
	var choice int
	var temp float64

	for {
		//selection menu, asks the user from which unit they want to convert from
		fmt.Println("Temperature Converter")
		fmt.Println("1: Celsius to Fahrenheit")
		fmt.Println("2: Fahrenheit to Celsius")
		fmt.Print("Choose an option (1 or 2): ")
		fmt.Scanln(&choice)

		// The formulas to convert from celcius to fahrenheit and vice versa are standard
		//the formulas can be found via google search, they are simple
		// loop to process the input of the user and ask for the temperature they wanna convert
		if choice == 1 {
			fmt.Print("Enter temperature in Celsius: ")
			fmt.Scanln(&temp)
			fahrenheit := (temp * 9 / 5) + 32
			fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", temp, fahrenheit)
			break
		} else if choice == 2 {
			fmt.Print("Enter temperature in Fahrenheit: ")
			fmt.Scanln(&temp)
			celsius := (temp - 32) * 5 / 9
			fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", temp, celsius)
			break
		} else {
			// make sure that the choice is valid, between 1 and 2 only
			fmt.Println("Invalid choice. please select 1 or 2 only.")
			fmt.Println("Try again:")
		}
	}

}
