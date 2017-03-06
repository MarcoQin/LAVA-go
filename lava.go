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

func (this *LAVA) LoadFile(filename string) {
	filename = strings.Trim(filename, "\n")
	filename = strings.Trim(filename, "\t")
	filename = strings.Trim(filename, "\r")
	filename = strings.Trim(filename, " ")
	fmt.Println("Receive file: " + filename)
	this.Filename = filename
	C.cp_load_file(C.CString(filename))
}

func (_ *LAVA) Pause() {
	C.cp_pause_audio()
}

func (_ *LAVA) Stop() {
	C.cp_stop_audio()
}

func (_ *LAVA) SetVolume(volume int) {
	C.cp_set_volume(C.int(volume))
}

func (_ *LAVA) SeekByPercent(percent float64) {
	C.cp_seek_audio(C.double(percent))
}

func (_ *LAVA) SeekBySecond(sec int) {
	C.cp_seek_audio_by_sec(C.int(sec))
}

func (_ *LAVA) SeekToPosition(pos int) {
	C.cp_seek_audio_by_absolute_pos(C.int(pos))
}

func (_ *LAVA) TimeLength() int {
	length := C.cp_get_time_length()
	return int(length)
}

func (_ *LAVA) CurrentTimePosition() float64 {
	pos := C.cp_get_current_time_pos()
	return float64(pos)
}

func (_ *LAVA) IsStopping() bool {
	state := C.cp_is_stopping()
	if int(state) > 0 {
		return true
	} else {
		return false
	}
}
