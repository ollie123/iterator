package iterator

type GetterFn func(int) interface{}

type SliceSource struct {
	i, len   int
	getterFn GetterFn
}

func NewSliceSource(len int, getterFn GetterFn) *SliceSource {
	return &SliceSource{
		len:      len,
		getterFn: getterFn,
	}
}

func (s *SliceSource) Next() (interface{}, bool) {
	if s.i < s.len {
		elem := s.getterFn(s.i)
		s.i++
		return elem, true
	}
	return nil, false
}

func (s *SliceSource) Reset() {
	s.i = 0
}
