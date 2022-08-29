package modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 服务端指令控制工具
func ServeCommand() {
	fmt.Print("\n")
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		command, _ := inputReader.ReadString('\n')
		switch strings.Join(strings.Fields(command), "") {
		// UserList 包括大小写两种模式
		case "UserList":
			GetUserList()
		case "userList":
			GetUserList()
		// Help 包括大小写两种模式
		case "Help":
			GetHelp()
		case "help":
			GetHelp()
		}
	}
}
