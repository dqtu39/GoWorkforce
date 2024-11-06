package errors

import "errors"

var (
    ErrInvalidCredentials = errors.New("invalid credentials")
    ErrUserNotFound      = errors.New("user not found")
    ErrUserExists        = errors.New("user already exists")
    ErrInvalidToken      = errors.New("invalid or expired token")
    ErrDatabaseError     = errors.New("database error")
) 