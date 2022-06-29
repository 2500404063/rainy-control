package device

import (
	"os"
)

func GetName() string {
	cachedir, _ := os.UserHomeDir()
	fd, err := os.OpenFile(cachedir+"/.rainy", os.O_RDONLY, 0)
	if err == nil {
		var buf = make([]byte, 128)
		fd.Read(buf)
		return string(buf)
	}
	return "unnamed"
}

func SetName(name string) {
	cachedir, _ := os.UserHomeDir()
	fd, err := os.OpenFile(cachedir+"/.rainy", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err == nil {
		fd.Write([]byte(name))
	}
	fd.Close()
}
