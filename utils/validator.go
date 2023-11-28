package utils

import (
	"errors"

	"github.com/RedFC/testproject/models"
)

// validation function for product
func ValidateProductInput(product models.Product) error {
	if product.Name == "" {
		return errors.New("invalid input data product 'Name' is required")
	} else if product.Description == "" {
		return errors.New("invalid input data product 'Description' is required")
	} else if product.Price <= 0 {
		return errors.New("invalid input data product 'Price' is required")

	}
	return nil
}

// validation function for user register
func ValidateUserInput(user models.User) error {
	if user.Name == "" {
		return errors.New("invalid input data user 'Name' is required")
	} else if user.Email == "" {
		return errors.New("invalid input data user 'Email' is required")
	} else if user.Password == "" {
		return errors.New("invalid input data user 'Password' is required")

	}
	return nil
}

// validation function for user login
func ValidateUserLogin(user models.Login) error {
	if user.Email == "" {
		return errors.New("invalid input data user 'Email' is required")
	} else if user.Password == "" {
		return errors.New("invalid input data user 'Password' is required")

	}
	return nil
}
