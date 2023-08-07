package domain

import (
	"com.github/salpreh/go-mars-rover/pkg/domain"
	"testing"
)

func Test_ShouldCreateNewRover(t *testing.T) {
	// given
	expected := domain.Rover{
		Position:  domain.Coordinate{1, 2},
		Direction: domain.North,
	}

	// when
	rover := domain.NewRover(1, 2, domain.North)

	// then
	if *rover != expected {
		t.Errorf("Constructor object is not as expected: E %+v | V %+v", expected, *rover)
	}
}

func Test_ShouldProcessForwardCommandY(t *testing.T) {
	// given
	rover := domain.NewRover(1, 2, domain.North)

	// when
	err := rover.ProcessCommand("f")

	// then
	if err != nil || rover.Position.X != 1 || rover.Position.Y != 3 {
		t.Errorf("Mars rover not moved correctly: %+v", rover)
	}
}

func Test_ShouldProcessForwardCommandX(t *testing.T) {
	// given
	rover := domain.NewRover(1, 2, domain.East)

	// when
	err := rover.ProcessCommand("f")

	// then
	if err != nil || rover.Position.X != 2 || rover.Position.Y != 2 {
		t.Errorf("Mars rover not moved correctly: %+v", rover)
	}
}

func Test_ShouldProcessBackwardCommandY(t *testing.T) {
	// given
	rover := domain.NewRover(1, 2, domain.North)

	// when
	err := rover.ProcessCommand("b")

	// then
	if err != nil || rover.Position.X != 1 || rover.Position.Y != 1 {
		t.Errorf("Mars rover not moved correctly: %+v", rover)
	}
}

func Test_ShouldProcessBackwardCommandX(t *testing.T) {
	// given
	rover := domain.NewRover(1, 2, domain.East)

	// when
	err := rover.ProcessCommand("b")

	// then
	if err != nil || rover.Position.X != 0 || rover.Position.Y != 2 {
		t.Errorf("Mars rover not moved correctly: %+v", rover)
	}
}

func Test_ShouldReturnErrorWhenUnknownCommand(t *testing.T) {
	// given
	rover := domain.NewRover(1, 2, domain.North)

	// when
	err := rover.ProcessCommand("g")

	// then
	if err == nil {
		t.Errorf("Expected error")
	}
}

func Test_ShouldProcessMultipleCommands(t *testing.T) {
	// given
	rover := domain.NewRover(1, 2, domain.North)
	commands := []string{"f", "f", "b", "f"}

	// when
	rover.ProcessCommands(commands)

	// then
	if rover.Position.X != 1 || rover.Position.Y != 4 {
		t.Errorf("Mars rover not moved correctly: %+v", rover)
	}
}
