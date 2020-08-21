package validators

import (
	"fmt"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// EnumIsInRange packages the int value of an enum, along with its name and exclusive min-max values, for validation.
// To call this validator on an enum, set Field: int(<enum>), where <enum> could be any enum with an underlying int (or similar) type.
// The Start and End values must be exclusive bounds, ie. 1 less than the minimum possible value and 1 more than the minimum possible value, respectively.
type EnumIsInRange struct {
	Name    string
	Field   int
	Message string
	Start   int
	End     int
}

// IsValid :	Validates that the value of the enum falls within the acceptable range of numerical enum values
func (v *EnumIsInRange) IsValid(errors *validate.Errors) {
	if v.Field > v.Start && v.Field < v.End {
		return
	}

	if len(v.Message) > 0 {
		errors.Add(validators.GenerateKey(v.Name), v.Message)
		return
	}

	errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("The number value of %s is not in the acceptable range of enum values.", v.Name))
}
