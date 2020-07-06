package restserializer

import "errors"

var ErrNotFound = errors.New("error not found")
var ErrBadRequest = errors.New("error bad request")
var ErrForbidden = errors.New("error access forbidden")
var ErrAlreadyExist = errors.New("error already exist")
var ErrUnknown = errors.New("error unknown")
