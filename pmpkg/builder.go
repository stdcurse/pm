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
	"fmt"
	"github.com/mibk/shellexec"
	"github.com/stdcurse/pm/output"
	"os"
	"os/exec"
)

func setupenv(p *Package, cmd *exec.Cmd) {
	for k, v := range c.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	pkgdir := fmt.Sprintf("%s/%s/pkgdir", c.Portdir, p.Name)
	srcdir := fmt.Sprintf("%s/%s/srcdir", c.Portdir, p.Name)

	cmd.Env = append(cmd.Env, "pkgdir="+pkgdir)
	cmd.Env = append(cmd.Env, "srcdir="+srcdir)
}

func custom(p *Package) {
	cmd, err := shellexec.Command(fmt.Sprintf("/bin/sh -c 'source /etc/profile; %s'", p.Script["instructions"].(string)))
	output.Check(err, "Something is wrong with build instructions", true)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	setupenv(p, cmd)

	srcdir := fmt.Sprintf("%s/%s/srcdir", c.Portdir, p.Name)

	res := cmd.Run()

	if _, err := os.Stat(srcdir); err == nil {
		output.Check(os.RemoveAll(srcdir), "Something went wrong with deleting directory", true)
	}

	if res != nil {
		fmt.Println(res)
		output.ErrorSimple("Process exited with non-zero exit-code, can't operate anymore")
	}
}
