package core

import (
	"log"

	pb "github.com/AlexanderKorovayev/microservice/shippy-service-user/proto/user"
	"github.com/dgrijalva/jwt-go"
)

var (
	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	key = []byte("SecretKey")
)

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type TokenService struct {
	Repo Repository
}

// Decode a token string into a token object
func (srv *TokenService) Decode(token string) (*jwt.MapClaims, error) { //(*CustomClaims, error) {
	log.Println(token)
	// Parse the token
	/*
		tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	*/
	tokenType, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	claims := tokenType.Claims.(jwt.MapClaims)
	log.Println(claims)
	// Validate the token and return the custom claims
	// проверка не работает
	/*
		if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
			return claims, nil
		} else {
			return nil, err
		}
	*/
	return &claims, err
}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 150000,
			//Issuer:    "microservice.service.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
