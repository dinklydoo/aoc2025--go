package day08

type DSU struct {
	dsu  []int
	size []int
	sets int
}

func (ds *DSU) find(n int) int {
	if ds.dsu[n] == n {
		return n
	}
	ds.dsu[n] = ds.find(ds.dsu[n])
	return ds.dsu[n]
}

func (ds *DSU) union(a, b int) {
	fa := ds.find(a)
	fb := ds.find(b)

	if fa == fb {
		return
	}

	if fa < fb {
		ds.dsu[fb] = fa
		ds.size[fa] += ds.size[fb]
	} else {
		ds.dsu[fa] = fb
		ds.size[fb] += ds.size[fa]
	}
	ds.sets--
}

func initDSU(n int) *DSU {
	dsu := DSU{
		make([]int, n),
		make([]int, n),
		n,
	}

	for i := 0; i < n; i++ {
		dsu.dsu[i] = i
	}

	for i := 0; i < n; i++ {
		dsu.size[i] = 1
	}

	return &dsu
}
