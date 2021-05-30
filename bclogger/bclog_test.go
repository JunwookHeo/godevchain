package bclogger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Init(true)

	x := 20

	INFOMSG("Info", "Test", x)
	INFOMSGF("Info Test %d\n", x)
	WARNMSG("Warning")
	WARNMSGF("Warning Test %d\n", x)
	ERRORMSG("Error")
	ERRORMSGF("Error Test %d\n", x)

	//FATALMSG("Fatal")
	//FATALMSGF("Fatal Test %d\n", x)
}
