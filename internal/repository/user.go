package repository

import (
	"context"
	"fmt"

	"github.com/mnsv1511/kasikorn-line-assignment/ent"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/usergreetings"
	"github.com/mnsv1511/kasikorn-line-assignment/ent/users"
)

func (r *RepositoryImpl) GetUser(userId int) (*ent.Users, error) {
	userData, err := r.clientDB.Users.Query().
		Where(users.ID(userId)).
		Only(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return userData, nil
}

func (r *RepositoryImpl) GetUserGreeting(userId int) (*ent.UserGreetings, error) {
	userData, err := r.clientDB.UserGreetings.Query().
		Where(usergreetings.UserIDEQ(userId)).
		WithUsers().
		Only(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return userData, nil
}
