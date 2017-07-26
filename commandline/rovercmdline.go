package commandline

import (
	"bufio"
	"fmt"
	"marsRover/cmdparser"
	"marsRover/marsrover"
	"os"
	"strings"
)

type RoverCommandLine struct {
	Iface  *os.File
	state  int
	Parser cmdparser.RoverParser
	Rover  marsrover.MarsExplorer
	reader *bufio.Reader
}

func (roverCommandLine *RoverCommandLine) SetParser(parser cmdparser.RoverParser) {
	roverCommandLine.Parser = parser
}

func (roverCommandLine *RoverCommandLine) SetRover(rover marsrover.MarsExplorer) {
	roverCommandLine.Rover = rover
}

func (roverCommandLine *RoverCommandLine) getState() int {
	return roverCommandLine.state
}

func (roverCommandLine *RoverCommandLine) start(iface *os.File) {
	roverCommandLine.Iface = iface
	fmt.Print("Please provide plateau dimensions,marsrover initial position and movement instructions:")
	roverCommandLine.state = 1
	roverCommandLine.reader = bufio.NewReader(roverCommandLine.Iface)

}

func (roverCommandLine *RoverCommandLine) read() {
	text := ""
	switch roverCommandLine.state {
	case 1:
		for n := 0; n < 3 && text != "x"; n++ {
			text, _ = roverCommandLine.reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			if n == 0 {
				roverCommandLine.Parser.ParsePlateauDimensions(roverCommandLine.Rover, text)
			}
			if n == 1 {
				roverCommandLine.Parser.ParseCoordinatesAndOrientation(roverCommandLine.Rover, text)
			}
			if n == 2 {
				roverCommandLine.Parser.ParseSpinAndMovement(roverCommandLine.Rover, text)
			}
		}

	}
	roverCommandLine.state = 2
}
