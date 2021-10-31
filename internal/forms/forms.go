package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/url"
	"strings"
)

// Form creates a custom from struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be empty")
		}
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string) bool {
	entry := f.Get(field)
	if entry == "" {
		f.Errors.Add(field, "This field cannot be empty")
		return false
	}
	return true
}

// MinLength checks for string minimum length, returns true if meets minimum length
func (f *Form) MinLength(field string, minLength int) bool {
	entry := f.Get(field)
	if len(entry) < minLength {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", minLength))
		return false
	}
	return true
}

// IsEmail checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
