package uuid

import (
	"regexp"

	"github.com/sipkg/validate/helper"
	"github.com/sipkg/validate/messages"
	"github.com/sipkg/validate/rules"
)

func init() {
	rules.Add("UUID", UUID)
}

// Used to check whether a string has at most N characters
// Fails if data is a string and its length is more than the specified comparator. Passes in all other cases.
func UUID(data rules.ValidationData) error {
	v, err := helper.ToString(data.Value)
	if err != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        messages.Translate("is not a string"),
		}
	}

	if !IsUUID(v) {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        messages.Translate("is an invalid UUID"),
		}
	}

	return nil
}

func IsUUID(uuid string) bool {
	hexPattern := "^(urn\\:uuid\\:)?\\{?([a-z0-9]{8})-([a-z0-9]{4})-([1-5][a-z0-9]{3})-([a-z0-9]{4})-([a-z0-9]{12})\\}?$"
	re := regexp.MustCompile(hexPattern)

	if match := re.FindStringSubmatch(uuid); match == nil {
		return false
	}
	return true
}
