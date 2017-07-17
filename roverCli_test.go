package marsrover

//import (
//	"testing"
//	"github.com/stretchr/testify/assert"
//	"os"
//	"io/ioutil"
//)

//func TestConsumePlateauDimensions(t *testing.T)  {
//	//Given
//	commandLine = new(RoverCli)
//	commandLine.setState(0)
//
//	commandLine.consume("5 5")
//
//
//
//}
//
//func TestCliShouldDisplayInitialInstructionsAndWaitForPlateauDimensions(t *testing.T)  {
//	//before
//	rescueStdout := os.Stdout
//	r, w, _ := os.Pipe()
//	os.Stdout = w
//
//	//Given
//	commandLine := new(RoverCommandLine)
//
//	//When
//	commandLine.start()
//
//	w.Close()
//	out, _ := ioutil.ReadAll(r)
//	os.Stdout = rescueStdout
//
//	//Then
//	assert.Equal(t,"Please provide plateau dimensions:",string(out))
//	assert.Equal(t,commandLine.getState(),1)
//
//
//
//}


