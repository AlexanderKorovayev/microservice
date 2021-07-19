package core

import (
	"context"
	"log"

	pb "github.com/AlexanderKorovayev/microservice/shippy-service-user/proto/user"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `sql:"id"`
	Name     string `sql:"name"`
	Email    string `sql:"email"`
	Company  string `sql:"company"`
	Password string `sql:"password"`
}

type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

func MarshalUserCollection(users []*pb.User) []*User {
	u := make([]*User, len(users))
	for _, val := range users {
		u = append(u, MarshalUser(val))
	}
	return u
}

func MarshalUser(user *pb.User) *User {
	return &User{
		ID:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: user.Password,
	}
}

func UnmarshalUserCollection(users []*User) []*pb.User {
	u := make([]*pb.User, 0)
	for _, val := range users {
		u = append(u, UnmarshalUser(val))
	}
	return u
}

func UnmarshalUser(user *User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: user.Password,
	}
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]*User, error) {
	users1 := make([]*User, 0)
	var users []User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, el := range users {
		users1 = append(users1, &el)
	}
	return users1, nil
}

func (r *PostgresRepository) Get(ctx context.Context, id string) (*User, error) {
	var user User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		log.Println("Error for get user")
		log.Println(result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func (r *PostgresRepository) Create(ctx context.Context, user *User) error {
	user.ID = uuid.NewV4().String()
	result := r.db.Create(user)
	if result.Error != nil {
		log.Println("Error for get user")
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	//var user *User
	var user User
	result := r.db.First(&user, "email = ?", email)
	if result.Error != nil {
		log.Println("Error for get by email")
		log.Println(result.Error)
		return nil, result.Error
	}

	return &user, nil
}
