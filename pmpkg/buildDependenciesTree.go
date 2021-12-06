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

import "github.com/stevenle/topsort"

func BuildDependenciesTree(pkgs []*Package) []*Package {
	graph := topsort.NewGraph()

	graph.AddNode("A")

	var walk func(parent string, pkg *Package)
	walk = func(parent string, pkg *Package) {
		graph.AddEdge(parent, pkg.Name)

		for _, x := range pkg.Dependencies {
			walk(pkg.Name, x)
		}
	}

	for _, x := range pkgs {
		graph.AddEdge("A", x.Name)

		for _, v := range x.Dependencies {
			walk(x.Name, v)
		}
	}

	r, err := graph.TopSort("A")

	if err != nil {
		return nil
	}

	var res []*Package
	for _, x := range r[:len(r)-1] {
		for _, v := range pkgs {
			if v.Name == x {
				res = append(res, v)
			}
		}
	}

	return res
}
