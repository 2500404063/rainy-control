package basic

import (
	"net"
	"strings"
)

const (
	IP = "127.0.0.1:6266"
)

var Connected bool = false

func Connect() {
	tcp_eip, _ := net.ResolveTCPAddr("tcp4", IP)
	tcp_conn, err := net.DialTCP("tcp4", nil, tcp_eip)
	if err == nil {
		go recvHandler(tcp_conn)
		Connected = true
	}
}

func CmdParse(cmd string) (string, []string) {
	start := strings.Index(cmd, ":")
	paras := strings.Split(cmd[start+1:], ",")
	return cmd[0:start], paras
}

func recvHandler(client *net.TCPConn) {
	client.SetKeepAlive(false)
	client.SetReadBuffer(1024)
	var recv_buf = make([]byte, 1024)
	for {
		len, err := client.Read(recv_buf)
		if err != nil {
			Connected = false
			client.Close()
			break
		}
		cmds := strings.Split(string(recv_buf[0:len]), "|")
		for _, v := range cmds {
			if v != "" {
				code, paras := CmdParse(v)
				Dispatch(client, code, paras)
			}
		}
	}
}
