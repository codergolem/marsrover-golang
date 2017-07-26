package commandline

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"marsRover/marsrover"
	"os"
	"testing"
)

func TestCliShouldDisplayInitialInstructionsAndWaitForPlateauDimensions(t *testing.T) {
	//before
	stdInAndOutMock := new(StdInAndOutMock)
	stdInAndOutMock.setFakeStdInAndOut()

	//Given
	commandLine := new(RoverCommandLine)
	//When
	commandLine.start(stdInAndOutMock.stdInMockReader)
	stdInAndOutMock.closeFakeWriters()
	stdInAndOutMock.restoreToRealStdInAndOut()
	output := stdInAndOutMock.getCapturedOutput()

	//Then
	assert.Equal(t, "Please provide plateau dimensions,marsrover initial position and movement instructions:", output)
	assert.Equal(t, commandLine.getState(), 1)

}

func TestReadInstructionsandCalculateRoverFinalPosition(t *testing.T) {
	//before

	reader, writer, _ := os.Pipe()

	//Given
	commandLine := new(RoverCommandLine)
	commandLine.Iface = reader
	commandLine.state = 1
	mockParser := new(MockParser)
	rover := new(RoverMock)
	commandLine.SetParser(mockParser)
	commandLine.SetRover(rover)
	commandLine.start(reader)

	mockParser.On("ParsePlateauDimensions", rover, "10 10").Return()
	mockParser.On("ParseCoordinatesAndOrientation", rover, "2 3 N").Return()
	mockParser.On("ParseSpinAndMovement", rover, "RMLMM").Return()

	writer.Write([]byte("10 10\n"))
	writer.Write([]byte("2 3 N\n"))
	writer.Write([]byte("RMLMM\n"))

	//When
	commandLine.read()

	writer.Close()

	//Then
	assert.Equal(t, 2, commandLine.getState())
	mockParser.AssertExpectations(t)

}

//Mocks

type MockParser struct {
	mock.Mock
}

func (parser *MockParser) ParsePlateauDimensions(rover marsrover.MarsExplorer, plateauDimensions string) {
	parser.Called(rover, plateauDimensions)
}

func (parser *MockParser) ParseCoordinatesAndOrientation(rover marsrover.MarsExplorer, coordinatesAndOrientation string) {
	parser.Called(rover, coordinatesAndOrientation)
}

func (parser *MockParser) ParseSpinAndMovement(rover marsrover.MarsExplorer, spinAndMovementInstructions string) {
	parser.Called(rover, spinAndMovementInstructions)
}

type StdInAndOutMock struct {
	realStdIn        *os.File
	realStdOut       *os.File
	stdInMockReader  *os.File
	stdInMockWriter  *os.File
	stdOutMockReader *os.File
	stdOutMockWriter *os.File
}

func (stdInAndOutMock *StdInAndOutMock) setFakeStdInAndOut() {
	stdInAndOutMock.realStdOut = os.Stdout
	stdInAndOutMock.realStdIn = os.Stdin

	stdInAndOutMock.stdInMockReader, stdInAndOutMock.stdInMockWriter, _ = os.Pipe()
	stdInAndOutMock.stdOutMockReader, stdInAndOutMock.stdOutMockWriter, _ = os.Pipe()

	os.Stdin = stdInAndOutMock.stdInMockReader
	os.Stdout = stdInAndOutMock.stdOutMockWriter

}

func (stdInAndOutMock *StdInAndOutMock) getCapturedInput() string {
	output, _ := ioutil.ReadAll(stdInAndOutMock.stdInMockReader)
	return string(output)
}

func (stdInAndOutMock *StdInAndOutMock) getCapturedOutput() string {
	output, _ := ioutil.ReadAll(stdInAndOutMock.stdOutMockReader)
	return string(output)
}

func (stdInAndOutMock *StdInAndOutMock) restoreToRealStdInAndOut() {
	os.Stdin = stdInAndOutMock.realStdIn
	os.Stdout = stdInAndOutMock.realStdOut
}

func (stdInAndOutMock *StdInAndOutMock) closeFakeWriters() {
	stdInAndOutMock.stdInMockWriter.Close()
	stdInAndOutMock.stdOutMockWriter.Close()
}

//Mocks

type RoverMock struct {
	mock.Mock
}

func (roverMock *RoverMock) SetCoordinates(xCoordinate, yCoordinate int) {
	roverMock.Called(xCoordinate, yCoordinate)
}

func (roverMock *RoverMock) SetOrientation(orientation string) {
	roverMock.Called(orientation)
}

func (roverMock *RoverMock) Spin(direction string) {
	roverMock.Called(direction)
}

func (roverMock *RoverMock) Move() {
	roverMock.Called()
}

func (roverMock *RoverMock) SetPlateau(plateau marsrover.Plateau) {
	roverMock.Called(plateau)
}
