package main

import (
	"fmt"
	"log"
	"os"
)

// To terminate the daemon use:
//  kill `cat pid`

func main() {
	fmt.Println("Helo Pull request 10")
	if (len(os.Args) > 1) && (os.Args[1] == "tests") {
                fmt.Println("Run tests...")
		file, err := os.Create("result.txt")
		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		defer file.Close()

		fmt.Fprintf(file, "Results OK")
		// os.Exit(1)
	}
	fmt.Println("goodbye cruel world")
}
