package alphanumeric

import (
	"regexp"

	"github.com/sipkg/validate/helper"
	"github.com/sipkg/validate/messages"
	"github.com/sipkg/validate/rules"
)

func init() {
	rules.Add("Alphanumeric", Alphanumeric)
}

// Validates that a string only contains alphabetic or numeric characters
func Alphanumeric(data rules.ValidationData) (err error) {
	v, ok := helper.ToString(data.Value)
	if ok != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        messages.Translate("is not a string"),
		}
	}

	if regexp.MustCompile(`[^a-zA-Z0-9]+`).MatchString(v) {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        messages.Translate("contains non-alphanumeric characters"),
		}
	}

	return nil
}
