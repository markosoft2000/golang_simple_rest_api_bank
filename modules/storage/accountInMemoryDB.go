package storage


import (
	"sync"
	"../../models"
)

type AccountInMemoryDB struct {
	m map[int] *models.Account
	lock sync.RWMutex
}

func NewInMemoryDB() DB {
	return &AccountInMemoryDB{m: make(map[int] *models.Account)}
}

func (d *AccountInMemoryDB) Get(key int) (*models.Account, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	v, ok := d.m[key]

	if !ok {
		return nil, ErrNotFound
	}

	return v, nil
}

func (d *AccountInMemoryDB) Set(key int, val *models.Account) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.m[key] = val
	_, ok := d.m[key];

	return ok
}

func (d *AccountInMemoryDB) Remove(key int) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	delete(d.m, key)
	_, ok := d.m[key];

	return !ok
}