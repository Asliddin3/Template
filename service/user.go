package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	pb "github.com/Asliddin3/Template/genproto"
	l "github.com/Asliddin3/Template/pkg/logger"
	"github.com/Asliddin3/Template/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
    user, err := s.storage.User().Create(req)
    if err != nil {
        s.logger.Error("error insert", l.Any("error insert user", err))
        return &pb.User{}, status.Error(codes.Internal, "something went wrong, please check user info")
    }

	return user, nil
}
