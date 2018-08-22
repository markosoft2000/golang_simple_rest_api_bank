package helpers

import (
	"net/http"
	"errors"
)

func GetRequestParam(r *http.Request, name string) (string, error) {
	for paramName, paramValue := range r.Form {
		if paramName == name {
			return paramValue[0], nil
		}
	}

	return "", errors.New("Error: param not found")
}
