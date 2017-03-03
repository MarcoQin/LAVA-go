package main

import (
	"bufio"
	"fmt"
	"github.com/marcoqin/LAVA-go"
	"os"
	"strings"
)

func main() {
	core := lava.LAVA{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("You entered: " + text)
	core.LoadFile(strings.Replace(text, "\\", "", -1))
	// core.LoadFile("/Users/marcoqin/media/1.flac")
	text, _ = reader.ReadString('\n')
	fmt.Println(text)
	core.Pause()
}
