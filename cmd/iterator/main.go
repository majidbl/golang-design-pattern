package main

import (
	"fmt"

	iterators "design_pattern/pattern/iterator/user"
	iteratorDomain "design_pattern/domain/iterator"
)

func main() {

	user1 := &iteratorDomain.User{
		Name: "a",
		Age:  30,
	}
	user2 := &iteratorDomain.User{
		Name: "b",
		Age:  20,
	}

	userCollection := &iterators.Collection{
		Users: []*iteratorDomain.User{user1, user2},
	}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}