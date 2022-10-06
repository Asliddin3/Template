package service

import (
	"context"

	pb "github.com/Asliddin3/Template/genproto"
	l "github.com/Asliddin3/Template/pkg/logger"
	"github.com/Asliddin3/Template/storage"
	"github.com/jmoiron/sqlx"
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

func (s *UserService) GetUsers(ctx context.Context, req *pb.Empty) (*pb.Users, error) {
	users, err := s.storage.User().GetUsers(req)
	if err != nil {
		s.logger.Error("error geting", l.Any("error geting users", err))
		return &pb.Users{}, status.Error(codes.Internal, "something went wrong while geting")
	}
	return users, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.RequesUser) (*pb.User, error) {
	user, err := s.storage.User().GetUser(req)
	if err != nil {
		s.logger.Error("error geting", l.Any("error geting user", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrong while geting")
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().UpdateUser(req)
	if err != nil {
		s.logger.Error("error update", l.Any("error updating user", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrong please check your input")
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.RequesUser) (*pb.Users, error) {
	users, err := s.storage.User().DeleteUser(req)
	if err != nil {
		s.logger.Error("error delete", l.Any("error deleting user", err))
		return &pb.Users{}, status.Error(codes.Internal, "something went wrong please check your input")
	}
	return users, nil
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("error insert", l.Any("error insert user", err))
		return &pb.User{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}

	return user, nil
}
