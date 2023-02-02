package repository

import (
	"gin_serve/app/model"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository is contract what userRepository can do tod db
type UserRepository interface {
	InsertUser(user model.User) model.User
	UpdateUser(user model.User) model.User
	VerifyCredential(email string) (model.User, error)
	IsDuplicateEmail(email string) *gorm.DB
	FindByEmail(email string) model.User
	ProfileUser(userID string) model.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

// inert user to database
func (db *userConnection) InsertUser(user model.User) model.User {
	// hash password
	user.Password = hashAndSalt([]byte(user.Password))

	db.connection.Save(&user)
	return user
}

// update user to database
func (db *userConnection) UpdateUser(user model.User) model.User {
	return model.User{}

}

// verify credential
func (db *userConnection) VerifyCredential(email string) (model.User, error) {
	var user model.User
	res := db.connection.Where("email=?", email).Take(&user)

	if res.Error == nil {
		return user, nil
	}
	return user, res.Error
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user model.User
	return db.connection.Where("email = ?", email).Take(&user)
}

// find user by email
func (db *userConnection) FindByEmail(email string) model.User {
	var user model.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) model.User {
	var user model.User
	db.connection.Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		log.Panicln("Failed to hash a password")
	}

	return string(hash)
}
