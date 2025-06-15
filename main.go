package main

import "C"
import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

var upperCasePool = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lowerCasePool = "abcdefghijklmnopqrstuvwxyz"
var digitsPool = "0123456789"
var symbolsPool = "!@#$%^&*()-_=+[]{}|;:,.<>?/~`"

func main() {
	fmt.Println("Password generator started in GoLang...")

	var passwordLength int
	var upperCase bool
	var lowerCase bool
	var digits bool
	var symbols bool
	var help bool

	flag.IntVar(&passwordLength, "passwordLength", 12, "length of final password")
	flag.BoolVar(&upperCase, "upper", true, "should include uppercase")
	flag.BoolVar(&lowerCase, "lower", true, "should include lowercase")
	flag.BoolVar(&digits, "digits", true, "should include digits")
	flag.BoolVar(&symbols, "symbols", false, "should include symbols")
	flag.BoolVar(&help, "help", false, "should include instructions")

	flag.Parse()

	pool := generatePassword(passwordLength, upperCase, lowerCase, digits, symbols)

	fmt.Println(passwordLength)

	if len(pool) <= 0 {
		fmt.Println("Please select at least one option")
		return
	}

	finalPassword := buildPasswordPool(passwordLength, pool)

	generateHelp(help)

	fmt.Println(finalPassword)
}

func generatePassword(passwordLength int, upperCase, lowerCase, digits, symbols bool) string {
	var passwordPool string

	if passwordLength < 10 {
		err := fmt.Errorf("The minimum password length is  %d", passwordLength)
		fmt.Println(err.Error())
		return ""
	}

	if upperCase {
		passwordPool += upperCasePool
	}
	if lowerCase {
		passwordPool += lowerCasePool
	}
	if digits {
		passwordPool += digitsPool
	}
	if symbols {
		passwordPool += symbolsPool
	}

	if passwordLength > 10 {
		fmt.Println("This password length might be too long to remember")
	}

	return passwordPool
}

func buildPasswordPool(passwordLength int, pool string) string {
	password := make([]rune, passwordLength)
	maxValue := big.NewInt(int64(len(pool)))
	for char := 0; char < passwordLength; char++ {
		randomNumber, err := rand.Int(rand.Reader, maxValue)
		if err != nil {
			fmt.Println("Error:", err)
			return ""
		}
		randomIndex := int(randomNumber.Int64())

		password[char] = rune(pool[randomIndex])
	}

	finalPassword := string(password)

	return finalPassword
}

func generateHelp(help bool) {
	if help {
		fmt.Println("See README.md for more information on how to use this program.")
		fmt.Println("👇🏾👇🏾============================????============================👇🏾👇🏾")
	}
}
