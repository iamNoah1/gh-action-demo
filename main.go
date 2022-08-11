package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Some input vat:", os.Getenv("INPUT_API_KEY"))
	fmt.Println("::set-output name=output_var::someOutputValue")
}
