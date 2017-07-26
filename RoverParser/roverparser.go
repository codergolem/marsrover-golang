package RoverParser

import "marsRover/marsrover"

type RoverParser interface {
	ParseCoordinatesAndOrientation(marsrover.MarsExplorer,string)
	ParseSpinAndMovement(marsrover.MarsExplorer,string)
	ParsePlateauDimensions(marsrover.MarsExplorer,string)
}
