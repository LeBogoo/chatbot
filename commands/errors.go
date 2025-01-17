package commands

import "errors"

var ErrCommandAlreadyExists = errors.New("command already exists")
var ErrAliasAlreadyExists = errors.New("alias already exists")
var ErrCommandNotFound = errors.New("command not found")
var ErrAliasNotFound = errors.New("alias not found")
var ErrInvalidPrefix = errors.New("invalid prefix")
