package service

import (
	"errors"
	"gin_serve/app/dto"
	"gin_serve/app/model"
	"gin_serve/app/repo"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) (dto.UserDTO, error)
	CreateUser(user dto.UserRegisterDTO) dto.UserDTO
	FindByEmail(email string) dto.UserDTO
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
func (service *authService) VerifyCredential(email string, password string) (dto.UserDTO, error) {

	userToCreate := dto.UserDTO{}

	user, err := service.userRepos.VerifyCredential(email)

	if err == nil {
		result := comparePassword(user.Password, password)
		if result && user.Email == email {
			err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))

			if err != nil {
				log.Fatalf("Failed map %v", err)
			}

			return userToCreate, nil
		} else {
			return dto.UserDTO{}, errors.New("password error")
		}
	}

	return dto.UserDTO{}, errors.New("user does not exist")
}

// CreateUser be used when create user
func (service *authService) CreateUser(user dto.UserRegisterDTO) dto.UserDTO {
	userToCreate := model.User{}

	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	userToDTO := dto.UserDTO{}

	res := service.userRepos.InsertUser(userToCreate)
	smapping.FillStruct(&userToDTO, smapping.MapFields(&res))

	return userToDTO
}

// FindByEmail be used when find user
func (service *authService) FindByEmail(email string) dto.UserDTO {

	userToDTO := dto.UserDTO{}
	user := service.userRepos.FindByEmail(email)
	err := smapping.FillStruct(&userToDTO, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	return userToDTO
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
