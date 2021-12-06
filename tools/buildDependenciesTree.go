package tools

import (
	"github.com/stdcurse/pm/pmpkg"
	"github.com/stevenle/topsort"
)

func BuildDependenciesTree(pkgs []*pmpkg.Package) []*pmpkg.Package {
	graph := topsort.NewGraph()

	graph.AddNode("A")

	var walk func(parent string, pkg *pmpkg.Package)
	walk = func(parent string, pkg *pmpkg.Package) {
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

	var res []*pmpkg.Package
	for _, x := range r[:1] {
		for _, v := range pkgs {
			if v.Name == x {
				res = append(res, v)
			}
		}
	}

	return res
}
