// Client code
package iterator

import (
	"fmt"
	"testing"
)

func Test_Iterator(t *testing.T) {
	user1 := &user{
		name: "a",
		age:  10,
	}

	user2 := &user{
		name: "b",
		age:  20,
	}

	userConllection := &userConllection{
		users: []*user{user1, user2},
	}

	iterator := userConllection.createIterator()

	if !iterator.hasNext() {
		t.Error("Has not next value")
	}

	_user1 := iterator.getNext()

	if fmt.Sprintf("%p", _user1) != fmt.Sprintf("%p", user1) {
		t.Error("Not same memory address")
	}

	if !iterator.hasNext() {
		t.Error("Has not next value")
	}

	_user2 := iterator.getNext()

	if fmt.Sprintf("%p", _user2) != fmt.Sprintf("%p", user2) {
		t.Error("Not same memory address")
	}
}
