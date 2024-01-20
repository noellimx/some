package main

type IDoer interface {
	Do(i int) IDoer
}

type Doer struct{}

func (d *Doer) Do(i int) IDoer {
	return d
}

type IWalker interface {
	Do(i string) IWalker
}

func (w *Walker) Do(i string) IWalker {
	some := len(i)
	w.Doer.Do(some)

	return w
}

type Walker struct {
	Doer
}

type GroupWalker struct {
}

func (g *GroupWalker) Do(i string) IWalker {
	return g
}
