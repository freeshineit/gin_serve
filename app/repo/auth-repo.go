package repo

import (
	"gin_serve/app/dto"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepo is contract what userRepository can do tod db
type UserRepo interface {
	InsertUser(user dto.User) dto.User
	UpdateUser(user dto.User) dto.User
	VerifyCredential(email string) (dto.User, error)
	IsDuplicateEmail(email string) *gorm.DB
	FindByEmail(email string) dto.User
	ProfileUser(userID string) dto.User
}

type userConnection struct {
	connection *gorm.DB
}

// NewUserRepo new user repository
func NewUserRepo(db *gorm.DB) UserRepo {
	return &userConnection{
		connection: db,
	}
}

// InsertUser be used when inert user to database
func (db *userConnection) InsertUser(user dto.User) dto.User {
	// hash password
	user.Password = hashAndSalt([]byte(user.Password))

	db.connection.Save(&user)
	return user
}

// UpdateUser be used when update user to database
func (db *userConnection) UpdateUser(user dto.User) dto.User {
	return dto.User{}
}

// VerifyCredential be used when verify credential
func (db *userConnection) VerifyCredential(email string) (dto.User, error) {
	var user dto.User
	res := db.connection.Where("email=?", email).Take(&user)

	if res.Error == nil {
		return user, nil
	}

	return user, res.Error
}

// IsDuplicateEmail be used when verify duplicate email
func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user dto.User
	res := db.connection.Where("email = ?", email).Take(&user)

	return res
}

// FindByEmail be used when find user by email form database
func (db *userConnection) FindByEmail(email string) dto.User {
	var user dto.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

// ProfileUser be used when find user by user id form database
func (db *userConnection) ProfileUser(userID string) dto.User {
	var user dto.User
	db.connection.Find(&user, userID)
	return user
}

// hashAndSalt be used when encrypt password
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		log.Panicln("Failed to hash a password")
	}

	return string(hash)
}
