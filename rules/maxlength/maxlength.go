package maxlength

import (
	"fmt"
	"strconv"

	"github.com/sipkg/validate/helper"
	"github.com/sipkg/validate/messages"
	"github.com/sipkg/validate/rules"
)

func init() {
	rules.Add("MaxLength", MaxLength)
}

// Used to check whether a string has at most N characters
// Fails if data is a string and its length is more than the specified comparator. Passes in all other cases.
func MaxLength(data rules.ValidationData) error {
	v, err := helper.ToString(data.Value)
	if err != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        messages.Translate("is not a string"),
		}
	}

	// We should always be provided with a length to validate against
	if len(data.Args) == 0 {
		return fmt.Errorf("no argument found in the validation struct (eg 'MaxLength:5')")
	}

	// Typecast our argument and test
	var max int
	if max, err = strconv.Atoi(data.Args[0]); err != nil {
		return err
	}
	// Typecast our argument and test

	if len(v) > max {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        fmt.Sprintf(messages.Translate("is too long; it must be at most %d characters long"), max),
		}
	}

	return nil
}
