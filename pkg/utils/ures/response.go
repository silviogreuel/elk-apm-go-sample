package ures

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func BodyAsString(r *http.Response) (string, error) {
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(bb), nil
}
