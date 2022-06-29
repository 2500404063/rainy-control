package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"felix.com/rainy/v1/basic"
)

func main() {
	basic.StartServer()
	fmt.Println("Server has started!")
	for {
		fmt.Print("Rainy@root> ")
		reader := bufio.NewReader(os.Stdin)
		cmd_line, _ := reader.ReadString('\n')
		cmd_line = strings.Trim(cmd_line, "\r")
		cmd_line = strings.Trim(cmd_line, "\n")
		cmd_line = strings.Trim(cmd_line, "\r\n")
		basic.Dispatch(cmd_line)
	}
}
