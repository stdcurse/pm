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
	"io/ioutil"
	"os"
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
