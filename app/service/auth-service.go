package service

import (
	"errors"
	"gin_serve/app/dto"
	"gin_serve/app/model"
	"gin_serve/app/repo"
	"gin_serve/helper"

	"github.com/mashingan/smapping"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) (dto.UserDTO, error)
	CreateUser(user dto.UserRegisterDTO) dto.UserDTO
	FindByEmail(email string) dto.UserDTO
	FindByID(id uint64) dto.UserDTO
	IsDuplicateEmail(email string) bool
	VerifyEmail(tokenStr string) bool
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
				zap.S().Errorf("Failed map %v", err)
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
		zap.S().Errorf("Failed map %v", err)
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
		zap.S().Errorf("Failed map %v", err)
	}

	return userToDTO
}

// FindByID be used when find user
func (service *authService) FindByID(id uint64) dto.UserDTO {

	userToDTO := dto.UserDTO{}
	user := service.userRepos.FindByID(id)
	err := smapping.FillStruct(&userToDTO, smapping.MapFields(&user))

	if err != nil {
		zap.S().Errorf("Failed map %v", err)
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
		zap.S().Infof(err.Error())
		return false
	}

	return true
}

func (service *authService) VerifyEmail(tokenStr string) bool {

	emailClaims, valid, err := helper.ValidateEmailTokenAndBackClaims(tokenStr)

	if err == nil && valid {
		zap.S().Info(emailClaims.UserID, emailClaims.Email)
		return service.userRepos.VerifyActiveEmail(emailClaims.UserID, emailClaims.Email)
	}

	return false
}
