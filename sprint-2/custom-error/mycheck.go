//go:build !solution

package mycheck

import (
	"errors"
	"strings"
)

type myError struct {
	Errs []error
}

func (err myError) Error() string {

	sb := strings.Builder{}

	for _, elem := range err.Errs {
		if len(sb.String()) != 0 {
			sb.WriteString(";")
		}

		sb.WriteString(elem.Error())
	}

	return sb.String()
}

func MyCheck(input string) error {
	var myErr myError

	ok := strings.ContainsAny(input, "0123456789")
	if ok {
		myErr.Errs = append(myErr.Errs, errors.New("found numbers"))
	}

	if len(input) > 20 {
		myErr.Errs = append(myErr.Errs, errors.New("line is too long"))
	}

	cnt := 0
	for _, ch := range input {
		if ch == ' ' {
			cnt++
		}
	}
	if cnt != 2 {
		myErr.Errs = append(myErr.Errs, errors.New("no two spaces"))
	}

	if len(myErr.Errs) == 0 {
		return nil
	}

	return myErr
}
