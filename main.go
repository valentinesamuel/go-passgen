package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"net/http"
)

var upperCasePool = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lowerCasePool = "abcdefghijklmnopqrstuvwxyz"
var digitsPool = "0123456789"
var symbolsPool = "!@#$%^&*()-_=+[]{}|;:,.<>?/~`"

type ResponseData struct {
	Data     any    `json:"data"`
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

func main() {
	var cliMode bool

	flag.BoolVar(&cliMode, "cli", true, "starts the program in cliMode")
	flag.Parse()

	fmt.Println("CLI Mode is", cliMode)
	if !cliMode {
		serverModule()
	} else {
		cliModule()
	}
}

func serverModule() {
	// Create a new ServeMux to register routes
	mux := http.NewServeMux()

	// Register routes on the new mux
	mux.HandleFunc("/generate", HandleGeneratePassword)

	// Create handler chain with both middlewares
	// The order is important - RequestLogger will run first, then HeaderLogger
	handler := RequestLogger(HeaderLogger(mux))

	fmt.Println("Started server .....")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting the server")
		return
	}
}

func cliModule() {
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

	pool, _ := generatePassword(passwordLength, upperCase, lowerCase, digits, symbols)

	fmt.Println(passwordLength)

	if len(pool) <= 0 {
		fmt.Println("Please select at least one option")
		return
	}

	finalPassword, _ := buildPasswordPool(passwordLength, pool)

	generateHelp(help)

	fmt.Println(finalPassword)
}

func generatePassword(passwordLength int, upperCase, lowerCase, digits, symbols bool) (string, error) {
	var passwordPool string

	if passwordLength < 10 {
		err := fmt.Errorf("the minimum password length is %d", 10)
		return "", err
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

	return passwordPool, nil
}

func buildPasswordPool(passwordLength int, pool string) (string, error) {
	password := make([]rune, passwordLength)
	maxValue := big.NewInt(int64(len(pool)))
	for char := 0; char < passwordLength; char++ {
		randomNumber, err := rand.Int(rand.Reader, maxValue)
		if err != nil {
			fmt.Println("Error:", err)
			return "", err
		}
		randomIndex := int(randomNumber.Int64())

		password[char] = rune(pool[randomIndex])
	}

	finalPassword := string(password)

	return finalPassword, nil
}

func generateHelp(help bool) {
	if help {
		fmt.Println("See README.md for more information on how to use this program.")
		fmt.Println("👇🏾👇🏾============================????============================👇🏾👇🏾")
	}
}
