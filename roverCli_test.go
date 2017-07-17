package marsrover

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

//import (
//	"testing"
//	"github.com/stretchr/testify/assert"
//	"os"
//	"io/ioutil"
//)

//func TestConsumePlateauDimensions(t *testing.T)  {
//	//Given
//	commandLine = new(Ro)
//	commandLine.setState(0)
//
//	commandLine.consume("5 5")
//
//
//
//}

func TestCliShouldDisplayInitialInstructionsAndWaitForPlateauDimensions(t *testing.T)  {
	//before
	stdOutMock := new(StdOutMock)
	stdOutMock.setFakeStdOut()

	//Given
	commandLine := new(RoverCommandLine)

	//When
	commandLine.start()
	stdOutMock.closeFakeWriter()
	stdOutMock.resetToRealStdOut()
	output := stdOutMock.getCapturedOutput()

	//Then
	assert.Equal(t,"Please provide plateau dimensions:",output)
	assert.Equal(t,commandLine.getState(),1)




}

type StdOutMock struct {
	realStdOut *os.File
	w *os.File
	r *os.File
}

func (stdOutMock *StdOutMock) setFakeStdOut()   {
	stdOutMock.realStdOut = os.Stdout
	stdOutMock.r, stdOutMock.w, _ = os.Pipe()
	os.Stdout = stdOutMock.w

}

func (stdOutMock *StdOutMock) getCapturedOutput() string  {
	output, _ := ioutil.ReadAll(stdOutMock.r)
	return string(output)
}

func (stdOutMock *StdOutMock) resetToRealStdOut() {
	os.Stdout = stdOutMock.realStdOut
}

func (stdOutMock *StdOutMock) closeFakeWriter() {
	stdOutMock.w.Close()
}


