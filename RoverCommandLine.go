package marsrover

import "fmt"

type RoverCommandLine struct {
	state int
}


func (roverCommandLine *RoverCommandLine) start() {
	fmt.Print("Please provide plateau dimensions:")
	roverCommandLine.state = 1
}

func (roverCommandLine *RoverCommandLine) processPlateauDimensions(plateauDimensions string)  {

}

func (roverCommandLine *RoverCommandLine) getState() int  {
	return roverCommandLine.state
}