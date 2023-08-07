package domain

import (
	"errors"
	"log"
)

type Rover struct {
	Position  Coordinate
	Direction Direction
	MarsMap   Map
}

func NewRover(xPosition int, yPosition int, direction Direction) *Rover {
	return &Rover{
		Position:  Coordinate{xPosition, yPosition},
		Direction: direction,
		MarsMap:   Map{10, 10},
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
	rov.adjustPosition()
}

func (rov *Rover) moveBackward() {
	xDirection, yDirection := getDirectionMovementVector(rov.Direction)
	rov.Position.X += xDirection * -1
	rov.Position.Y += yDirection * -1
	rov.adjustPosition()
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

// adjustPosition Adjusts current rover position checking Mars map.
// If limits of the map has been exceeded wraps position to other end (e.g. for 5 by 5 map if x position is 6 it will be corrected to 1)
// Position is 0 indexed, so for a map of 5 width last position is 4.
func (rov *Rover) adjustPosition() {
	for rov.Position.X >= rov.MarsMap.Width {
		rov.Position.X -= rov.MarsMap.Width
	}

	for rov.Position.X < 0 {
		rov.Position.X += rov.MarsMap.Width
	}

	for rov.Position.Y >= rov.MarsMap.Height {
		rov.Position.Y -= rov.MarsMap.Height
	}

	for rov.Position.Y < 0 {
		rov.Position.Y += rov.MarsMap.Height
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
