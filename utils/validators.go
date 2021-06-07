package utils

import (
	"edith/models"
	"regexp"

	valid "github.com/asaskevich/govalidator"
)

// IsEmpty check if string is empty
func IsEmpty(str string) (bool, string) {
	if valid.HasWhitespaceOnly(str) && str != "" {
		return true, "Must not be empty"
	}

	return false, ""
}

// ValidateRegister validates the body of the user for registration
func ValidateRegister(u *models.User) *models.UserErrors {
	e := &models.UserErrors{}
	e.Err, e.UserName = IsEmpty(u.UserName)

	if !valid.IsEmail(u.Email) {
		e.Err, e.Email = true, "Must be a valid email"
	}

	re := regexp.MustCompile("\\d") //regex check for atleast one integer
	if !(len(u.Password) >= 8 && valid.HasLowerCase(u.Password) && valid.HasUpperCase(u.Password) && re.MatchString(u.Password)) {
		e.Err, e.Password = true, "Password should be atleast 8 characters and it must be a combination of uppercase letters, lowercase letters and numbers"
	}

	return e
}
