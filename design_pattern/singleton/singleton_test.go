package sigleton

import (
	"errors"
	"testing"
)

func Test_Singleton(t *testing.T) {
	input := []int {1, 1}
	output := 2

	a1 := GetInstance_1()
	a1.AddCount(input[0])
	a2 := GetInstance_1()
	a2.AddCount(input[1])

	b1 := GetInstance_2()
	b1.AddCount(input[0])
	b2 := GetInstance_2()
	b2.AddCount(input[1])

	c1 := GetInstance_3()
	c1.AddCount(input[0])
	c2 := GetInstance_3()
	c2.AddCount(input[1])

	if a2.GetCount() != output {
		t.Error(errors.New("Wrong output. - Case 1"))
	}

	if b2.GetCount() != output {
		t.Error(errors.New("Wrong output. - Case 2"))
	}

	if c2.GetCount() != output {
		t.Error(errors.New("Wrong output. - Case 3"))
	}
}