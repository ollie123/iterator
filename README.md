# Iterators in Golang

Just playing around. You're probably better off sticking with for loops in actual Go code.

```go
package main

import (
	"fmt"
	"math"

	"github.com/ollie123/iterator"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	source := iterator.NewSliceSource(len(numbers), func(i int) interface{} {
		return numbers[i]
	})
	iterator.NewIterator(source).Filter(func(elem interface{}) bool {
		return elem.(int)%2 == 0
	}).Map(func(elem interface{}) interface{} {
		return math.Pow(float64(elem.(int)), 2)
	}).Collect(func(elem interface{}) {
		fmt.Println(elem)
	})
}
```

```
4
16
36
64
```