package sigleton

import (
	"errors"
	"testing"
)

func Test_Singleton(t *testing.T) {
	a := GetInstance()
	a.PlusOne()

	b := GetInstance()
	b.PlusOne()

	c := GetInstance()

	if c.MinusOne() != 1 {
		t.Error(errors.New("s count is not 1"))
	}
}