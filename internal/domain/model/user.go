package model

import (
	"time"
	"yourapp/internal/core/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	model.BaseModel
	model.SoftDeleteModel
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	FullName string `gorm:"not null" json:"full_name"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

type VerificationToken struct {
	ID        string    `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	UserID    string    `gorm:"not null"`
	Token     string    `gorm:"not null;uniqueIndex"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      User           `gorm:"foreignKey:UserID"`
}
