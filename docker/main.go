package main

import (
	"fmt"
	"os"
)

// To terminate the daemon use:
//  kill `cat pid`

func main() {
	fmt.Println("Pull request 10")
	if (len(os.Args) > 1) && (os.Args[1] == "tests") {
                fmt.Println("Run tests...")
		file, err := os.Create("result.txt")
		if err != nil {
			fmt.Println("Cannot create file", err)
			os.Exit(10)
		}
		defer file.Close()
		fmt.Println("Results OK")
		fmt.Fprintf(file, "Results OK")
//		os.Exit(1)
	}
	fmt.Println("goodbye")
}
