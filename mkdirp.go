package mkdirp

import (
	"os"
	"path"
	"regexp"
)

var sep = regexp.MustCompile(`(\}/){1}|([\{\},]){1}`)

var (
	startSubtree = "{"
	eachSubtree  = "}/"
	endSubtree   = "}"
	sibling      = ","
)

type TreeNode interface {
	newNode() *treeNode
	setName(string)
	getName() string
	getParent() *treeNode
	getChildren() []*treeNode
}

type treeNode struct {
	name     string
	children []*treeNode
	parent   *treeNode
	root     *treeRoot
}

type TreeRoot interface {
	TreeNode
	addNode(note *treeNode)
	GetPaths() []string
}

type treeRoot struct {
	treeNode
	nodes []*treeNode
}

func (t *treeRoot) addNode(node *treeNode) {
	t.nodes = append(t.nodes, node)
}

func (t *treeRoot) getLeaves() (leaves []*treeNode) {
	for _, n := range t.nodes {
		if len(n.children) == 0 {
			leaves = append(leaves, n)
		}
	}
	return
}

func (t *treeRoot) GetPaths() (paths []string) {
	for _, n := range t.nodes {
		if len(n.children) == 0 { // node is leaf
			var p string
			for {
				p = path.Join(n.name, p)
				if n.parent == nil {
					break
				}
				n = n.parent
			}
			paths = append(paths, p)
		}
	}
	return
}

func newTree() (tree *treeRoot) {
	tree = &treeRoot{treeNode{children: []*treeNode{}}, []*treeNode{}}
	tree.root = tree
	tree.nodes = append(tree.nodes, &tree.treeNode)
	return
}

func (t *treeNode) newNode() *treeNode {
	node := &treeNode{children: []*treeNode{}, parent: t, root: t.root}
	t.children = append(t.children, node)
	t.root.addNode(node)
	return node
}

func (t *treeNode) getName() string {
	return t.name
}

func (t *treeNode) setName(name string) {
	t.name = name
}

func (t *treeNode) getParent() *treeNode {
	return t.parent
}

func (t *treeNode) getChildren() []*treeNode {
	return t.children
}

func build(command string, tree TreeNode) {
	if len(command) == 0 {
		return
	}
	l := sep.FindStringIndex(command)
	var n TreeNode
	if tree.getName() != "" {
		n = tree.newNode()
	} else {
		n = tree
	}
	if l != nil {
		n.setName(command[:l[0]])
		next := command[l[1]:]
		if s := string(command[l[0]:l[1]]); s == startSubtree {
			build(next, n)
		} else if s == eachSubtree {
			for _, b := range tree.getChildren() {
				build(next, b)
			}
		} else if s == endSubtree {
			build(next, tree.getParent())
		} else { // sibling
			build(next, tree)
		}
	} else {
		n.setName(command)
	}
}

func MkTree(command string) TreeRoot {
	tree := newTree()
	build(command, tree)
	return tree
}

func Mk(command string, perm os.FileMode) (err error) {
	paths := MkTree(command).GetPaths()

	for _, p := range paths {
		err = os.MkdirAll(p, perm)
		if err != nil {
			return
		}
	}
	return
}
