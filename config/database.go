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
