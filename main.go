package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			line := scanner.Text()
			cleaned := cleanInput(line)
			if len(cleaned) > 0 {
				fmt.Printf("Your command was: %s\n", cleaned[0])
			}
		}
	}
}
