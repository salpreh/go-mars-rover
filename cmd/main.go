package main

import (
	"com.github/salpreh/go-mars-rover/pkg/domain"
	"fmt"
)

func main() {
	rover := createRover()
	commands := []string{"f", "f", "r", "b", "r", "f", "l", "l", "l", "f"} // Ends in -2, 1

	fmt.Printf("Current rover status: %+v\n", *rover)

	rover.ProcessCommands(commands)

	fmt.Printf("Rover status after commands: %+v\n", *rover)
}

func createRover() *domain.Rover {
	return domain.NewRover(0, 0, domain.North)
}
