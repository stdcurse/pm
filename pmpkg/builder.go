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
