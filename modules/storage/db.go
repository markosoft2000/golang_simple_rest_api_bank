package storage


import (
	"errors"
	"../../models"
)

var (
	ErrNotFound = errors.New("Error: item not found")
)

type DB interface {
	Get(key int) (*models.Account, error)
	Set(key int, val *models.Account) bool
	Remove(key int) bool
}