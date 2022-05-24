// Concrete iterator
package iterator

type userIterator struct {
	nextIndex int
	users []*user
}

func (u *userIterator) hasNext() bool {
	return u.nextIndex < len(u.users) 
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.nextIndex]
		u.nextIndex++
		return user
	}
	return nil
}