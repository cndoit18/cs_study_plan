package flyweight

func New(name string) *tree {
	return &tree{name: name}
}

type tree struct {
	name string
}

func (f *tree) Name() string {
	return f.name
}

type Trees struct {
	Trees []*tree
}

func GenerateTrees(n int) *Trees {
	t := New("tree")
	ts := &Trees{
		Trees: make([]*tree, 0, n),
	}
	for i := 0; i < n; i++ {
		ts.Trees = append(ts.Trees, t)
	}

	return ts
}
