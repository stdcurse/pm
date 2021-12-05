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

package install

import (
	"fmt"
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

	for _, pkgName := range c.Args().Slice() {
		p := pmpkg.NewPackage(pkgName, db, cfg)

		if p == nil {
			output.ErrorSimple(fmt.Sprintf("Package %s is not found", pkgName))
		}

		fmt.Println(*p)
	}

	return nil
}
