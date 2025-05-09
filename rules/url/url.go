package url

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/sipkg/validate/helper"
	"github.com/sipkg/validate/messages"
	"github.com/sipkg/validate/rules"
)

func init() {
	rules.Add("URL", URL)
}

// Validates a URL using url.Parse() in the net/url library.
// For a valid URL, the following need to be present in a parsed URL:
// * Scheme (either http or https)
// * Host (without a backslash)
func URL(data rules.ValidationData) error {
	v, err := helper.ToString(data.Value)
	if err != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        messages.Translate("is not a string"),
		}
	}

	parsed, err := url.Parse(v)
	if err != nil {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        messages.Translate("is not a valid URL"),
		}
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        fmt.Sprintf(messages.Translate("has an invalid scheme '%s'"), parsed.Scheme),
		}
	}

	if parsed.Host == "" || strings.IndexRune(parsed.Host, '\\') > 0 {
		return rules.ErrInvalid{
			ValidationData: data,
			Failure:        fmt.Sprintf(messages.Translate("has an invalid host '%s'"), parsed.Host),
		}
	}

	return nil
}
