package linq

import (
	"fmt"
	golinq "github.com/ahmetb/go-linq/v3"
)

func WorkFlow() {
	m := make(map[int]bool)
	m[10] = true
	fmt.Println(golinq.From(m).Results())

	ages := []int{21, 46, 46, 55, 17, 21, 55, 55}

	var distinctAges []int
	golinq.From(ages).
		OrderBy(
			func(item interface{}) interface{} { return item },
		).
		Distinct().
		ToSlice(&distinctAges)

	fmt.Println(distinctAges)

	query := golinq.From([]int{1, 2, 3, 4, 5}).Where(func(i interface{}) bool {
		return i.(int) <= 3
	})

	next := query.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		fmt.Println(item)
	}

	var ss []int
	golinq.From([]int{1, 2, 3, 4, 5}).Where(func(i interface{}) bool {
		return i.(int) <= 3
	}).ToSlice(&ss)
}
