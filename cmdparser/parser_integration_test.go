package cmdparser

import "testing"
import (
	"github.com/stretchr/testify/assert"
	"marsRover/marsrover"
)

func TestParserSetConfigurationForRover(t *testing.T)  {
	//Given
	parser := DefaultParser{}
	rover := new(marsrover.Rover)

	//When
	parser.ParsePlateauDimensions(rover,"10 10")
	parser.ParseCoordinatesAndOrientation(rover,"2 3 N")
	parser.ParseSpinAndMovement(rover,"RMLMM")
	xCoordinate,yCoordinate := rover.Coordinates()

	//Then
	assert.Equal(t,"N",rover.Orientation())
	assert.Equal(t,3,xCoordinate)
	assert.Equal(t,5,yCoordinate)

}
