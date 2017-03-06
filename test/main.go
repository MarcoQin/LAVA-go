package main

import (
	"bufio"
	"fmt"
	"github.com/marcoqin/LAVA-go"
	"os"
	"strconv"
	"strings"
)

func main() {
	core := lava.LAVA{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path: ")
	var text string
	text, _ = reader.ReadString('\n')
	fmt.Println("You entered: " + text)
	core.LoadFile(strings.Replace(text, "\\", "", -1))
	// core.LoadFile("/Users/marcoqin/media/1.flac")
	// text, _ = reader.ReadString('\n')
	// fmt.Println(text)
	// core.Pause()

	for {
		text, _ = reader.ReadString('\n')
		text = strings.Trim(text, "\n")
		if text == "q" {
			break
		}
		if strings.Contains(text, "set_volume") {
			r := strings.Fields(text)
			if len(r) > 1 {
				volume := r[1]
				v, _ := strconv.Atoi(volume)
				core.SetVolume(v)
			}
		}
	}
}
