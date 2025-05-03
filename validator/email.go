package validator

import (
	"fmt"
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (v *Validator) IsEmail(field, email string) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if !emailRegexp.MatchString(email) {
		v.Errors[field] = fmt.Sprintf("%s not a valid email", email)
		return false
	}

	return true
}
