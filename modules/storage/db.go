package storage

import (
	"errors"
	"../../models"
)

var (
	ErrNotFound = errors.New("error: item not found")
	ErrAlreadyExists = errors.New("error: item already exists")
	ErrCreateFailed = errors.New("error: item creation failed")
)

type DB interface {
	Get(key int) (*models.Account, error)
	Create(key int, val *models.Account) error
	Remove(key int) bool
}

type DBPlus interface {
	DB
	Len() int
}