package spider

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Spider interface {
	SpiderRun() error
}

type SpiderMp4 struct {
	Url  string
	Path string
	Name string
}

func (s *SpiderMp4) SpiderRun() error {
	rsp, err := http.Get(s.Url)
	if err != nil {
		return err
	}

	defer rsp.Body.Close()

	f, err := os.Create(s.Path + s.Name)
	if err != nil {
		return err
	}
	defer f.Close()

	fileSize, err := io.Copy(f, rsp.Body)
	if err != nil {
		return err
	}

	return nil
}
