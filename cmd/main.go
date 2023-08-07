package main

import (
	"com.github/salpreh/go-mars-rover/pkg/domain"
	"fmt"
)

func main() {
	rover := createRover()
	commands := []string{"f", "f", "f", "b", "f"}

	fmt.Printf("Current rover status: %+v\n", *rover)

	rover.ProcessCommands(commands)

	fmt.Printf("Rover status after commands: %+v\n", *rover)
}

func createRover() *domain.Rover {
	return domain.NewRover(0, 0, domain.North)
}
