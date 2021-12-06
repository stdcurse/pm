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

package emerge

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/stdcurse/pm/config"
	"github.com/stdcurse/pm/output"
	"github.com/stdcurse/pm/pmpkg"
	"github.com/stdcurse/pm/tools"
	"github.com/urfave/cli/v2"
)

func Command(c *cli.Context) error {
	tools.NeedRoot()
	cfg := config.NewConfig(c.String("config"))
	db := config.NewDatabase(cfg.Database)

	if c.Args().Len() == 0 {
		output.ErrorSimple("You must specify packages names")
	}

	var pkgs []*pmpkg.Package

	for _, pkgName := range c.Args().Slice() {
		p := pmpkg.NewPackage(pkgName, db, cfg)

		if p == nil {
			output.ErrorSimple(fmt.Sprintf("Package %s is not found", pkgName))
		}

		pkgs = append(pkgs, p)
	}

	tree := pmpkg.BuildDependenciesTree(pkgs)
	if tree == nil {
		output.ErrorSimple("Cyclic dependence found, can't operate")
	}

	output.Info("The following new packages will be installed:")
	for _, x := range tree {
		output.Info("- " + x.Name)
	}

	for _, x := range tree {
		fmt.Println()
		output.Info("Emerging " + x.Name + "...")
		x.Fetch()
		x.Emerge()

		var files []string
		pkgdir := fmt.Sprintf("%s/%s/pkgdir", cfg.Portdir, x.Name)

		output.Check(filepath.Walk(pkgdir, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path == pkgdir {
				return nil
			}
			files = append(files, strings.Replace(path, pkgdir, "", 1))
			return nil
		}), "Something went wrong with walking through pkgdir", true)

		output.Info("Merging pkgdir with filesystem...")
		output.Check(tools.BetterRenameFilesRecursively(pkgdir, "/"), "Something went wrong with merging files", true)

		if _, err := os.Stat(pkgdir); err == nil {
			output.Check(os.RemoveAll(pkgdir), "Something went wrong with deleting directory", true)
		}

		db.Add(x.Name, x.Version, x.Release, files)

		output.Info("Package " + x.Name + " was successfully emerged!")
	}

	return nil
}
