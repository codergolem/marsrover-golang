package marsrover

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

)


func TestCliShouldDisplayInitialInstructionsAndWaitForPlateauDimensions(t *testing.T)  {
	//before
	stdInAndOutMock := new(StdInAndOutMock)
	stdInAndOutMock.setFakeStdInAndOut()

	//Given
	commandLine := new(RoverCommandLine)

	//When
	commandLine.start()
	stdInAndOutMock.closeFakeWriters()
	stdInAndOutMock.restoreToRealStdInAndOut()
	output := stdInAndOutMock.getCapturedOutput()


	//Then
	assert.Equal(t,"Please provide plateau dimensions:",output)
	assert.Equal(t,commandLine.getState(),1)

}

//func TestReadPlateauDimensions(t *testing.T)  {
//	//before
//
//	stdInAndOutMock := new(StdInAndOutMock)
//	stdInAndOutMock.setFakeStdInAndOut()
//
//	//Given
//	commandLine := new(RoverCommandLine)
//	commandLine.state = 1
//	mockParser := new(MockParser)
//	rover := new(Rover)
//	commandLine.SetParser(mockParser)
//
//	mockParser.On("ParsePlateauDimensions",rover,"10 10").Return()
//	mockParser.On("ParseCoordinatesAndOrientation",rover,"2 3").Return()
//	mockParser.On("ParseSpinAndMovement",rover,"RMLMM").Return()
//
//	//When
//	commandLine.read()
//
//	stdInAndOutMock.stdInMockWriter.Write([]byte("10 10\n"))
//	stdInAndOutMock.stdInMockWriter.Write([]byte("2 3 N\n"))
//	stdInAndOutMock.stdInMockWriter.Write([]byte("RMLMM\n"))
//
//	stdInAndOutMock.closeFakeWriters()
//	stdInAndOutMock.restoreToRealStdInAndOut()
//
//
//
//	//Then
//	assert.Equal(t,2,commandLine.getState())
//	mockParser.AssertExpectations(t)
//
//}

//Mocks


type MockParser struct {
	mock.Mock
}

func (parser *MockParser)ParsePlateauDimensions(rover MarsExplorer,plateauDimensions string) {
	parser.Called(rover,plateauDimensions)
}

func (parser *MockParser)ParseCoordinatesAndOrientation(rover MarsExplorer,coordinatesAndOrientation string) {
	parser.Called(rover,coordinatesAndOrientation)
}

func (parser *MockParser)ParseSpinAndMovement(rover MarsExplorer,spinAndMovementInstructions string) {
	parser.Called(rover,spinAndMovementInstructions)
}

type StdInAndOutMock struct {
	realStdIn *os.File
	realStdOut *os.File
	stdInMockReader *os.File
	stdInMockWriter *os.File
	stdOutMockReader *os.File
	stdOutMockWriter *os.File
}

func (stdInAndOutMock *StdInAndOutMock) setFakeStdInAndOut()   {
	stdInAndOutMock.realStdOut = os.Stdout
	stdInAndOutMock.realStdIn  = os.Stdin

	stdInAndOutMock.stdInMockReader, stdInAndOutMock.stdInMockWriter, _ = os.Pipe()
	stdInAndOutMock.stdOutMockReader, stdInAndOutMock.stdOutMockWriter, _ = os.Pipe()

	os.Stdin = stdInAndOutMock.stdInMockReader
	os.Stdout = stdInAndOutMock.stdOutMockWriter

}

func (stdInAndOutMock *StdInAndOutMock) getCapturedInput() string  {
	output, _ := ioutil.ReadAll(stdInAndOutMock.stdInMockReader)
	return string(output)
}

func (stdInAndOutMock *StdInAndOutMock) getCapturedOutput() string  {
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


