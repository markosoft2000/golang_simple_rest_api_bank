package storage

import (
	"sync"
	"../../models"
)

type AccountInMemoryDB struct {
	m map[int] *models.Account
	lock sync.RWMutex
}

var createMutex = new(sync.RWMutex)

func NewInMemoryDB() DB {
	return &AccountInMemoryDB{m: make(map[int] *models.Account)}
}

func (d *AccountInMemoryDB) Get(key int) (*models.Account, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	v, ok := d.m[key]

	if !ok {
		return nil, ErrNotFound
	}

	return v, nil
}

func (d *AccountInMemoryDB) set(key int, val *models.Account) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.m[key] = val
	_, ok := d.m[key];

	return ok
}

func (d *AccountInMemoryDB) Create(key int, val *models.Account) error {
	createMutex.Lock()
	defer createMutex.Unlock()

	_, ok := d.Get(key)
	if ok == nil {
		return ErrAlreadyExists
	}

	if ok := d.set(key, val); !ok {
		return ErrCreateFailed
	}

	return nil
}

func (d *AccountInMemoryDB) Remove(key int) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	delete(d.m, key)
	_, ok := d.m[key];

	return !ok
}

func (d *AccountInMemoryDB) Len() int {
	return len(d.m)
}