package marsrover

type MarsExplorer interface {
	SetPlateau(plateau Plateau)
	SetCoordinates(int, int)
	SetOrientation(string)
	Spin(string)
	Move()
}

type Rover struct {
	currentXCoordinate int
	currentYCoordinate int
	currentOrientation string
	Plateau            Plateau
}

func (rover *Rover) SetCoordinates(xCoordinate int, yCoordinate int) {
	rover.currentXCoordinate = xCoordinate
	rover.currentYCoordinate = yCoordinate

}

func (rover *Rover) Spin(direction string) {
	if direction == "L" {
		switch rover.Orientation() {
		case "N":
			rover.SetOrientation("W")
		case "W":
			rover.SetOrientation("S")
		case "S":
			rover.SetOrientation("E")
		case "E":
			rover.SetOrientation("N")

		}
	}

	if direction == "R" {
		switch rover.Orientation() {
		case "N":
			rover.SetOrientation("E")
		case "W":
			rover.SetOrientation("N")
		case "S":
			rover.SetOrientation("W")
		case "E":
			rover.SetOrientation("S")

		}
	}

}

func (rover *Rover) Move() {
	switch rover.Orientation() {
	case "N":
		rover.SetCoordinates(rover.currentXCoordinate, rover.currentYCoordinate+1)
	case "W":
		rover.SetCoordinates(rover.currentXCoordinate-1, rover.currentYCoordinate)
	case "S":
		rover.SetCoordinates(rover.currentXCoordinate, rover.currentYCoordinate-1)
	case "E":
		rover.SetCoordinates(rover.currentXCoordinate+1, rover.currentYCoordinate)

	}
}

func (rover *Rover) Coordinates() (int, int) {
	return rover.currentXCoordinate, rover.currentYCoordinate
}

func (rover *Rover) SetOrientation(orientation string) {
	rover.currentOrientation = orientation
}

func (rover *Rover) SetPlateau(plateau Plateau) {
	rover.Plateau = plateau
}

func (rover *Rover) Orientation() string {
	return rover.currentOrientation
}
