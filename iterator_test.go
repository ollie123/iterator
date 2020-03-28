package iterator

import (
	"strings"
	"testing"
)

type Artist struct {
	Country string
	Name    string
}

func TestIterator(t *testing.T) {
	artists := []*Artist{
		&Artist{"Japan", "Hatsune Miku"},
		&Artist{"Japan", "Kagamine Rin"},
		&Artist{"Japan", "Kagamine Len"},
		&Artist{"Japan", "Megurine Luka"},
		&Artist{"China", "Luo Tianyi"},
		&Artist{"Japan", "Yuzuki Yukari"},
		&Artist{"Japan", "Otomachi Una"},
	}
	source := NewSliceSource(len(artists), func(i int) interface{} {
		return artists[i]
	})
	var name string
	NewIterator(source).Filter(func(elem interface{}) bool {
		return elem.(*Artist).Country == "China"
	}).Map(func(elem interface{}) interface{} {
		return elem.(*Artist).Name
	}).Map(func(elem interface{}) interface{} {
		names := strings.Split(elem.(string), " ")
		return names[len(names)-1]
	}).Collect(func(elem interface{}) {
		name = elem.(string)
	})
	if name != "Tianyi" {
		t.Errorf("name = %s; want Tianyi", name)
	}
}
