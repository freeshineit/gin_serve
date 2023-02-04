package service

import (
	"errors"
	"gin_serve/app/dto"
	"gin_serve/app/repo"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) (dto.User, error)
	CreateUser(user dto.UserRegisterDTO) dto.User
	FindByEmail(email string) dto.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepos repo.UserRepo
}

func NewAuthService(userRepo repo.UserRepo) AuthService {
	return &authService{
		userRepos: userRepo,
	}
}

// VerifyCredential be used when verify email and password
func (service *authService) VerifyCredential(email string, password string) (dto.User, error) {

	user, err := service.userRepos.VerifyCredential(email)

	if err == nil {
		result := comparePassword(user.Password, password)
		if result && user.Email == email {
			return user, nil
		} else {
			return dto.User{}, errors.New("password error")
		}
	}

	return dto.User{}, errors.New("user does not exist")
}

// CreateUser be used when create user
func (service *authService) CreateUser(user dto.UserRegisterDTO) dto.User {
	userToCreate := dto.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.userRepos.InsertUser(userToCreate)

	return res
}

// FindByEmail be used when find user
func (service *authService) FindByEmail(email string) dto.User {
	return service.userRepos.FindByEmail(email)
}

// 判读重复邮箱
func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepos.IsDuplicateEmail(email)

	// 重复
	return res.Error == nil
}

// Compare password
func comparePassword(hashPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))

	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}
