// Concrete iterator
package iterator

type userIterator struct {
	nextIndex int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.nextIndex < len(u.users) {
		return true
	}

	return false
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.nextIndex]
		u.nextIndex++
		return user
	}
	return nil
}