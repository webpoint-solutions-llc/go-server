package user

import (
	"github.com/jellydator/validation"
	"github.com/jellydator/validation/is"
)

type CreateUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (i CreateUserInput) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.FirstName, validation.Required),
		validation.Field(&i.LastName, validation.Required),
		validation.Field(&i.Email, validation.Required, is.Email),
		validation.Field(&i.Password, validation.Required, validation.Length(6, 40)),
	)
}
