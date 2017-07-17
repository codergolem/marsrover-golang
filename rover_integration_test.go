package marsrover

import (
	"testing"
	"github.com/onsi/gomega"
)

func TestMarsRoversMovesToTheNorth(t *testing.T)  {
	gomega.RegisterTestingT(t)

	//Given
	rover := new(Rover)
	rover.SetOrientation("E")
	rover.SetCoordinates(1,2)

	//When
	rover.Spin("L")
	rover.Move()

	//Then
	gomega.Expect(rover.currentXCoordinate).To(gomega.Equal(1))
	gomega.Expect(rover.currentYCoordinate).To(gomega.Equal(3))
	gomega.Expect(rover.currentOrientation).To(gomega.Equal("N"))
}

func TestMarsRoversMovesToTheEast(t *testing.T)  {
	gomega.RegisterTestingT(t)

	//Given
	rover := new(Rover)
	rover.SetOrientation("N")
	rover.SetCoordinates(1,2)

	//When
	rover.Spin("R")
	rover.Move()
	rover.Spin("L")
	rover.Move()
	rover.Spin("R")
	rover.Move()

	//Then
	gomega.Expect(rover.currentXCoordinate).To(gomega.Equal(3))
	gomega.Expect(rover.currentYCoordinate).To(gomega.Equal(3))
	gomega.Expect(rover.currentOrientation).To(gomega.Equal("E"))
}
