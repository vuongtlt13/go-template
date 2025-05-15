package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a system user
type User struct {
	ID        uint64         `gorm:"primarykey" json:"id"`
	Email     string         `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Password  string         `gorm:"size:100;not null" json:"-"`
	FirstName string         `gorm:"size:50" json:"first_name"`
	LastName  string         `gorm:"size:50" json:"last_name"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Roles     []Role         `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}

// HashPassword hashes the user's password
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword checks if the provided password matches the hashed password
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
