package models

import (
	"errors"
)

//go:generate gen-validator -model=User
type User struct {
	Name string `validator:"min=3,max=12"`
	Age  int    `validator:"min=18"`
}

func (obj *User) ErrorMaxName() error {
	return errors.New("Ensure this value is less than or equal to 12")
}

func (obj *User) ErrorMinAge() error {
	return errors.New("Ensure this value is greater than or equal to 18")
}
