package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: likscan MAIN SOURCE\n")
		return
	}

}
