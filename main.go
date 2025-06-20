package main

import (
	"errors"
	customerror "go_lang/customError"
	"go_lang/users"
	"log"
	"strconv"
)

func main() {
	log.Println("hello world")
	result, err := divide(1, 0)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println(result)
	result, err = divide(40, 19)
	log.Println("Result of division:", result)
	newUser := users.NewUser("John Doe", "tW3oJ@example.com", 30, "123 Main St", true)
	capitalizedName := newUser.CapatalizeUserName()
	log.Println("Capitalized User Name:", capitalizedName)

	// parsing age
	age, cerr := parseAge("pra")
	if cerr.Message != "" {
		log.Println ("Custom Error:", cerr.Message, "Type:", cerr.Type, "Error:", cerr.Err)
	}
		age,_ = parseAge("25")
		log.Println("Parsed Age:", age)
	
}

func divide(a,b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func parseAge (age string) (*int, customerror.CustomError) {
	convertedAge, err := strconv.Atoi(age)	
	if err != nil {
		return nil, customerror.CustomError{
			Message: "Failed to convert age",
			Err: err,
			Type: "ConversionError",
		}
	}
	if convertedAge < 1|| convertedAge > 120 {
		return nil, customerror.CustomError{
			Message: "Age must be between 1 and 120",
			Err: errors.New("invalid age range"),
			Type: "RangeError",
		}
	}
	return &convertedAge, customerror.CustomError{}


}