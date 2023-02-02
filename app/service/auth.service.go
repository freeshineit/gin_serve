package service

import (
	"errors"
	"gin_serve/app/dto"
	"gin_serve/app/model"
	"gin_serve/app/repository"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) (model.User, error)
	CreateUser(user dto.UserCreateDTO) model.User
	FindByEmail(email string) model.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) (model.User, error) {

	user, err := service.userRepository.VerifyCredential(email, password)

	if err == nil {
		result := comparePassword(user.Password, password)
		if result && user.Email == email {
			return user, nil
		}

		return model.User{}, err
	}

	return model.User{}, errors.New("failed to verify credential")
}

func (service *authService) CreateUser(user dto.UserCreateDTO) model.User {
	userToCreate := model.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.userRepository.InsertUser(userToCreate)

	return res
}
func (service *authService) FindByEmail(email string) model.User {
	return service.userRepository.FindByEmail(email)

}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return res.Error != nil
}

// Compare password
func comparePassword(hashPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	if err != nil {
		log.Panicln(err.Error())

		return false
	}

	return true
}
