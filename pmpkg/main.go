package pmpkg

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/stdcurse/pm/config"
	"github.com/stdcurse/pm/output"
	"io/ioutil"
	"os"
	"text/template"
)

type Package struct {
	Name         string
	Version      string
	Description  string
	Dependencies []Package

	PackageDir string

	DatabaseEntry *config.DatabaseEntry
}

func NewPackage(name string, db *config.Database, cfg *config.Config) *Package {
	if _, err := os.Stat(cfg.Portdir + "/" + name); err != nil {
		return nil
	}

	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/build.yaml", cfg.Portdir, name))
	output.Check(err, "Something went wrong with package file", true)

	tmpl, err := template.New(name).Parse(string(content))
	output.Check(err, "Something went wrong with parsing template", true)

	var buff bytes.Buffer
	err = tmpl.Execute(&buff, nil)
	output.Check(err, "Something went wrong with parsing template", true)

	var p Package
	output.Check(yaml.Unmarshal(buff.Bytes(), &p), "Something went wrong with package file content", true)

	p.PackageDir = cfg.Portdir + "/" + name
	p.DatabaseEntry = db.Probe(name)

	return &p
}
