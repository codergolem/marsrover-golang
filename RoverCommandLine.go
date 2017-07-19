package marsrover

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type RoverCommandLine struct {
	state int
	Parser RoverParser
	Rover MarsExplorer
}


func (roverCommandLine *RoverCommandLine) SetParser(parser RoverParser) {
	roverCommandLine.Parser = parser
}

func (roverCommandLine *RoverCommandLine) SetRover(rover MarsExplorer) {
	roverCommandLine.Rover = rover
}

func (roverCommandLine *RoverCommandLine) getState() int  {
	return roverCommandLine.state
}

func (roverCommandLine *RoverCommandLine) start() {
	fmt.Print("Please provide plateau dimensions:")
	roverCommandLine.state = 1
	//roverCommandLine.read()

}

func (roverCommandLine *RoverCommandLine) read()  {
	reader := bufio.NewReader(os.Stdin)

	text := ""
	switch roverCommandLine.state {
	case 1:
		for n:= 0; n < 3 && text != "x" ; n++ {
			text, _ = reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			if(n == 0) {
				roverCommandLine.Parser.ParsePlateauDimensions(roverCommandLine.Rover,text)
			}
			if(n == 1) {
				roverCommandLine.Parser.ParseCoordinatesAndOrientation(roverCommandLine.Rover,text)
			}
			if(n == 2) {
				roverCommandLine.Parser.ParseSpinAndMovement(roverCommandLine.Rover,text)
			}
		}
	
	}
}