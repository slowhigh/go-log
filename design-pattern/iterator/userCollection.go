// Concrete collection
package iterator

type userConllection struct {
	users []*user
}

func (u *userConllection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}
