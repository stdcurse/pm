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
