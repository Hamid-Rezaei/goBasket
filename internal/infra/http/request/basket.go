package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type BasketCreate struct {
	Data  string `json:"data,omitempty"   validate:"required,alphanumunicode"`
	State string `json:"state,omitempty" validate:"required,eq=COMPLETED|eq=PENDING"`
}

type BasketUpdate struct {
	Data  string `json:"data,omitempty"   validate:"omitempty,alphanumunicode"`
	State string `json:"state,omitempty" validate:"omitempty,eq=COMPLETED|eq=PENDING"`
}

func (bc BasketCreate) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(bc); err != nil {
		return fmt.Errorf("create request validation failed %w", err)
	}

	return nil
}

func (bu BasketUpdate) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(bu); err != nil {
		return fmt.Errorf("update request validation failed %w", err)
	}

	return nil
}
