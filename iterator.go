package iterator

type Iterator interface {
	Source
	Collect(CollectFn)
	Filter(FilterFn) Iterator
	Map(MapFn) Iterator
}

type iterator struct {
	Source
}

func NewIterator(source Source) Iterator {
	return &iterator{
		Source: source,
	}
}

func (it *iterator) Collect(collectFn CollectFn) {
	collect(it, collectFn)
}

func (it *iterator) Filter(filterFn FilterFn) Iterator {
	return newFilterer(it, filterFn)
}

func (it *iterator) Map(mapFn MapFn) Iterator {
	return newMapper(it, mapFn)
}