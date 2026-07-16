package passgenerator

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/LeonardoCG12/ForgeKey/variables"
)

func GeneratePass(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("Password length must be greater than zero.")
	}

	res := make([]byte, n)
	maxChars := big.NewInt(int64(len(variables.Chars)))

	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, maxChars)
		if err != nil {
			return "", fmt.Errorf("Failed to read from crypto/rand: %w.", err)
		}

		res[i] = variables.Chars[num.Int64()]
	}

	return string(res), nil
}
