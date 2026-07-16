package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/LeonardoCG12/ForgeKey/utils/passgenerator"
	"github.com/LeonardoCG12/ForgeKey/utils/printer"
)

func main() {
	var dig int

	if len(os.Args) > 1 {
		parsedDig, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Conversion error. Length must be a valid number: %v.", err)
		}
		dig = parsedDig
	} else {
		fmt.Print("Enter password length: ")
		_, err := fmt.Scanf("%d", &dig)
		if err != nil {
			log.Fatalf("Error reading input length: %v.", err)
		}
	}

	pass, err := passgenerator.GeneratePass(dig)
	if err != nil {
		log.Fatalf("Error generating password: %v.", err)
	}

	printer.PrintValues(dig, pass)
}
