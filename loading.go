package main

import (
	"fmt"
	"time"
)

//i don't have the code that actually does stuff, so i'll just add random tasks

func loading_screen(ip string) {
	clear()
	progress := ProgressBar{
		total:       100,
		length:      50,
		last_suffix: 0,
		enabled:     true,
	}
	time.Sleep(1 * time.Second)
	progress.change(0, "Loading ...", "Initializing\n")
	time.Sleep(500 * time.Millisecond)
	progress.change(20, "Loading ...", "Connecting to "+"Localhost:"+ip+"\n")
	time.Sleep(1 * time.Second)
	progress.change(40, "Loading ...", "We will be done soon\n")
	time.Sleep(1 * time.Second)
	progress.change(60, "Loading ...", "Welcome to Tick\n")
	time.Sleep(1 * time.Second)
	progress.change(80, "Loading ...", "Downloading Assests\n")
	time.Sleep(1 * time.Second)
	progress.change(100, "Loading ...", "Final task\n")
	time.Sleep(1 * time.Second)
	progress.clean()
	clear()
}

func PrintTitle() {
	fmt.Println(Cyan + "████████ ██  ██████     ████████  █████   ██████     ████████  ██████  ███████")
	fmt.Println("   ██    ██ ██             ██    ██   ██ ██             ██    ██    ██ ██      ")
	fmt.Println("   ██    ██ ██             ██    ███████ ██             ██    ██    ██ █████")
	fmt.Println("   ██    ██ ██             ██    ██   ██ ██             ██    ██    ██ ██")
	fmt.Println("   ██    ██  ██████        ██    ██   ██  ██████        ██     ██████  ███████" + Reset)
	fmt.Println()
	fmt.Println()
}
