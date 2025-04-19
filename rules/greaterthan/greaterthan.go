package greaterthan

import (
	"fmt"
	"strconv"

	"github.com/sipkg/validate/helper"
	"github.com/sipkg/validate/messages"
	"github.com/sipkg/validate/rules"
)

func init() {
	rules.Add("GreaterThan", GreaterThan)
}

// Passes if the data is a float/int and is greater than the specified integer.
// Note that this is *not* a greater than or equals check, and this only comapres
// floats/ints to a predefined integer specified in your tag.
// Fails if the data is not a float/int or the data is less than or equals the comparator
func GreaterThan(data rules.ValidationData) error {
	v, err := helper.ToFloat64(data.Value)
	if err != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        messages.Translate("is not numeric"),
		}
	}

	// We should always be provided with a length to validate against
	if len(data.Args) == 0 {
		return fmt.Errorf("no argument found in the validation struct (eg 'GreaterThan:5')")
	}

	// Typecast our argument and test
	var min float64
	if min, err = strconv.ParseFloat(data.Args[0], 64); err != nil {
		return err
	}

	if v < min {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        fmt.Sprintf(messages.Translate("must be greater than %f"), min),
		}
	}

	return nil
}
