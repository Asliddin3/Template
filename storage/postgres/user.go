package postgres

import (
	pb "github.com/Asliddin3/Template/genproto"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	userResp := pb.User{}
	err := r.db.QueryRow(`insert into users (name, last_name) values($1, $2) returning id, name, last_name`, user.Name, user.LastName).Scan(&userResp.Id, &userResp.Name, &userResp.LastName)
	if err != nil {
		return &pb.User{}, err
	}

	return &userResp, nil
}
