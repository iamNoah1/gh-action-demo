package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Some input var:", os.Getenv("INPUT_API_TOKEN"))
	fmt.Print("::set-output name=output_var::someOutputValue")
}
