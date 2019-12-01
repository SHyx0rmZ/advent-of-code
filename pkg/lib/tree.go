package lib

type Tree interface {
	Walk(func(Tree))
	Value() interface{}
}

func TreeDo(f func(interface{}), t Tree) {
	t.Walk(func(tree Tree) {
		f(tree.Value())
	})
}

func Walk(t Tree, f func(interface{})) {
	f(t.Value())
	t.Walk(func(tree Tree) {
		Walk(tree, f)
	})
}
