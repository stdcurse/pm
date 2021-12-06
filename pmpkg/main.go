/*
	Copyright (c) 2021 Nikita Nikiforov <vokestd@gmail.com>

	This software is provided 'as-is', without any express or implied
	warranty. In no event will the authors be held liable for any damages
	arising from the use of this software.

	Permission is granted to anyone to use this software for any purpose,
	including commercial applications, and to alter it and redistribute it
	freely, subject to the following restrictions:

	1. The origin of this software must not be misrepresented; you must not
		 claim that you wrote the original software. If you use this software
		 in a product, an acknowledgement in the product documentation would be
		 appreciated but is not required.
	2. Altered source versions must be plainly marked as such, and must not be
		 misrepresented as being the original software.
	3. This notice may not be removed or altered from any source distribution.
*/

package pmpkg

import (
	"bytes"
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/stdcurse/pm/tools"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/goccy/go-yaml"
	"github.com/stdcurse/pm/config"
	"github.com/stdcurse/pm/output"
)

type Package struct {
	Name         string
	Version      string
	Description  string
	Dependencies []*Package
	Script       map[string]interface{}
	Sources      []string
	Release      uint

	PackageDir string

	DatabaseEntry *config.DatabaseEntry
}

var c *config.Config

func NewPackage(name string, db *config.Database, cfg *config.Config) *Package {
	c = cfg
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

	if p.Release == 0 {
		p.Release = 1
	}

	return &p
}

func (p *Package) Fetch() {
	pkgdir := fmt.Sprintf("%s/%s/pkgdir", c.Portdir, p.Name)
	srcdir := fmt.Sprintf("%s/%s/srcdir", c.Portdir, p.Name)

	if _, err := os.Stat(srcdir); err == nil {
		output.Check(os.RemoveAll(pkgdir), "Something went wrong with deleting directory", true)
	}
	if _, err := os.Stat(srcdir); err == nil {
		output.Check(os.RemoveAll(pkgdir), "Something went wrong with deleting directory", true)
	}

	output.Check(os.MkdirAll(pkgdir, 0644), "Something went wrong with creating directory", true)
	output.Check(os.MkdirAll(srcdir, 0644), "Something went wrong with creating directory", true)

	for _, x := range p.Sources {
		output.Info("Downloading " + x + "...")

		basename := filepath.Base(x)
		path := srcdir + "/" + basename

		output.Check(tools.DownloadFile(x, path), "Something went wrong with downloading file", true)

		switch filepath.Ext(basename) {
		case "gz",
			"xz",
			"zip",
			"bz2",
			"lz",
			"zst":
			output.Info("Extracting " + basename + "...")
			output.Check(archiver.Unarchive(path, srcdir), "Something went wrong with extracting an archive", true)
		}
	}
}

func (p *Package) Emerge() {
	if p.Script["type"].(string) == "custom" {
		custom(p)
	} else {
		output.ErrorSimple("Unknown script type")
	}
}
