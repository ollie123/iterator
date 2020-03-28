package iterator

type MapFn func(interface{}) interface{}

type mapper struct {
	Iterator
	mapFn MapFn
}

func newMapper(it Iterator, mapFn MapFn) *mapper {
	return &mapper{
		Iterator: it,
		mapFn:    mapFn,
	}
}

func (it *mapper) Next() (interface{}, bool) {
	elem, ok := it.Iterator.Next()
	if !ok {
		return nil, false
	}
	return it.mapFn(elem), true
}

func (it *mapper) Collect(collectFn CollectFn) {
	collect(it, collectFn)
}

func (it *mapper) Filter(filterFn FilterFn) Iterator {
	return newFilterer(it, filterFn)
}

func (it *mapper) Map(mapFn MapFn) Iterator {
	return newMapper(it, mapFn)
}
