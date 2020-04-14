package validator

import (
	"errors"
	"regexp"
	"strconv"
	// "log"
)

// ValidateString : used for validating string
func ValidateString(name string , str string , minimum int, maximum int, regex string ) error {

	strLength := len(str)

	if strLength < minimum {
		return errors.New(name + " cannot be less than "+ strconv.Itoa(minimum) + " characters")
	}

	if strLength > maximum {
		return errors.New(name + " cannot be more than "+ strconv.Itoa(maximum) + " characters")
	}

	if regex != "" {
		re := regexp.MustCompile(regex)
		val := re.MatchString(str)
		if val == false {
			return errors.New(name + " is invalid")
		}
	}

	return nil
}

// ValidateMobile : used for validating string
func ValidateMobile(name string , mobile string , minimum int, maximum int , regex string) error {
	

	mobLength := len(mobile)
	re := regexp.MustCompile(regex)
	val := re.MatchString(mobile)
	if val == false {
		return errors.New(name + " is invalid")
	}

	if mobLength < minimum {
		return errors.New(name + " cannot be less than "+ strconv.Itoa(minimum) + " digits")
	}

	if mobLength > maximum {
		return errors.New(name + " cannot be more than "+ strconv.Itoa(maximum) + " digits")
	}

	
	

	return nil
}