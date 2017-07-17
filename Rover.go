package marsrover


type MarsExplorer interface {
	SetPlateau(Plateau)
	SetCoordinates(int,int)
	SetOrientation(string)
	Spin(string)
	Move()

}

type Rover struct {
	currentXCoordinate int
	currentYCoordinate int
	currentOrientation string
	Plateau Plateau
	Parser	Parser
}

func (rover *Rover)SetCoordinates(xCoordinate int,yCoordinate int)  {
	rover.currentXCoordinate = xCoordinate
	rover.currentYCoordinate = yCoordinate
	
}

func (rover *Rover) Spin(direction string) {
	if (direction == "L") {
		switch rover.GetOrientation() {
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

	if (direction == "R") {
		switch rover.GetOrientation() {
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

func (rover *Rover) Move()  {
	switch rover.GetOrientation() {
	case "N":
		rover.SetCoordinates(rover.currentXCoordinate,rover.currentYCoordinate + 1)
	case "W":
		rover.SetCoordinates(rover.currentXCoordinate - 1,rover.currentYCoordinate)
	case "S":
		rover.SetCoordinates(rover.currentXCoordinate,rover.currentYCoordinate - 1)
	case "E":
		rover.SetCoordinates(rover.currentXCoordinate + 1,rover.currentYCoordinate)

	}
}

func (rover *Rover) GetCoordinates() (int,int) {
	return rover.currentXCoordinate, rover.currentYCoordinate
}

func (rover *Rover) SetOrientation(orientation string) () {

	rover.currentOrientation = orientation
}

func (rover *Rover) SetPlateau(plateau Plateau)  {
	rover.Plateau = plateau
}

func (rover *Rover) GetOrientation()  string {
	return rover.currentOrientation
}

