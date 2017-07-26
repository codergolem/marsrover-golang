package cmdparser

import (
	"strings"
	"strconv"
	"marsRover/marsrover"
)

type RoverParser interface {
	ParseCoordinatesAndOrientation(marsrover.MarsExplorer,string)
	ParseSpinAndMovement(marsrover.MarsExplorer,string)
	ParsePlateauDimensions(marsrover.MarsExplorer,string)
}


type DefaultParser struct {

}

func (parser *DefaultParser)ParseCoordinatesAndOrientation(rover marsrover.MarsExplorer,coordinatesAndOrientation string)  {
	parsedCoordinatesAndOrientation := strings.Split(coordinatesAndOrientation," ")

	xCoordinate,_ := strconv.Atoi(parsedCoordinatesAndOrientation[0])
	yCoordinate,_ := strconv.Atoi(parsedCoordinatesAndOrientation[1])
	orientation := parsedCoordinatesAndOrientation[2]


	rover.SetCoordinates(xCoordinate,yCoordinate)
	rover.SetOrientation(orientation)
}

func (parser *DefaultParser)ParseSpinAndMovement(rover marsrover.MarsExplorer,spindAndMovementsInstructions string)  {
	parsedSpinAndMoveInstructions := strings.Split(spindAndMovementsInstructions,"")

	for _,singleInstruction := range parsedSpinAndMoveInstructions {
		if singleInstruction == "L" || singleInstruction == "R" {
			rover.Spin(singleInstruction)
		}
		if singleInstruction == "M" {
			rover.Move()
		}
	}

}

func (parser *DefaultParser)ParsePlateauDimensions(rover marsrover.MarsExplorer,plateauDimensions string)  {
	parsedPlateauDimensions := strings.Split(plateauDimensions," ")
	parsedXCoordinate,_ := strconv.Atoi(parsedPlateauDimensions[0])
	parsedYCoordinate,_ := strconv.Atoi(parsedPlateauDimensions[1])

	plateau:= marsrover.Plateau{parsedXCoordinate, parsedYCoordinate}

	rover.SetPlateau(plateau)

}
