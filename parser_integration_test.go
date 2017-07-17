package marsrover

import "testing"
import "github.com/stretchr/testify/assert"

func TestParserSetConfigurationForRover(t *testing.T)  {
	//Given
	parser := Parser{}
	rover := new(Rover)

	//When
	parser.ParsePlateauDimensions(rover,"10 10")
	parser.ParseCoordinatesAndOrientation(rover,"2 3 N")
	parser.ParseSpinAndMovement(rover,"RMLMM")
	xCoordinate,yCoordinate := rover.GetCoordinates()

	//Then
	assert.Equal(t,"N",rover.GetOrientation())
	assert.Equal(t,3,xCoordinate)
	assert.Equal(t,5,yCoordinate)

}
