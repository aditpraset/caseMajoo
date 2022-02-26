package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"size:255;not null;unique" json:"user_name"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedBy int64     `gorm:"default:0" json:"created_by"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy int64     `gorm:"default:0" json:"updated_by"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type ResponseUser struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
	Data    []User
}

func (u *User) Prepare() {
	u.ID = 0
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.CreatedBy = 0
	u.CreatedAt = time.Now()
	u.UpdatedBy = 0
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.UserName == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Name == "" {
			return errors.New("Required Name")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.UserName == "" {
			return errors.New("Required Username")
		}

		return nil

	default:
		if u.UserName == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Name == "" {
			return errors.New("Required Name")
		}

		return nil
	}
}
