package tools

import (
	"github.com/stdcurse/pm/output"
	"io"
	"net/http"
	"os"
)

func NeedRoot() {
	if os.Geteuid() != 0 {
		output.ErrorSimple("This action requires root access")
	}
}

func DownloadFile(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
