package http

import (
	"github.com/gookit/validate"
)

type Validator struct{}

func (v *Validator) Engine() any {
	return ""
}

func (v *Validator) ValidateStruct(out any) error {
	c := validate.Struct(out)

	if c.Validate() {
		return nil
	} else {
		return c.Errors
	}
}
