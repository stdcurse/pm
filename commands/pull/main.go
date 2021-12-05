package pull

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/stdcurse/pm/config"
	"github.com/stdcurse/pm/output"
	"github.com/stdcurse/pm/tools"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Command(c *cli.Context) error {
	tools.NeedRoot()
	cfg := config.NewConfig(c.String("config"))

	tmp, err := ioutil.TempDir("", "pm-pull")
	output.Check(err, "Something went wrong with creating temporary directory", true)

	for _, v := range cfg.Repos {
		output.Info(fmt.Sprintf("Pulling %s...", v.Url))

		filename := tmp + "/" + filepath.Base(v.Url)
		output.Check(tools.DownloadFile(v.Url, filename), "Something went wrong with downloading file", true)
		output.Check(archiver.Unarchive(filename, tmp), "Something went wrong with extracting an archive", true)

		tools.RenameDirectoryRecursively(tmp+"/"+v.Path, cfg.Portdir)
	}

	output.Info("Done")
	output.Check(os.RemoveAll(tmp), "Something went wrong with removing temporary directory", true)

	return nil
}
