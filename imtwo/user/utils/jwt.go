package utils

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

var JwtKey=[]byte("penguinGo")

func CreateToken(userId int64) string {
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"iss":		"penguin",
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour).Unix(),
		"sub":      "user",
		"user_id":	 userId,
	})

	tokenString,err:= token.SignedString(JwtKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

type AuthToken struct {
	Token string
}

func (a *AuthToken) GetRequestMetadata(ctx context.Context,uri ...string)(map[string]string,error) {
	return map[string]string{
		"authorization":a.Token,
	},nil


}

func (a *AuthToken) RequireTransportSecurity() bool {
	return false
}

type Claims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}


//从context 的metadata 中，取出token

func getTokenFromContext(ctx context.Context)(string,error)  {
	//md 的类型是 type MD map[string][]string
	md,ok:= metadata.FromIncomingContext(ctx)
	if !ok {
		return "",errors.New("ErrNoMetadataInContext")
	}

	token,ok := md["authorization"]
	if !ok || len(token)==0 {
		return "",errors.New("ErrNoAuthorizationInMetadata")
	}
	log.Println(token)

	return token[0],nil

}

func CheckAuth(ctx context.Context) int64 {
	tokenStr ,err:= getTokenFromContext(ctx)
	if err != nil {
		panic("get token from context error")
	}

	var clientClaims Claims
	token ,err:= jwt.ParseWithClaims(tokenStr,&clientClaims,func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			panic("ErrInvalidAlgorithm")
		}
		return JwtKey, nil
	})
	if err != nil {
		panic(err)
	}

	if !token.Valid {
		panic("ErrInvalidToken")
	}


	return clientClaims.UserId
}
