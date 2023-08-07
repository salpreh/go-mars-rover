package domain

import (
	"errors"
	"log"
)

type Rover struct {
	Position  Coordinate
	Direction Direction
}

func NewRover(xPosition int, yPosition int, direction Direction) *Rover {
	return &Rover{
		Position:  Coordinate{xPosition, yPosition},
		Direction: direction,
	}
}

func (rov *Rover) ProcessCommands(commands []string) {
	for _, command := range commands {
		err := rov.ProcessCommand(command)
		if err != nil {
			log.Printf("Unknown command skipped: %s\n", err)
		}
	}
}

func (rov *Rover) ProcessCommand(command string) error {
	switch command {
	case Forward:
		rov.moveForward()
	case Backward:
		rov.moveBackward()
	case TurnRight:
		rov.turnRight()
	case TurnLeft:
		rov.turnLeft()
	default:
		return errors.New("Unknown command: " + command)
	}

	return nil
}

func (rov *Rover) moveForward() {
	xDirection, yDirection := getDirectionMovementVector(rov.Direction)
	rov.Position.X += xDirection
	rov.Position.Y += yDirection
}

func (rov *Rover) moveBackward() {
	xDirection, yDirection := getDirectionMovementVector(rov.Direction)
	rov.Position.X += xDirection * -1
	rov.Position.Y += yDirection * -1
}

func (rov *Rover) turnRight() {
	rov.Direction = (rov.Direction + 1) % 4
}

func (rov *Rover) turnLeft() {
	if rov.Direction == 0 {
		rov.Direction = 3
	} else {
		rov.Direction = (rov.Direction - 1) % 4
	}
}

func getDirectionMovementVector(direction Direction) (int, int) {
	switch direction {
	case North:
		return 0, 1
	case South:
		return 0, -1
	case East:
		return 1, 0
	case West:
		return -1, 0
	default:
		return 0, 0
	}
}

type Coordinate struct {
	X int
	Y int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

const (
	Forward   = "f"
	Backward  = "b"
	TurnLeft  = "l"
	TurnRight = "r"
)
