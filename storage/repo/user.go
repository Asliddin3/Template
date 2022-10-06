package repo

import (
    pb "github.com/Asliddin3/Template/genproto"
)

//UserStorageI ...
type UserStorageI interface {
    Create(*pb.User) (*pb.User, error)
}
