package repo

import (
	"gin_serve/app/model"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepo is contract what userRepository can do tod db
type UserRepo interface {
	InsertUser(user model.User) model.User
	UpdateUser(user model.User) model.User
	VerifyCredential(email string) (model.User, error)
	IsDuplicateEmail(email string) *gorm.DB
	FindByEmail(email string) model.User
	ProfileUser(userID string) model.User
	VerifyActiveEmail(userID uint64, email string) bool
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
func (db *userConnection) InsertUser(user model.User) model.User {
	// hash password
	user.Password = hashAndSalt([]byte(user.Password))

	db.connection.Save(&user)
	return user
}

// UpdateUser be used when update user to database
func (db *userConnection) UpdateUser(user model.User) model.User {
	return model.User{}
}

// VerifyCredential be used when verify credential
func (db *userConnection) VerifyCredential(email string) (model.User, error) {
	var user model.User
	res := db.connection.Where("email=?", email).Take(&user)

	if res.Error == nil {
		return user, nil
	}

	return user, res.Error
}

// IsDuplicateEmail be used when verify duplicate email
func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user model.User
	res := db.connection.Where("email = ?", email).Take(&user)

	return res
}

// FindByEmail be used when find user by email form database
func (db *userConnection) FindByEmail(email string) model.User {
	var user model.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

// ProfileUser be used when find user by user id form database
func (db *userConnection) ProfileUser(userID string) model.User {
	var user model.User
	db.connection.Find(&user, userID)
	return user
}

// VerifyActiveEmail
func (db *userConnection) VerifyActiveEmail(userID uint64, email string) bool {
	user := model.User{ID: userID, Email: email}

	err := db.connection.Model(&user).Update("is_active", 1).Error

	if err != nil {
		return false
	}

	return true
}

// hashAndSalt be used when encrypt password
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		zap.S().Panicln("Failed to hash a password")
	}

	return string(hash)
}
