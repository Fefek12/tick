package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

type ProgressBar struct {
	total       int
	length      int
	last_suffix int
	enabled     bool
}

func (prb *ProgressBar) change(amount int, prefix, suffix string) {
	if prb.enabled {
		percent := float64(amount) / float64(prb.total)
		filledLength := int(float64(prb.length) * percent)
		fill := "█"
		end := "█"
		if amount == prb.total {
			end = "█"
		}
		for len(suffix) < prb.last_suffix {
			suffix += " "
		}
		bar := strings.Repeat(fill, filledLength) + end + strings.Repeat("-", (prb.length-filledLength))
		fmt.Printf("\r%s [%s] %s", prefix, bar, suffix)
		// if amount == prb.total {
		// 	fmt.Println()
		// }
		prb.last_suffix = len(suffix) + 3
	}
}

func (prb ProgressBar) clean() {
	if prb.enabled {
		fmt.Printf("\r" + strings.Repeat(" ", prb.last_suffix+prb.length+2) + "\r")
	}
}
