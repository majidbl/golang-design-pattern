package user

import "design_pattern/domain/iterator"
type Collection struct {
	Users []*iterator.User
}

func (u *Collection) CreateIterator() iterator.Iterator {
	return &Iterator{
		Users: u.Users,
	}
}
