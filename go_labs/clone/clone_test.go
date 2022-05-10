package clone

import "testing"

type item struct {
	value int
}

func Test_Clone(t *testing.T) {
	item1 := &item{value: 1}
	var item2 *item
	item2 = item1
	item1.value = 3
	t.Log(item1)
	t.Logf("item1 %p", item1)
	t.Log(item2)
	t.Logf("item2 %p", item2)
}
