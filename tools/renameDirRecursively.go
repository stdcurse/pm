package tools

import (
	"io/ioutil"
	"os"
)

func RenameDirectoryRecursively(from, to string) error {
	d, err := ioutil.ReadDir(from)
	if err != nil {
		return err
	}

	for _, x := range d {
		if _, err := os.Stat(to + "/" + x.Name()); err != nil {
			os.RemoveAll(to + "/" + x.Name())
		}
		if err = os.Rename(from+"/"+x.Name(), to+"/"+x.Name()); err != nil {
			return err
		}
	}

	return nil
}
