package validate

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/sipkg/validate/rules"
	_ "github.com/sipkg/validate/rules/alpha"
	_ "github.com/sipkg/validate/rules/alphanumeric"
	_ "github.com/sipkg/validate/rules/email"
	_ "github.com/sipkg/validate/rules/greaterthan"
	_ "github.com/sipkg/validate/rules/length"
	_ "github.com/sipkg/validate/rules/lessthan"
	_ "github.com/sipkg/validate/rules/maxlength"
	_ "github.com/sipkg/validate/rules/minlength"
	_ "github.com/sipkg/validate/rules/notempty"
	_ "github.com/sipkg/validate/rules/notzero"
	_ "github.com/sipkg/validate/rules/notzerotime"
	_ "github.com/sipkg/validate/rules/regexp"
	_ "github.com/sipkg/validate/rules/url"
	_ "github.com/sipkg/validate/rules/uuid"
)

// Takes a struct, loops through all fields and calls check on any fields that
// have a validate tag. If the field is an anonymous struct recursively run
// validation on it.
func Run(object any, fieldsSlice ...string) error {
	pass := true // We'll override this if checking returns false
	err := ValidationError{}

	// If we have been passed a slice of fields to valiate - to check only a
	// subset of fields - change the slice into a map for O(1) lookups instead
	// of O(n).
	fields := map[string]struct{}{}
	for _, field := range fieldsSlice {
		fields[field] = struct{}{}
	}

	// If we're passed a pointer to a struct we need to dereference the pointer before
	// inspecting its tags
	value := reflect.ValueOf(object)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Iterate through each field of the struct and validate
	typ := value.Type() // A Type's Field method returns StructFields
	for i := range value.NumField() {
		var validateTag string
		var validateError error

		// Is this an anonymous struct? If so, we also need to validate on this.
		if typ.Field(i).Anonymous {
			if anonErr := Run(value.Field(i).Interface(), fieldsSlice...); anonErr != nil {
				// The validation failed: set pass to false and merge the anonymous struct's
				// validation errors with our current validation error above to give a complete
				// error message.
				pass = false

				// A non validation error occurred: return this immediately
				if _, ok := anonErr.(ValidationError); !ok {
					return anonErr
				}

				err.Merge(anonErr.(ValidationError))
			}
		}

		if len(fields) > 0 {
			// We're only checking for a subset of fields; if this field isn't
			// included in the subset of fields to validate we can skip.
			if _, ok := fields[typ.Field(i).Name]; !ok {
				continue
			}
		}

		if validateTag = typ.Field(i).Tag.Get("validate"); validateTag == "" {
			continue
		}

		// Validate this particular field against the options in our tag
		if validateError = validateField(value.Field(i).Interface(), typ.Field(i).Name, validateTag); validateError == nil {
			continue
		}

		// If there was no validation rule defined for the given tag return
		// that error immediately.
		if _, ok := validateError.(rules.ErrNoValidationMethod); ok {
			return validateError
		}

		pass = false
		err.addFailure(typ.Field(i).Name, validateError.Error())
	}

	if pass {
		return nil
	}

	return err
}

var rxRegexp = regexp.MustCompile(`Regexp:\/.+/`)

// Takes a field's value and the validation tag and applies each check
// until either a test fails or all tests pass.
func validateField(data any, fieldName, tag string) (err error) {
	// A tag can specify multiple validation rules which are delimited via ','.
	// However, because we allow regular expressions we can't split the tag field
	// via all commas to find our validation rules: we need to extract the regular expression
	// first (in case it specifies a comma), and *then* run through our validation rules.
	if match := rxRegexp.FindString(tag); match != "" {
		// If we fail validating the regexp we can break here
		if err := validateRule(data, fieldName, match); err != nil {
			return err
		}
		// Now we need to replace our regular expression from the tag list to continue
		// normally.
		tag = rxRegexp.ReplaceAllString(tag, "")
	}

	for tag != "" {
		var next string

		i := strings.Index(tag, ",")
		if i >= 0 {
			tag, next = tag[:i], tag[i+1:]
		}

		if err := validateRule(data, fieldName, tag); err != nil {
			return err
		}

		// Continue with the next tag
		tag = next
	}

	return nil
}

// Given a validation rule from a tag, run the associated validation methods and return
// the result.
func validateRule(data any, fieldName, rule string) error {
	var args []string

	// Remove any preceeding spaces from comma separated tags
	rule = strings.TrimLeft(rule, " ")

	// If the rule is empty we don't need to process anything. This only happens
	// if we have a regex followed by another rule:
	//   `validate:"Regexp:/.+/, NotEmpty"`
	// Becomes:
	//   `validate:", NotEmpty"`
	// After processing in validateField()
	if rule == "" {
		return nil
	}

	// rule is the method we want to call. If it has a colon we need to further
	// process the rule to extract arguments to our validation method.
	i := strings.Index(rule, ":")
	if i > 0 {
		var arg string
		rule, arg = rule[:i], rule[i+1:]
		args = append(args, arg)
	}

	// Attempt to validate the data using methods registered with the rules
	// sub package
	if method, err := rules.Get(rule); err != nil {
		return err
	} else {
		data := rules.ValidationData{
			Field: fieldName,
			Value: data,
			Args:  args,
		}
		return method(data)
	}
}
