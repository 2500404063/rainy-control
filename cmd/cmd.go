package cmd

import (
	"bytes"
	"os"
	"os/exec"
)

var current_path string
var path string

func init() {
	current_path, _ = os.UserHomeDir()
	path = "C:/Windows/System32/cmd.exe"
}

func ExecCmd(cmdline string, dir string, dirType string) ([]byte, error) {
	var out_buf, err_buf bytes.Buffer
	if dir != "" {
		if dirType == "cd" {
			current_path = dir
		} else if dirType == "path" {
			path = dir
		}
		return out_buf.Bytes(), nil
	} else {
		cc := exec.Command("cmd", "/c", cmdline)
		cc.Dir = current_path
		cc.Path = path
		cc.Stdout = &out_buf
		cc.Stderr = &err_buf
		err := cc.Run()
		return out_buf.Bytes(), err
	}
}
