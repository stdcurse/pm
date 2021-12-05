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

package config

import (
	"github.com/goccy/go-yaml"
	"github.com/stdcurse/pm/output"
	"io/ioutil"
)

type DatabaseEntry struct {
	Name    string
	Version string
	Files   []string
}

type Database []DatabaseEntry

var dbPath string

func NewDatabase(path string) *Database {
	dbPath = path

	content, err := ioutil.ReadFile(path)
	output.Check(err, "Something went wrong with config file", true)

	var db Database
	output.Check(yaml.Unmarshal(content, &db), "Something went wrong with config file content", true)

	return &db
}

func (db *Database) update() {
	c, err := yaml.Marshal(*db)
	output.Check(err, "Something went wrong with database. Unfortunately, it's fatal", true)
	output.Check(ioutil.WriteFile(dbPath, c, 0664), "Something went wrong with database. Unfortunately, it's fatal", true)
}

func (db *Database) Add(name string, version string, files []string) {
	*db = append(*db, DatabaseEntry{
		Name:    name,
		Version: version,
		Files:   files,
	})

	db.update()
}

func (db *Database) Probe(name string) *DatabaseEntry {
	for _, v := range *db {
		if v.Name == name {
			return &v
		}
	}

	return nil
}
