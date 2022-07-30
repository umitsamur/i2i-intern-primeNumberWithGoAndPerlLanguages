package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isPrime(number int) (message string) {
	var isPrime bool
	isPrime = true
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			isPrime = false
		}
	}
	if isPrime == true {
		message = "Number is prime number"
	} else {
		message = "Number is not prime number"
	}
	return message
}

func check(err error) {
	if err != nil {
		log.Fatalf("unable to file: %v", err)
	}
}

func stringToInt(text string) (number int) {
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Casting problem. Please enter integer data")
	}
	return number
}

func intToString(number int) (text string) {
	text = strconv.Itoa(number)
	return text
}

func main() {
	var number int

	txtWrite, err := os.Create("result.txt")
	check(err)
	txtFile, err := os.Open("a.txt")
	check(err)

	scanner := bufio.NewScanner(txtFile)
	for scanner.Scan() {
		number = stringToInt(scanner.Text())
		stringTypeOfNumberVariable := intToString(number)
		if number < 2 {
			fmt.Println("Number must be greater or equal than 2")
		} else {
			txtWrite.WriteString(stringTypeOfNumberVariable + " => " + isPrime(number) + "\n")
			fmt.Printf("%d => %s \n", number, isPrime(number))
		}
	}
	txtFile.Close()
	txtWrite.Close()
}
