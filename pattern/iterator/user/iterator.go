package user

import "design_pattern/domain/iterator"
type Iterator struct {
	Index int
	Users []*iterator.User
}

func (u *Iterator) HasNext() bool {
	if u.Index < len(u.Users) {
		return true
	}
	return false

}
func (u *Iterator) GetNext() *iterator.User {
	if u.HasNext() {
		user := u.Users[u.Index]
		u.Index++
		return user
	}
	return nil
}

