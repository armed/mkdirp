package mkdirp

import (
	"regexp"
)

var sep = regexp.MustCompile(`(\}/){1}|([\{\},]){1}`)

var (
	startSubtree = "{"
	eachSubtree  = "}/"
	endSubtree   = "}"
	sibling      = ","
)

type dirTree struct {
	name     string
	branches []*dirTree
	parent   *dirTree
}

func (t *dirTree) newBranch() (branch *dirTree) {
	branch = &dirTree{branches: []*dirTree{}, parent: t}
	t.branches = append(t.branches, branch)
	return
}

func build(data string, tree *dirTree) {
	if len(data) <= 1 {
		return
	}
	l := sep.FindStringIndex(data)
	if l != nil {
		br := tree.newBranch()
		br.name = data[:l[0]]
		next := data[l[1]:]
		if s := string(data[l[0]:l[1]]); s == startSubtree {
			build(next, br)
		} else if s == eachSubtree {
			for _, b := range tree.branches {
				build(next, b)
			}
		} else if s == endSubtree {
			build(next, tree.parent)
		} else { // sibling
			build(next, tree)
		}
	} else {
		br := tree.newBranch()
		br.name = data
	}
}

func mkTree(command string) *dirTree {
	tree := &dirTree{branches: []*dirTree{}}
	build(command, tree)
	return tree.branches[0]
}
