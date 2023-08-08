package domain

import (
	"com.github/salpreh/go-mars-rover/pkg/domain"
	"reflect"
	"testing"
)

func Test_ShouldCreateNewRover(t *testing.T) {
	// given
	expected := domain.Rover{
		Position:  domain.Coordinate{1, 2},
		Direction: domain.North,
		MarsMap:   *domain.NewMap(10, 10),
	}

	// when
	rover := domain.NewRover(1, 2, domain.North)

	// then
	if !reflect.DeepEqual(*rover, expected) {
		t.Errorf("Constructor object is not as expected: E %+v | V %+v", expected, *rover)
	}
}

func Test_ShouldProcessForwardCommandY(t *testing.T) {
	// given
	rover := createMarsRover(domain.North)

	// when
	err := rover.ProcessCommand("f")

	// then
	if err != nil || rover.Position.X != 1 || rover.Position.Y != 3 {
		t.Errorf("Mars rover not moved correctly: %+v", rover)
	}
}

func Test_ShouldProcessForwardCommandX(t *testing.T) {
	// given
	rover := createMarsRover(domain.East)

	// when
	err := rover.ProcessCommand("f")

	// then
	if err != nil || rover.Position.X != 2 || rover.Position.Y != 2 {
		t.Errorf("Mars rover not moved correctly: %+v", rover)
	}
}

func Test_ShouldProcessBackwardCommandY(t *testing.T) {
	// given
	rover := createMarsRover(domain.North)

	// when
	err := rover.ProcessCommand("b")

	// then
	if err != nil || rover.Position.X != 1 || rover.Position.Y != 1 {
		t.Errorf("Mars rover not moved correctly: %+v", rover)
	}
}

func Test_ShouldProcessBackwardCommandX(t *testing.T) {
	// given
	rover := createMarsRover(domain.East)

	// when
	err := rover.ProcessCommand("b")

	// then
	if err != nil || rover.Position.X != 0 || rover.Position.Y != 2 {
		t.Errorf("Mars rover not moved correctly: %+v", rover)
	}
}

func Test_ShouldReturnErrorWhenUnknownCommand(t *testing.T) {
	// given
	rover := createMarsRover(domain.North)

	// when
	err := rover.ProcessCommand("g")

	// then
	if err == nil {
		t.Errorf("Expected error")
	}
}

func Test_ShouldProcessMultipleCommands(t *testing.T) {
	// given
	rover := createMarsRover(domain.North)
	commands := []string{"f", "f", "b", "f"}

	// when
	err := rover.ProcessCommands(commands)

	// then
	if err != nil || rover.Position.X != 1 || rover.Position.Y != 4 {
		t.Errorf("Mars rover not moved correctly: %+v", rover)
	}
}

func Test_ShouldStopProcessingCommandsWhenUnknownCommand(t *testing.T) {
	// given
	rover := createMarsRover(domain.North)
	commands := []string{"f", "f", "g", "f"}

	// when
	err := rover.ProcessCommands(commands)

	// then
	if err == nil {
		t.Errorf("Expected error")
	}
}

func Test_ShouldTurnRight(t *testing.T) {
	// given
	rover := createMarsRover(domain.North)

	// when
	err := rover.ProcessCommand("r")

	// then
	if err != nil || rover.Direction != domain.East {
		t.Errorf("Mars rover not turned correctly: %+v", rover)
	}
}

func Test_ShouldTurnLeft(t *testing.T) {
	// given
	rover := createMarsRover(domain.North)

	// when
	err := rover.ProcessCommand("l")

	// then
	if err != nil || rover.Direction != domain.West {
		t.Errorf("Mars rover not turned correctly: %+v", rover)
	}
}

func Test_ShouldWrapPositionAtXEdge(t *testing.T) {
	// given
	rover := createMarsRover(domain.East)

	// when
	rover.ProcessCommands([]string{"b", "b", "b"})

	// then
	if rover.Position.X != 3 {
		t.Errorf("Mars rover position not wrapped correctly: %+v", rover)
	}
}

func Test_ShouldWrapPositionAtYEdge(t *testing.T) {
	// given
	rover := createMarsRover(domain.North)

	// when
	rover.ProcessCommands([]string{"f", "f", "f"})

	// then
	if rover.Position.Y != 0 {
		t.Errorf("Mars rover position not wrapped correctly: %+v", rover)
	}
}

func Test_ShouldReturnErrorWhenMoveForwardToObstacle(t *testing.T) {
	// given
	rover := createMarsRoverWithObstaclesMap(domain.West)

	// when
	err := rover.ProcessCommand("f")

	// then
	if err == nil {
		t.Errorf("Expected obstacle error")
	}
}

func Test_ShouldReturnErrorWhenMoveBackwardToObstacle(t *testing.T) {
	// given
	rover := createMarsRoverWithObstaclesMap(domain.South)

	// when
	err := rover.ProcessCommand("b")

	// then
	if err == nil {
		t.Errorf("Expected obstacle error")
	}
}

func Test_ShouldReturnErrorWhenProcessingCommandsAndFindsObstacle(t *testing.T) {
	// given
	rover := createMarsRoverWithObstaclesMap(domain.North)
	commands := []string{"r", "f", "f", "l", "f", "f", "l", "f"}

	// when
	err := rover.ProcessCommands(commands)

	// then
	if err == nil {
		t.Errorf("Expected error")
	}
}

func createMarsRover(direction domain.Direction) *domain.Rover {
	return &domain.Rover{
		Position:  domain.Coordinate{1, 2},
		Direction: direction,
		MarsMap:   *domain.NewMap(5, 5),
	}
}

func createMarsRoverWithObstaclesMap(direction domain.Direction) *domain.Rover {
	obstacles := []domain.Coordinate{
		{0, 2},
		{1, 3},
		{3, 4},
	}

	return &domain.Rover{
		Position:  domain.Coordinate{1, 2},
		Direction: direction,
		MarsMap:   *domain.NewMapWithObstacles(5, 5, obstacles),
	}
}
