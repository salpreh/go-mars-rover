package domain

import (
	"errors"
	"fmt"
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
		MarsMap:   *NewMap(10, 10),
	}
}

func (rov *Rover) ProcessCommands(commands []string) error {
	for i, command := range commands {
		err := rov.ProcessCommand(command)
		if err != nil {
			log.Printf("Unable to process command no. %d. Skipping following commands. ERR: %s\n", i, err)

			return errors.Join(
				errors.New(fmt.Sprintf("Stopped processing at command no. %d", i)),
				err,
			)
		}
	}

	return nil
}

func (rov *Rover) ProcessCommand(command string) error {
	var err error
	switch command {
	case Forward:
		err = rov.moveForward()
	case Backward:
		err = rov.moveBackward()
	case TurnRight:
		rov.turnRight()
	case TurnLeft:
		rov.turnLeft()
	default:
		err = errors.New("Unknown command: " + command)
	}

	return err
}

func (rov *Rover) moveForward() error {
	xDirection, yDirection := getMovementDirectionVector(rov.Direction)
	newPosition := Coordinate{
		X: rov.Position.X + xDirection,
		Y: rov.Position.Y + yDirection,
	}
	rov.adjustPosition(&newPosition)

	return rov.updatePosition(&newPosition)
}

func (rov *Rover) moveBackward() error {
	xDirection, yDirection := getMovementDirectionVector(rov.Direction)
	newPosition := Coordinate{
		X: rov.Position.X + xDirection*-1,
		Y: rov.Position.Y + yDirection*-1,
	}
	rov.adjustPosition(&newPosition)

	return rov.updatePosition(&newPosition)
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

// adjustPosition Adjusts provided coordinate checking Mars map.
// If limits of the map has been exceeded wraps position to other end (e.g. for 5 by 5 map if x position is 6 it will be corrected to 1)
// Coordinate is 0 indexed, so for a map of 5 width last position is 4.
func (rov *Rover) adjustPosition(coord *Coordinate) {
	for coord.X >= rov.MarsMap.Width() {
		coord.X -= rov.MarsMap.Width()
	}

	for coord.X < 0 {
		coord.X += rov.MarsMap.Width()
	}

	for coord.Y >= rov.MarsMap.Height() {
		coord.Y -= rov.MarsMap.Height()
	}

	for coord.Y < 0 {
		coord.Y += rov.MarsMap.Height()
	}
}

func (rov *Rover) updatePosition(newPosition *Coordinate) error {
	if rov.MarsMap.HasObstacle(*newPosition) {
		return errors.New(fmt.Sprintf("Unable to move to position, obstacle in coordinates %+v", newPosition))
	}

	rov.Position = *newPosition

	return nil
}

func getMovementDirectionVector(direction Direction) (int, int) {
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
