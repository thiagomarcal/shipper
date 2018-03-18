package main

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	pb "github.com/thiagomarcal/shipper/user-service/proto/user"
)

var (
	key = []byte("mySuperSecretKeyLol")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {

	log.Println("Decoding Token ...")

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {

		log.Println("Problem parsing", err)

		return nil, err
	}

	if token.Valid {
		claims, ok := token.Claims.(*CustomClaims)
		if ok {
			return claims, nil
		} else {
			return nil, errors.New("Problem with claims")
		}
	}

	return nil, err

}

func (srv *TokenService) Encode(user *pb.User) (string, error) {

	log.Println("Encoding Token ...")

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "shipper.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
