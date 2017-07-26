package marsrover

type MarsExplorer interface {
	SetPlateau(plateau Plateau)
	SetCoordinates(int, int)
	SetOrientation(string)
	Spin(string)
	Move()
}

type Rover struct {
	CurrentXCoordinate int
	CurrentYCoordinate int
	CurrentOrientation string
	Plateau            Plateau
}

func (rover *Rover) SetCoordinates(xCoordinate int, yCoordinate int) {
	rover.CurrentXCoordinate = xCoordinate
	rover.CurrentYCoordinate = yCoordinate

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
		rover.SetCoordinates(rover.CurrentXCoordinate, rover.CurrentYCoordinate+1)
	case "W":
		rover.SetCoordinates(rover.CurrentXCoordinate-1, rover.CurrentYCoordinate)
	case "S":
		rover.SetCoordinates(rover.CurrentXCoordinate, rover.CurrentYCoordinate-1)
	case "E":
		rover.SetCoordinates(rover.CurrentXCoordinate+1, rover.CurrentYCoordinate)

	}
}

func (rover *Rover) Coordinates() (int, int) {
	return rover.CurrentXCoordinate, rover.CurrentYCoordinate
}

func (rover *Rover) SetOrientation(orientation string) {

	rover.CurrentOrientation = orientation
}

func (rover *Rover) SetPlateau(plateau Plateau) {
	rover.Plateau = plateau
}

func (rover *Rover) Orientation() string {
	return rover.CurrentOrientation
}
