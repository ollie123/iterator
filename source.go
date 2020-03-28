package iterator

type Source interface {
	Next() (interface{}, bool)
}