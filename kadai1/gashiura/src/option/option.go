package option

import (
	"errors"
	"regexp"
)

func Valid(s, d string) error {
	r := regexp.MustCompile("^(jpeg|png|gif)$")
	if !r.MatchString(s) || !r.MatchString(d) {
		return errors.New("please set \"jpeg\" or \"png\" or \"gif\" as an option.")
	}
	return nil
}
