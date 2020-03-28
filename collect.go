package iterator

type CollectFn func(interface{})

func collect(it Iterator, collectFn CollectFn) {
	for elem, ok := it.Next(); ok; elem, ok = it.Next() {
		collectFn(elem)
	}
}
