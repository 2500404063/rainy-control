package basic

import (
	"io"
	"net/http"
	"os"
)

func Dld(url string, save string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	fd, err := os.Create(save)
	if err != nil {
		return err
	}
	io.Copy(fd, res.Body)
	fd.Close()
	res.Body.Close()
	return nil
}
