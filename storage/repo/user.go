package repo

import (
    pb "github.com/Asliddin3/Template/genproto"
)

//UserStorageI ...
type UserStorageI interface {
  Create(*pb.User) (*pb.User, error)
	DeleteUser(*pb.RequesUser)(*pb.Users,error)
	GetUsers(*pb.Empty)(*pb.Users,error)
	UpdateUser(*pb.User)(*pb.User,error)
	GetUser(*pb.RequesUser)(*pb.User,error)
}
