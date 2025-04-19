package validate

// @TODO: Clean up the tests a bit

import (
	"testing"
	"time"

	"github.com/sipkg/validate/rules"
)

type Anonymous struct {
	Email string `validate:"NotEmpty"`
}

func TestAnonymousStructs(t *testing.T) {
	object := struct {
		Anonymous
		Name string `validate:"NotEmpty"`
	}{
		Name: "",
	}

	err := Run(object)
	if err == nil {
		t.Fatalf("Expected Validate to validate anonymous fields")
	}

	vErr := err.(ValidationError)

	// Validation errors should concatenate the struct and anonymous struct
	// errors together
	if len(vErr.Fields) != 2 {
		t.Fatalf("Expected ValidationError to merge Anonymous Struct errors")
	}

	if _, ok := vErr.Fields["Name"]; !ok {
		t.Fatalf("Expected ValidationError.Field to contain standard field names")
	}

	if _, ok := vErr.Fields["Email"]; !ok {
		t.Fatalf("Expected ValidationError.Field to contain anonymous field names")
	}
}

func TestNotEmpty(t *testing.T) {
	invalid := []any{
		1,
		1.5,
		int8(1),
		float64(2.333),
		struct{}{},
		[]string{"test"},
		'a',
		"",
	}

	object := struct {
		Data any `validate:"NotEmpty"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid NotEmpty values to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	object.Data = "valid"
	if err := Run(object); err != nil {
		t.Errorf("Unexpected error with valid field: %s", err.Error())
	}
}

func TestNotZero(t *testing.T) {
	invalid := []any{
		"a",
		struct{}{},
		0,
		int8(0),
		float32(0),
	}
	valid := []any{
		float64(2),
		int16(1231),
		1241631,
		float32(0.1),
		1,
		0.5,
	}

	object := struct {
		Data any `validate:"NotZero"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid NotZero values to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	for _, v := range valid {
		object.Data = v
		err := Run(object)
		if err != nil {
			t.Errorf("Unexpected error with valid NotZero value")
		}
	}
}

func TestEmail(t *testing.T) {
	invalid := []any{
		1,
		1.5,
		int8(1),
		float64(2.333),
		struct{}{},
		[]string{"test"},
		'a',
		"test@example",
		"test@example.",
		"testexample.",
		"example.com",
	}
	valid := []any{
		"test@example.com",
		"test@example.org.uk",
		"test@example.ru",
	}

	object := struct {
		Data any `validate:"Email"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid Email values to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	for _, v := range valid {
		object.Data = v
		err := Run(object)
		if err != nil {
			t.Errorf("Unexpected error with valid Email value")
		}
	}
}

func TestMinLength(t *testing.T) {
	invalid := []any{
		1,
		1.5,
		int8(1),
		float64(2.333),
		struct{}{},
		[]string{"test"},
		'a',
		"t",
	}
	valid := []any{
		"aa",
		"test",
	}

	object := struct {
		Data any `validate:"MinLength:2"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid MinLength:2 values to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	for _, v := range valid {
		object.Data = v
		err := Run(object)
		if err != nil {
			t.Errorf("Unexpected error with valid MinLength:2 value")
		}
	}
}

func TestMaxLength(t *testing.T) {
	invalid := []any{
		1,
		1.5,
		int8(1),
		float64(2.333),
		struct{}{},
		[]string{"test"},
		'a',
		"test",
	}
	valid := []any{
		"t",
	}

	object := struct {
		Data any `validate:"MaxLength:2"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid MaxLength:2 values to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	for _, v := range valid {
		object.Data = v
		err := Run(object)
		if err != nil {
			t.Errorf("Unexpected error with valid MaxLength:2 value")
		}
	}
}

func TestLength(t *testing.T) {
	invalid := []any{
		1,
		1.5,
		int8(1),
		float64(2.333),
		struct{}{},
		[]string{"test"},
		'a',
		"t",
		"foobar",
	}
	valid := []any{
		"test",
	}

	object := struct {
		Data any `validate:"Length:4"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid Length:4 values to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	for _, v := range valid {
		object.Data = v
		err := Run(object)
		if err != nil {
			t.Errorf("Unexpected error with valid Length:4 value")
		}
	}
}

func TestGreaterThan(t *testing.T) {
	invalid := []any{
		"a",
		0,
		1.5,
		int16(2),
		float64(1.25),
		49.99,
	}
	valid := []any{
		int64(100),
		float32(192.123),
		12311,
		123.6,
		50,
	}

	object := struct {
		Data any `validate:"GreaterThan:50"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid GreaterThan:50 values to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	for _, v := range valid {
		object.Data = v
		err := Run(object)
		if err != nil {
			t.Errorf("Unexpected error with valid GreaterThan:50 value")
		}
	}
}

func TestValidateUUID(t *testing.T) {
	invalid := []any{
		1,
		1.5,
		int8(1),
		float64(2.333),
		struct{}{},
		[]string{"test"},
		[]byte("test"),
		[]rune("test"),
		'a',
		"t",
		"foobar",
		"E55A815A-BA16-4FB9-AE01-644204CC433A", // Uppercase V4 - invalid hex digits
	}
	valid := []any{
		"fb623672-40dd-11e3-91ea-ce3f5508acd9", // V1
		"8563d95d-efb0-4e87-95d8-1d6c5debf298", // V4
	}

	object := struct {
		Data any `validate:"UUID"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid UUID values to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	for _, v := range valid {
		object.Data = v
		err := Run(object)
		if err != nil {
			t.Errorf("Unexpected error with valid UUID value")
		}
	}
}

func TestValidateNotZeroTime(t *testing.T) {
	invalid := []any{
		1,
		1.5,
		int8(1),
		float64(2.333),
		struct{}{},
		[]string{"test"},
		[]byte("test"),
		[]rune("test"),
		'a',
		"t",
		time.Time{},
	}
	valid := []any{
		time.Date(1984, 1, 1, 12, 0o0, 0o0, 0o0, time.UTC),
		time.Now(),
	}

	object := struct {
		Data any `validate:"NotZeroTime"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid NotZeroTime values to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	for _, v := range valid {
		object.Data = v
		err := Run(object)
		if err != nil {
			t.Errorf("Unexpected error with valid NotZeroTime value")
		}
	}
}

func TestValidateURL(t *testing.T) {
	invalid := []any{
		"test",
		"test",
		"http://",
		"example.c\\",
		"example.com",
		"http//example.com/",
		"http::/example.com/",
		"http://example\\.com",
	}

	valid := []any{
		"http://example.com",
		"http://example.com/",
		"HTTP://example.com/",
		"https://www.example.com/index.html",
	}

	object := struct {
		Data any `validate:"URL"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid URL values to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	for _, v := range valid {
		object.Data = v
		err := Run(object)
		if err != nil {
			t.Errorf("Unexpected error with valid URL value")
		}
	}
}

// Tests a regexp sandwiched in the middle of two other validation rules
func TestValidateRegexp(t *testing.T) {
	invalid := []any{
		1,
		'a',
		"0aaa0",
	}
	valid := []any{
		"aaaaa0",
		"aaa123456789",
	}

	object := struct {
		Data any `validate:"MinLength:1, Regexp:/^[a-zA-Z]{3,5}[0-9]+$/, NotEmpty"`
	}{}

	for _, v := range invalid {
		object.Data = v
		err := Run(object)
		if err == nil {
			t.Errorf("Expected invalid regexp to fail validation")
		}
		if _, ok := err.(rules.ErrNoValidationMethod); ok {
			t.Error(err.Error())
		}
	}

	for _, v := range valid {
		object.Data = v
		err := Run(object)
		if err != nil {
			t.Errorf("Unexpected error with valid regexp value: %s", err.Error())
		}
	}
}

func TestWithPointer(t *testing.T) {
	object := &struct {
		Data any `validate:"MinLength:1"`
	}{
		Data: "a",
	}

	if err := Run(object); err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
}

func TestValidateFields(t *testing.T) {
	object := &struct {
		Invalid any `validate:"MinLength:10"`
		Valid   any `validate:"NotEmpty"`
	}{
		Invalid: "a",
		Valid:   "a",
	}

	err := Run(object)
	if err == nil {
		t.Fatal()
	}

	vErr, ok := err.(ValidationError)
	if !ok {
		t.Fatal()
	}

	if len(vErr.Fields) != 1 {
		t.Fatal()
	}

	if _, ok := vErr.Fields["Invalid"]; !ok {
		t.Fatal()
	}
}
