package validators

import (
	"fmt"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// IsNotNull packages wraps Nulls type (ie. any type that implements the nullable interface) with its name and message for validation
// To call this validator on a nullable value, set Field: nulls.Nulls{Value: <nullable>}, where <nullable> could be nulls.Int, nulls.String, etc.
type IsNotNull struct {
	Name    string
	Field   nulls.Nulls
	Message string
}

// IsValid checks if nullable field is null; if so returns error
func (v *IsNotNull) IsValid(errors *validate.Errors) {
	if v.Field.Interface() != nil {
		return
	}

	if len(v.Message) > 0 {
		errors.Add(validators.GenerateKey(v.Name), v.Message)
		return
	}

	errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s can not be null.", v.Name))
}
