package marsrover

import "testing"
import (
	"github.com/onsi/gomega"
)

func TestSetCoordinates(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	const coordinateX int = 6
	const coordinateY int = 7

	rover := new(Rover)

	//When
	rover.SetCoordinates(coordinateX, coordinateY)
	actualXCoordinate, actualYCoordinate := rover.Coordinates()

	//Then
	gomega.Expect(actualXCoordinate).To(gomega.Equal(coordinateX))
	gomega.Expect(actualYCoordinate).To(gomega.Equal(coordinateY))

}

func TestSetOrientation(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	const orientation string = "N"
	rover := new(Rover)

	//When
	rover.SetOrientation(orientation)
	actualOrientation := rover.Orientation()

	//Then
	gomega.Expect(actualOrientation).To(gomega.Equal(orientation))

}

func TestSetPlateau(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	plateau := Plateau{7, 8}
	rover := new(Rover)

	//When
	rover.SetPlateau(plateau)

	//then
	gomega.Expect(rover.Plateau).To(gomega.Equal(plateau))
}

func TestRoverSpinToTheLeftWhenStartingFromNorth(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	const orientation string = "N"
	rover := new(Rover)
	rover.SetOrientation(orientation)

	//When
	rover.Spin("L")
	orientationAfterSpinning := rover.Orientation()

	//Then
	gomega.Expect(orientationAfterSpinning).To(gomega.Equal("W"))

}

func TestRoverSpinToTheLeftWhenStartingFromWest(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	const orientation string = "W"
	rover := new(Rover)
	rover.SetOrientation(orientation)
	//When
	rover.Spin("L")
	orientationAfterSpinning := rover.Orientation()

	//Then
	gomega.Expect(orientationAfterSpinning).To(gomega.Equal("S"))

}

func TestRoverSpinToTheLeftWhenStartingFromSouth(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	const orientation string = "S"
	rover := new(Rover)
	rover.SetOrientation(orientation)

	//When
	rover.Spin("L")
	orientationAfterSpinning := rover.Orientation()

	//Then
	gomega.Expect(orientationAfterSpinning).To(gomega.Equal("E"))

}

func TestRoverSpinToTheLeftWhenStartingFromEast(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	const orientation string = "E"
	rover := new(Rover)
	rover.SetOrientation(orientation)

	//When
	rover.Spin("L")
	orientationAfterSpinning := rover.Orientation()

	//Then
	gomega.Expect(orientationAfterSpinning).To(gomega.Equal("N"))

}

func TestRoverSpinToTheRightWhenStartingFromNorth(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	const orientation string = "N"
	rover := new(Rover)
	rover.SetOrientation(orientation)

	//When
	rover.Spin("R")
	orientationAfterSpinning := rover.Orientation()

	//Then
	gomega.Expect(orientationAfterSpinning).To(gomega.Equal("E"))

}

func TestRoverSpinToTheRightWhenStartingFromEast(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	const orientation string = "E"
	rover := new(Rover)
	rover.SetOrientation(orientation)

	//When
	rover.Spin("R")
	orientationAfterSpinning := rover.Orientation()

	//Then
	gomega.Expect(orientationAfterSpinning).To(gomega.Equal("S"))

}

func TestRoverSpinToTheRightWhenStartingFromSouth(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	const orientation string = "S"
	rover := new(Rover)
	rover.SetOrientation(orientation)

	//When
	rover.Spin("R")
	orientationAfterSpinning := rover.Orientation()

	//Then
	gomega.Expect(orientationAfterSpinning).To(gomega.Equal("W"))

}

func TestRoverSpinToTheRightWhenStartingFromWest(t *testing.T) {
	gomega.RegisterTestingT(t)
	//Given
	const orientation string = "W"
	rover := new(Rover)
	rover.SetOrientation(orientation)

	//When
	rover.Spin("R")
	orientationAfterSpinning := rover.Orientation()

	//Then
	gomega.Expect(orientationAfterSpinning).To(gomega.Equal("N"))

}

func TestRoverMovesWhenInitialOrientationIsNorth(t *testing.T) {
	gomega.RegisterTestingT(t)

	//Given
	const orientation string = "N"
	rover := new(Rover)
	rover.SetOrientation(orientation)
	rover.SetCoordinates(1, 3)

	//When
	rover.Move()

	//Then
	gomega.Expect(rover.currentXCoordinate).To(gomega.Equal(1))
	gomega.Expect(rover.currentYCoordinate).To(gomega.Equal(4))
	gomega.Expect(rover.currentOrientation).To(gomega.Equal(orientation))

}

func TestRoverMovesWhenInitialOrientationIsSouth(t *testing.T) {
	gomega.RegisterTestingT(t)

	//Given
	const orientation string = "S"
	rover := new(Rover)
	rover.SetOrientation(orientation)
	rover.SetCoordinates(1, 3)

	//When
	rover.Move()

	//Then
	gomega.Expect(rover.currentXCoordinate).To(gomega.Equal(1))
	gomega.Expect(rover.currentYCoordinate).To(gomega.Equal(2))
	gomega.Expect(rover.currentOrientation).To(gomega.Equal(orientation))

}

func TestRoverMovesWhenInitialOrientationIsWest(t *testing.T) {
	gomega.RegisterTestingT(t)

	//Given
	const orientation string = "W"
	rover := new(Rover)
	rover.SetOrientation(orientation)
	rover.SetCoordinates(1, 3)

	//When
	rover.Move()

	//Then
	gomega.Expect(rover.currentXCoordinate).To(gomega.Equal(0))
	gomega.Expect(rover.currentYCoordinate).To(gomega.Equal(3))
	gomega.Expect(rover.currentOrientation).To(gomega.Equal(orientation))

}

func TestRoverMovesWhenInitialOrientationIsEast(t *testing.T) {
	gomega.RegisterTestingT(t)

	//Given
	const orientation string = "E"
	rover := new(Rover)
	rover.SetOrientation(orientation)
	rover.SetCoordinates(1, 3)

	//When
	rover.Move()

	//Then
	gomega.Expect(rover.currentXCoordinate).To(gomega.Equal(2))
	gomega.Expect(rover.currentYCoordinate).To(gomega.Equal(3))
	gomega.Expect(rover.currentOrientation).To(gomega.Equal(orientation))

}
