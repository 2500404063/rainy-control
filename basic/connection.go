package basic

import (
	"net"
)

const (
	Port = 6266
)

type ClientHandle struct {
	id     int
	name   string
	socket *net.TCPConn
}

var Server *net.TCPListener
var Clients = make(map[int]*ClientHandle)

func StartServer() (err error) {
	listener, err := net.ListenTCP("tcp4", &net.TCPAddr{
		IP:   nil,
		Port: Port,
	})
	if err != nil {
		panic(err.Error())
	}
	Server = listener
	go clientAcception()
	return nil
}

func clientAcception() {
	var client_count int
	for {
		client, _ := Server.AcceptTCP()
		var buf = make([]byte, 128)
		client.Write([]byte(CmdFormat(GET_NAME, nil)))
		len, _ := client.Read(buf)
		name := string(buf[0:len])
		Clients[client_count] = &ClientHandle{
			id:     client_count,
			name:   name,
			socket: client,
		}
		client_count++
	}
}
