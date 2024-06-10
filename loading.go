package main

import (
	"fmt"
	"time"
)

//i don't have the code that actually does stuff, so i'll just add random tasks

func loading_screen(ip string) {
	clear()
	printTitle()
	progress := ProgressBar{
		total:       100,
		length:      50,
		last_suffix: 0,
		enabled:     true,
	}
	time.Sleep(1 * time.Second)
	progress.change(0, "Loading ...", "Doing something\n")
	time.Sleep(500 * time.Millisecond)
	progress.change(20, "Loading ...", "Connecting to "+ip+"\n")
	time.Sleep(1 * time.Second)
	progress.change(40, "Loading ...", "Doing something\n")
	time.Sleep(1 * time.Second)
	progress.change(60, "Loading ...", "Doing something else\n")
	time.Sleep(1 * time.Second)
	progress.change(80, "Loading ...", "Idk\n")
	time.Sleep(1 * time.Second)
	progress.change(100, "Loading ...", "Final task\n")
	time.Sleep(1 * time.Second)
	progress.clean()
	clear()
	printTitle()
}

func printTitle() {
	fmt.Println(Cyan + "████████ ██  ██████     ████████  █████   ██████     ████████  ██████  ███████")
	fmt.Println("   ██    ██ ██             ██    ██   ██ ██             ██    ██    ██ ██      ")
	fmt.Println("   ██    ██ ██             ██    ███████ ██             ██    ██    ██ █████")
	fmt.Println("   ██    ██ ██             ██    ██   ██ ██             ██    ██    ██ ██")
	fmt.Println("   ██    ██  ██████        ██    ██   ██  ██████        ██     ██████  ███████" + Reset)
	fmt.Println()
	fmt.Println()
}
