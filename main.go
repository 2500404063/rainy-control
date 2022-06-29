package main

import (
	"time"

	"felix.com/rainy/v1/basic"
	"github.com/kardianos/service"
)

func trying() {
	for {
		ticker := time.NewTicker(3 * time.Second)
		<-ticker.C
		if !basic.Connected {
			basic.Connect()
		}
	}
}

type Program struct{}

func (*Program) Start(s service.Service) error {
	go trying()
	return nil
}

func (*Program) Stop(s service.Service) error {
	return nil
}

func main() {
	pro := &Program{}
	basic.Letrun(pro, true)
}
