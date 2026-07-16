package printer

import "fmt"

func PrintValues(dig int, pass string) {
	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Printf("Length   : %d\n", dig)
	fmt.Printf("Password : %s\n", pass)
	fmt.Println("----------------------------------------")
	fmt.Println()
}
