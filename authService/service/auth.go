package service

import (
	"errors"

	"github.com/dionisiopro/dobet-auth/domain"
)

type (
	Publisher interface {
		Publish(name string, event domain.Event) error
	}
	Authrepo interface {
		Login(phone string) (domain.User, error)
		SignUp(user domain.User) (string, error)
		DeleteUser(userid string) error
		UpdateUser(userId string, user domain.User) error
		AddRefreshToken(userid, refreshtoken string) error
		GetRefreshTokens(userid string) ([]string, error)
	}
	PasswordVerifier interface {
		VerifyPassword(password, hash string) bool
	}

	authService struct {
		PasswordVerifier PasswordVerifier
		authRepo         Authrepo
		jwtmanager       jwtManager
		Publisher        Publisher
	}

	jwtManager interface {
		GenerateAcessToken(user domain.User) (string, error)
		GenerateRefreshToken(userid string) (string, error)
		VerifyToken(incomingtoken string) bool
		IsTokenExpired(token string) (bool, error)
		ExtractClaimsFromAcessToken(acessToken string) (domain.TokenClaims, error)
	}
)

func NewAuthService(Authrepo Authrepo, jwtmanager jwtManager,Publisher Publisher) *authService {
	service := &authService{
		PasswordVerifier: PasswordHandler{},
		authRepo:         Authrepo,
		jwtmanager:       jwtmanager,
		Publisher:        Publisher,
	}
	return service
}
func (service *authService) CreateUser(user domain.User) error {
	_, err := service.authRepo.SignUp(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *authService) DeleteUser(userId string) error {
	err := service.authRepo.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
func (service *authService) UpdateUser(userid string, user domain.User) error {
	err := service.authRepo.UpdateUser(userid, user)
	if err != nil {
		return err
	}
	return nil
}

func (service *authService) Login(user domain.LoginDetails) (token, refreshToken string, err error) {
	if user.Phone == "" || user.Password == "" {
		return "", "", errors.New("invalid user details")
	}
	localuser, err := service.authRepo.Login(user.Phone)
	if err != nil {
		return "", "", err
	}
	ok := service.PasswordVerifier.VerifyPassword(user.Password, localuser.Hashed_password)
	if !ok {
		return "", "", errors.New("user invalid")
	}
	token, err = service.jwtmanager.GenerateAcessToken(localuser)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = service.jwtmanager.GenerateRefreshToken(localuser.User_id)
	if err != nil {
		return "", "", err
	}
	if err != nil {
		return "", "", err
	}
	err = service.authRepo.AddRefreshToken(user.Phone, refreshToken)
	service.publishLoginEvent(localuser.User_id)
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}

func (service *authService) Logout(token string) error {
	ok := service.jwtmanager.VerifyToken(token)
	if !ok {
		return errors.New("token invalid")
	}
	claims, err := service.jwtmanager.ExtractClaimsFromAcessToken(token)
	if err != nil {
		return err
	}
	userId := claims.Id
	err = service.publishLogOutEvent(userId)
	if err != nil {
		return err
	}
	return nil
}

func (service *authService) RefreshToken(token string) (acessToken, refreshToken string, err error) {
	ok, err := service.jwtmanager.IsTokenExpired(token)
	if err != nil || !ok {
		return "", "", errors.New("can not refresh with this your token")
	}
	claims, err := service.jwtmanager.ExtractClaimsFromAcessToken(token)
	if err != nil {
		return "", "", err
	}
	if err != nil {
		return "", "", err
	}
	user := domain.User{
		First_name:   claims.First_name,
		Last_name:    claims.Last_name,
		Phone_number: claims.Phone,
		User_id:      claims.Id,
	}
	refreshtokens, err := service.authRepo.GetRefreshTokens(user.User_id)
	if err != nil {
		return "", "", err
	}
	if len(refreshtokens) <= 0 {
		return "", "", errors.New("invalid refresh token")
	}
	refreshCount := 0
	for _, rt := range refreshtokens {
		if rt == refreshToken {
			refreshCount++
		}
	}
	if refreshCount <= 0 {
		return "", "", errors.New("invalid refresh token")
	}
	acessToken, err = service.jwtmanager.GenerateAcessToken(user)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = service.jwtmanager.GenerateRefreshToken(claims.Id)
	if err != nil {
		return "", "", err
	}
	err = service.authRepo.AddRefreshToken(user.Phone_number, refreshToken)
	if err != nil {
		return "", "", err
	}
	return acessToken, refreshToken, nil
}

func (service *authService) publishLoginEvent(id string) error {
	userlogin := domain.LoginEvent{
		UserId: id,
	}
	err := service.Publisher.Publish(domain.USERLOGIN, userlogin)
	if err != nil {
		return err
	}
	return nil
}

func (service *authService) publishLogOutEvent(id string) error {
	userlogout := domain.LogOutEvent{
		UserId: id,
	}

	err := service.Publisher.Publish(domain.USERLOGOUT, userlogout)
	if err != nil {
		return err
	}
	return nil
}
