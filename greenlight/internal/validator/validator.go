package validator

import "regexp"

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
	v
}
func (v Validator) Valid() bool {
	return len(v.Errors) == 0
}
func (v Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}
func (v Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}
func PermittedValue[T comparable](value T, permittedValue ...T) bool {
	for i := range permittedValue {
		if value == permittedValue[i] {
			return true
		}
	}
	return false
}
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}
