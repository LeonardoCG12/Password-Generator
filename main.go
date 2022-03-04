package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
)

func init() {
	checkSafety()
}

func checkSafety() {
	buf := make([]byte, 1)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		panic(fmt.Sprintf("It is not safe to generate a password beacause crypto/rand is not available\n Error: %#v", err))
	}
}

func generatePass(n int64) (string, error) {
	var i int64
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%&*()-_=+-,./|\\{}[]^~;:?<>\"'"
	res := make([]byte, n)

	for i = 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))

		if err != nil {
			return "", err
		}

		res[i] = chars[num.Int64()]
	}

	return string(res), nil
}

func printValues(dig int64, pass string, err error) {

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\033[1;32mDigits:\033[1;36m %v", dig)
	fmt.Printf("\n\033[1;32mPassword:\033[1;36m %v\033[0m\n\n", pass)
}

func main() {
	argLen := len(os.Args)

	if argLen > 1 {
		dig, err := strconv.ParseInt(os.Args[1], 10, 64)

		if err != nil {
			panic(fmt.Sprintf("%#v", err))
		}

		pass, err := generatePass(int64(dig))
		printValues(dig, pass, err)

	} else {
		var dig int64

		fmt.Printf("\n\033[1;33mLength:\033[1;36m ")
		fmt.Scanf("%d", &dig)
		pass, err := generatePass(int64(dig))
		printValues(dig, pass, err)
	}
}
