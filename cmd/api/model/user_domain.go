package model

import (
	"github.com/walber-vaz/crud-go/configs/logger"
	"github.com/walber-vaz/crud-go/configs/rest_error"
	"golang.org/x/crypto/bcrypt"
)

type UserDomain struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	Age       int8
}

func NewUserDomain(
	email,
	password,
	firstName,
	lastName string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		email,
		password,
		firstName,
		lastName,
		age,
	}
}

func (u *UserDomain) EncryptPassword() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Error encrypting password", err)
		return
	}
	u.Password = string(hashedPassword)
}

func (u *UserDomain) ComparePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		logger.Error("Error comparing password", err)
		return false
	}
	return true
}

type UserDomainInterface interface {
	Create() *rest_error.RestError
	Update(string) *rest_error.RestError
	Delete(string) *rest_error.RestError
	GetByEmail(email string) (*UserDomain, *rest_error.RestError)
	GetAll() ([]UserDomain, *rest_error.RestError)
	GetById(id string) (*UserDomain, *rest_error.RestError)
}
