package validator

import "github.com/go-playground/validator/v10"

type Wrapper struct {
	validator *validator.Validate
}

func (w *Wrapper) Validate(i any) error {
	return w.validator.Struct(i)
}

func New() *Wrapper {
	return &Wrapper{
		validator: validator.New(),
	}
}
