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
