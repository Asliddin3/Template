package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "github.com/Asliddin3/Template/genproto"
)

type userRepo struct{
	db *sqlx.DB
}

func (r *userRepo) Create(user *pb.User) (*pb.User,error){
	userRepo :=pb.User{}
	err:=r.db.QueryRow(`insert into users(name,last_name) values($1,$2)
	returning id,name,last_name`,user.Name,user.LastName).Scan(&userRepo.Id,&userRepo.Name,&userRepo.LastName)
	if err != nil {
		return &pb.User{}, err
	}
	return &userRepo,nil
}