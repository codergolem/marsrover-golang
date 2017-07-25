package cmdparser

import (
	"strings"
	"strconv"
	"marsRover/rover"
	"marsRover/plateau"
)

type RoverParser interface {
	ParseCoordinatesAndOrientation(rover.MarsExplorer,string)
	ParseSpinAndMovement(rover.MarsExplorer,string)
	ParsePlateauDimensions(rover.MarsExplorer,string)
}


type DefaultParser struct {

}

func (parser *DefaultParser)ParseCoordinatesAndOrientation(rover rover.MarsExplorer,coordinatesAndOrientation string)  {
	parsedCoordinatesAndOrientation := strings.Split(coordinatesAndOrientation," ")

	xCoordinate,_ := strconv.Atoi(parsedCoordinatesAndOrientation[0])
	yCoordinate,_ := strconv.Atoi(parsedCoordinatesAndOrientation[1])
	orientation := parsedCoordinatesAndOrientation[2]


	rover.SetCoordinates(xCoordinate,yCoordinate)
	rover.SetOrientation(orientation)
}

func (parser *DefaultParser)ParseSpinAndMovement(rover rover.MarsExplorer,spindAndMovementsInstructions string)  {
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

func (parser *DefaultParser)ParsePlateauDimensions(rover rover.MarsExplorer,plateauDimensions string)  {
	parsedPlateauDimensions := strings.Split(plateauDimensions," ")
	parsedXCoordinate,_ := strconv.Atoi(parsedPlateauDimensions[0])
	parsedYCoordinate,_ := strconv.Atoi(parsedPlateauDimensions[1])

	plateau:= plateau.Plateau{parsedXCoordinate,parsedYCoordinate}

	rover.SetPlateau(plateau)

}
