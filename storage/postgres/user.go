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
func (r *userRepo) GetUsers(req *pb.Empty) (*pb.Users, error) {
	usersRepo := pb.Users{}
	rows, err := r.db.Query(`
	select id,name,last_name from users`)
	for rows.Next() {
		var userRepo pb.User
		err = rows.Scan(&userRepo.Id, &userRepo.Name, &userRepo.LastName)
		if err != nil {
			return &pb.Users{}, err
		}
		usersRepo.Users = append(usersRepo.Users, &userRepo)
	}
	return &usersRepo, nil
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	userResp := pb.User{}
	err := r.db.QueryRow(`insert into users (name, last_name) values($1, $2) returning id, name, last_name`, user.Name, user.LastName).Scan(&userResp.Id, &userResp.Name, &userResp.LastName)
	if err != nil {
		return &pb.User{}, err
	}

	return &userResp, nil
}

func (r *userRepo) UpdateUser(user *pb.User) (*pb.User, error) {
	userRepo := pb.User{}
	_, err := r.db.Exec(`update users set name=$1,
	last_name=$2 where id=$3`, user.Name, user.LastName, user.Id)
	if err != nil {
		return &pb.User{}, err
	}
	return &userRepo, nil
}

func (r *userRepo) GetUser(req *pb.RequesUser) (*pb.User, error) {
	user := pb.User{}
	err := r.db.QueryRow(`select id,name,last_name
	from users where id=$1`, req.Id).Scan(&user.Id, &user.Name, &user.LastName)
	if err != nil {
		return &pb.User{}, err
	}
	return &user, nil
}

func (r *userRepo) DeleteUser(user *pb.RequesUser) (*pb.Users, error) {
	usersRepo := pb.Users{}
	_, err := r.db.Exec(`delet from users where id=$1`, user.Id)
	if err != nil {
		return &pb.Users{}, err
	}
	rows, err := r.db.Query(`
	select id,name,last_name from users`)
	for rows.Next() {
		var userRepo pb.User
		err = rows.Scan(&userRepo.Id, &userRepo.Name, &userRepo.LastName)
		if err != nil {
			return &pb.Users{}, err
		}
		usersRepo.Users = append(usersRepo.Users, &userRepo)
	}
	return &usersRepo, nil
}
