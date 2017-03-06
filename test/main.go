package main

import (
	"bufio"
	"fmt"
	"github.com/marcoqin/LAVA-go"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var core lava.LAVA
var cmds map[string]func(string)

const Red string = "\033[0;31m"
const Green string = "\033[0;32m"
const Orange string = "\033[0;33m"
const Blue string = "\033[0;34m"
const Purple string = "\033[0;35m"
const Cyan string = "\033[0;36m"
const LightGray string = "\033[0;37m"
const DarkGray string = "\033[1;30m"
const Yellow string = "\033[1;33m"
const White string = "\033[1;37m"
const End string = "\033[0m"

func Log(s string, color string) {
	for _, line := range strings.Split(s, "\n") {
		fmt.Printf("%s%s%s\n", color, line, End)
	}
}

func set_volume(cmd string) {
	r := strings.SplitN(cmd, " ", 2)
	if len(r) > 1 {
		volume := r[1]
		v, _ := strconv.Atoi(volume)
		core.SetVolume(v)
	}
}

func load(cmd string) {
	r := strings.SplitN(cmd, " ", 2)
	if len(r) > 1 {
		file := r[1]
		file = strings.Trim(file, "\n")
		core.LoadFile(strings.Replace(file, "\\", "", -1))
	}
}

func pause(_ string) {
	core.Pause()
}

func stop(_ string) {
	core.Stop()
}

func stat(_ string) {
	var s string
	s = fmt.Sprintf("Now  Playing: %s\nTotal Length: %d\nCurrent  Pos: %.2f", core.Filename, core.TimeLength(), core.CurrentTimePosition())
	Log(s, Blue)
}

func main() {
	core = lava.LAVA{}
	cmds = make(map[string]func(string))
	cmds["set_volume"] = set_volume
	cmds["load"] = load
	cmds["pause"] = pause
	cmds["stop"] = stop
	cmds["stat"] = stat

	reader := bufio.NewReader(os.Stdin)
	var cmd string
	for {
		fmt.Printf("%sPlayer%s %s>>>%s ", Green, End, Red, End)
		cmd, _ = reader.ReadString('\n')
		cmd = strings.Trim(cmd, "\n")
		if cmd == "q" {
			break
		}
		r := strings.SplitN(cmd, " ", 2)
		if len(r) > 0 {
			f, ok := cmds[r[0]]
			if ok {
				f(cmd)
			}
		}
	}
}
