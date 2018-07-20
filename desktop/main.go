package main

import (
	"go_messenger/desktop/windows"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	windows.Draw()
}
