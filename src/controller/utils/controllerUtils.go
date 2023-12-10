package utils

import (
	"errors"
	"regexp"
)

func CepValidator(cep string) error {
	re := regexp.MustCompile(`^[0-9]{8}$`)

	if !(re.MatchString(cep)) {
		return errors.New("cep is invalid")
	}

	return nil
}
