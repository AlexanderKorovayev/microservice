package core

import (
	"context"

	pb "github.com/AlexanderKorovayev/microservice/shippy-service-user/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type Handler struct {
	repository Repository
	//tokenService authable
}

func (s *Handler) Get(ctx context.Context, req *pb.User) (*pb.Response, error) {
	result, err := s.repository.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	user := UnmarshalUser(result)

	return &pb.Response{User: user}, nil
}

func (s *Handler) GetAll(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	results, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	users := UnmarshalUserCollection(results)

	return &pb.Response{Users: users}, nil
}

/*
func (s *Handler) Auth(ctx context.Context, req *pb.User) (*pb.Token, error) {
	user, err := s.repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	token, err := s.tokenService.Encode(req)
	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}
*/

func (s *Handler) Create(ctx context.Context, req *pb.User) (*pb.Response, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashedPass)
	if err := s.repository.Create(ctx, MarshalUser(req)); err != nil {
		return nil, err
	}

	// Strip the password back out, so's we're not returning it
	req.Password = ""
	res.User = req

	return &pb.Response{User: req}, nil
}

/*
func (s *Handler) ValidateToken(ctx context.Context, req *pb.Token) (*pb.Token, error) {
	claims, err := s.tokenService.Decode(req.Token)
	if err != nil {
		return nil, err
	}

	if claims.User.Id == "" {
		return nil, errors.New("invalid user")
	}

	return &pb.Token{Valid: true}, nil
}
*/
