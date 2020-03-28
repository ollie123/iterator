package iterator

type FilterFn func(interface{}) bool

type filterer struct {
	Iterator
	filterFn FilterFn
}

func newFilterer(it Iterator, filterFn FilterFn) *filterer {
	return &filterer{
		Iterator: it,
		filterFn: filterFn,
	}
}

func (it *filterer) Next() (interface{}, bool) {
	for {
		elem, ok := it.Iterator.Next()
		if !ok {
			return nil, false
		}
		if it.filterFn(elem) {
			return elem, true
		}
	}
}

func (it *filterer) Collect(collectFn CollectFn) {
	collect(it, collectFn)
}

func (it *filterer) Filter(filterFn FilterFn) Iterator {
	return newFilterer(it, filterFn)
}

func (it *filterer) Map(mapFn MapFn) Iterator {
	return newMapper(it, mapFn)
}