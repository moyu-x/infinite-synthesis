package fiber

import (
	"github.com/gookit/validate"
)

type Validator struct{}

func (v *Validator) Validate(out any) error {
	c := validate.Struct(out)

	if c.Validate() {
		return nil
	} else {
		return c.Errors
	}
}
