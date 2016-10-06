//START1 OMIT
package main

import (
	"github.com/daved/grpcbasic0/idl"
	"golang.org/x/net/context"
)

type UserService struct {
	Users *idl.Users
}

func NewUserService() *UserService {
	return &UserService{
		Users: &idl.Users{
			Users: []*idl.User{
				{Id: 1, Name: "Alice", Age: 21, Fortune: fortunes.get(1)},
				{Id: 2, Name: "Bob", Age: 30, Fortune: fortunes.get(2)},
				{Id: 3, Name: "Charlie", Age: 25, Fortune: fortunes.get(3)},
				{Id: 4, Name: "Dana", Age: 18, Fortune: fortunes.get(4)},
			},
		},
	}
}

//END1 OMIT

//START2 OMIT
func (s *UserService) RecordUser(ctx context.Context, req *idl.UserRecordReq) (*idl.User, error) {
	// ...
	return nil, nil //OMIT
}

func (s *UserService) GetUser(ctx context.Context, req *idl.UserGetReq) (*idl.User, error) {
	// ...
	return nil, nil //OMIT
}

// GetUsers ...
func (s *UserService) GetUsers(ctx context.Context, req *idl.UsersGetReq) (*idl.Users, error) {
	// ...
	return nil, nil //OMIT
}

//END2 OMIT
