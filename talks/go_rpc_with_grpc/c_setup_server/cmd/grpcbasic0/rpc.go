//START1 OMIT
package main

import (
	"context" //OMIT
	//OMIT
	"github.com/daved/grpcbasic0/pb"
)

type UserService struct {
	UsersResp *pb.UsersResp
}

func NewUserService() *UserService {
	return &UserService{
		UsersResp: &pb.UsersResp{
			Users: []*pb.UserResp{
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
func (s *UserService) NewUser(ctx context.Context, req *pb.NewUserReq) (*pb.UserResp, error) {
	// ...
	return nil, nil //OMIT
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.UserResp, error) {
	// ...
	return nil, nil //OMIT
}

// GetUsers ...
func (s *UserService) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.UsersResp, error) {
	// ...
	return nil, nil //OMIT
}

//END2 OMIT

type forts struct{}

func (f *forts) get(int) string { return "" }

var fortunes = &forts{}
