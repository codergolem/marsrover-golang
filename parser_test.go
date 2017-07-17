package marsrover

import (
	"testing"
	"github.com/stretchr/testify/mock"
)


func TestParseInitialCoordinatesAndOrientation(t *testing.T )  {
	//Given
	roverMock := new(RoverMock)
	roverParser:= new(Parser)
	roverMock.On("SetCoordinates",3,5).Return()
	roverMock.On("SetOrientation","N").Return()

	//When
	roverParser.ParseCoordinatesAndOrientation(roverMock,"3 5 N")

	//Then
	roverMock.AssertExpectations(t)

}

func TestParseSpinAndMovement(t *testing.T )  {
	//Given
	roverMock := new(RoverMock)
	roverParser:= new(Parser)
	roverMock.On("Spin","L").Return()
	roverMock.On("Move").Return()

	//When
	roverParser.ParseSpinAndMovement(roverMock,"LM")

	//Then
	roverMock.AssertExpectations(t)
}

func TestParsePlateauDimensions(t *testing.T) {
	//given
	roverMock := new(RoverMock)
	roverParser:= new(Parser)
	plateau := Plateau{5,6}
	roverMock.On("SetPlateau",plateau).Return()

	//When
	roverParser.ParsePlateauDimensions(roverMock,"5 6")

	//
	roverMock.AssertExpectations(t)

}


//Mocks

type RoverMock struct {
	mock.Mock
}

func (roverMock *RoverMock) SetCoordinates(xCoordinate,yCoordinate int)  {
	roverMock.Called(xCoordinate,yCoordinate)
}

func (roverMock *RoverMock) SetOrientation(orientation string)  {
	roverMock.Called(orientation)
}

func (roverMock *RoverMock) Spin(direction string)  {
	roverMock.Called(direction)
}

func (roverMock *RoverMock) Move()  {
	roverMock.Called()
}

func (roverMock *RoverMock) SetPlateau(plateau Plateau) {
	roverMock.Called(plateau)
}