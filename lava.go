package lava

/*
#include "player.h"
#cgo CFLAGS:  -I /usr/local/include
#cgo LDFLAGS: -lavformat -lavcodec -lswscale -lavutil -lswresample -lz -framework sdl2
*/
import "C"
import (
	"fmt"
	"strings"
)

func init() {
	C.global_init()
}

type LAVA struct {
	Volume   int
	Filename string
}

func (_ *LAVA) LoadFile(filename string) {
	filename = strings.Trim(filename, "\n")
	filename = strings.Trim(filename, "\t")
	filename = strings.Trim(filename, "\r")
	filename = strings.Trim(filename, " ")
	fmt.Println("Receive file: " + filename)
	C.cp_load_file(C.CString(filename))
}

func (_ *LAVA) Pause() {
	C.cp_pause_audio()
}
