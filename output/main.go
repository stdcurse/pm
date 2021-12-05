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

package output

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Info(str string) {
	fmt.Printf("[*] %s\n", str)
}

func ErrorSimple(str string) {
	color.New(color.FgRed, color.Bold).Print("[ERROR]")
	fmt.Printf(" %s\n", str)

	os.Exit(1)
}

func Check(err error, str string, die ...bool) {
	if err != nil {
		color.New(color.FgRed, color.Bold).Print("[ERROR]")
		fmt.Printf(" %s: %s\n", str, err)

		if len(die) > 0 {
			if die[0] {
				os.Exit(1)
			}
		}
	}
}
