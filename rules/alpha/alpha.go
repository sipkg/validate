package alpha

import (
	"regexp"

	"github.com/sipkg/validate/helper"
	"github.com/sipkg/validate/messages"
	"github.com/sipkg/validate/rules"
)

func init() {
	rules.Add("Alpha", Alpha)
}

// Validates that a string only contains alphabetic characters
func Alpha(data rules.ValidationData) (err error) {
	v, ok := helper.ToString(data.Value)
	if ok != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        messages.Translate("is not a string"),
		}
	}

	if regexp.MustCompile(`[^a-zA-Z]+`).MatchString(v) {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        messages.Translate("contains non-alphabetic characters"),
		}
	}

	return nil
}
