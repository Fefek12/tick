package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Fefek12/tick/Server"
)

func intro() {
	PrintTitle()
	buffReader := bufio.NewReader(os.Stdin)
	fmt.Print("Join or Host: ")
	response, _ := buffReader.ReadString('\n')
	res := strings.ToLower(strings.TrimSpace(response))
	switch res {
	case "join":
		fmt.Print("Enter Port under LocalHost to Join: ")
		joinAddr, _ := buffReader.ReadString('\n')
		joinAddr = strings.TrimSpace(joinAddr)
		loading_screen(joinAddr)
		c, err := NewClient(joinAddr)
		if err != nil {
			panic(err)
		}
		for {
			c.Render()
			buffReader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter X or O and Cords")
			res, err := buffReader.ReadString('\n')
			if err != nil {
				fmt.Println("Error Processing Input")
			}
			go c.SendDelta(res)
		}
	case "host":
		fmt.Print("Enter Port under LocalHost to Host: ")
		hostAddr, _ := buffReader.ReadString('\n')
		hostAddr = strings.TrimSpace(hostAddr)
		s := Server.NewServer(hostAddr)
		s.Start()
	}
}
