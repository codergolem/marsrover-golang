package marsrover

import (
	"strings"
	"strconv"
)

type Parser struct {

}

func (parser *Parser)ParseCoordinatesAndOrientation(rover MarsExplorer,coordinatesAndOrientation string)  {
	parsedCoordinatesAndOrientation := strings.Split(coordinatesAndOrientation," ")

	xCoordinate,_ := strconv.Atoi(parsedCoordinatesAndOrientation[0])
	yCoordinate,_ := strconv.Atoi(parsedCoordinatesAndOrientation[1])
	orientation := parsedCoordinatesAndOrientation[2]


	rover.SetCoordinates(xCoordinate,yCoordinate)
	rover.SetOrientation(orientation)
}

func (parser *Parser)ParseSpinAndMovement(rover MarsExplorer,spindAndMovementsInstructions string)  {
	parsedSpinAndMoveInstructions := strings.Split(spindAndMovementsInstructions,"")

	for _,singleInstruction := range parsedSpinAndMoveInstructions {
		if(singleInstruction == "L" || singleInstruction == "R") {
			rover.Spin(singleInstruction)
		}
		if(singleInstruction == "M") {
			rover.Move()
		}
	}

}

func (parser *Parser)ParsePlateauDimensions(rover MarsExplorer,plateauDimensions string)  {
	parsedPlateauDimensions := strings.Split(plateauDimensions," ")
	parsedXCoordinate,_ := strconv.Atoi(parsedPlateauDimensions[0])
	parsedYCoordinate,_ := strconv.Atoi(parsedPlateauDimensions[1])

	plateau:= Plateau{parsedXCoordinate,parsedYCoordinate}

	rover.SetPlateau(plateau)

}
