package model

import (
	"regexp"

	"github.com/pkg/errors"
)

type UserCreateReq struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"pass"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	Country     string `json:"country"`
}

type UserCreateRes struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CreateUserInput struct {
	ID             string
	Email          string
	Name           string
	HashedPassword string
	PhoneNumber    string
	Address        string
	City           string
	State          string
	ZipCode        string
	Country        string
}

func (req *UserCreateReq) Validate() error {
	if err := validateEmail(req.Email); err != nil {
		return err
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if req.PhoneNumber != "" {
		if err := validatePhoneNumber(req.PhoneNumber); err != nil {
			return err
		}
	}

	return nil
}

func validateEmail(email string) error {
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func validatePhoneNumber(phone string) error {
	var phoneRegex = regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	if !phoneRegex.MatchString(phone) {
		return errors.New("invalid phone number format")
	}
	return nil
}