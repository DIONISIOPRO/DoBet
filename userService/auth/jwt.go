package auth

import (
	"errors"
	"github/namuethopro/dobet-user/domain"
	"time"
	"github.com/golang-jwt/jwt"
)

type JWTManager struct {
	PrivateKey []byte
}

func NewJwtManager(PrivateKey []byte) *JWTManager {
	return &JWTManager{
		PrivateKey: PrivateKey,
	}
}


func (manager *JWTManager) GenerateAcessToken(user domain.User) (string, error) {
	claims := &domain.TokenClaims{
		Admin:      user.IsAdmin,
		First_name: user.First_name,
		Last_name:  user.Last_name,
		Phone:      user.Phone_number,
		StandardClaims: jwt.StandardClaims{
			Id: user.User_id,
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(30)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(manager.PrivateKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (manager JWTManager) GenerateRefreshToken(userid string) (string, error) {
	claims := &domain.RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        userid,
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24*7)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(manager.PrivateKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (manager *JWTManager) VerifyToken(incomingtoken string) bool {
	token, err := jwt.ParseWithClaims(incomingtoken, &domain.TokenClaims{}, func(t *jwt.Token) (interface{}, error){
		return manager.PrivateKey, nil
	})
	if err != nil{
		return false
	}
	claims, ok := token.Claims.(*domain.TokenClaims)
	if !ok{
		if err != nil{
			return false
		}
	}
	isTokenExpires := claims.ExpiresAt < time.Now().Unix()
	if isTokenExpires{
		return false
	}
	_, isHMACMethothod := token.Method.(*jwt.SigningMethodHMAC)
	return isHMACMethothod
}
func (manager *JWTManager) IsTokenExpired(incomingtoken string) (bool, error) {
	token, _ := jwt.ParseWithClaims(incomingtoken, &domain.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return manager.PrivateKey, nil
	})
	claims, ok := token.Claims.(*domain.TokenClaims)
	isTokenExpires := claims.ExpiresAt < time.Now().Unix()
	_, isHMACMethothod := token.Method.(*jwt.SigningMethodHMAC)
	if !ok || !isHMACMethothod{
		return false, errors.New("token invalid")
	}
	return isTokenExpires, nil
}

func (manager *JWTManager) ExtractClaimsFromAcessToken(acessToken string) (domain.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(acessToken, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return manager.PrivateKey, nil
	})
	if err != nil{
		return domain.TokenClaims{},errors.New("token invalid")
	}
	claims, ok := token.Claims.(*domain.TokenClaims)
	isTokenExpires := claims.ExpiresAt < time.Now().Unix()
	_, isHMACMethothod := token.Method.(*jwt.SigningMethodHMAC)
	if !ok && isTokenExpires && !isHMACMethothod {
		return domain.TokenClaims{}, errors.New("token is invalid")
	}
	return *claims, nil
}
