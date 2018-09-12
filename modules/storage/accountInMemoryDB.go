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
var syncPaymentMutex = new(sync.Mutex)

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

func (d *AccountInMemoryDB) Create(key int, val *models.Account) error {
	createMutex.Lock()
	defer createMutex.Unlock()

	_, ok := d.Get(key)
	if ok == nil {
		return ErrAlreadyExists
	}

	d.m[key] = val
	if _, ok := d.m[key]; !ok {
		return ErrCreateFailed
	}

	return nil
}

func (d *AccountInMemoryDB) Remove(key int) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	delete(d.m, key)
	_, ok := d.m[key]

	return !ok
}

func (d *AccountInMemoryDB) Len() int {
	return len(d.m)
}

func (d *AccountInMemoryDB) PayToAccount(sendAccountPtr *models.Account, receiveAccountPtr *models.Account, summ string) bool {
	if receiveAccountPtr.GetId() == sendAccountPtr.GetId() {
		return false
	}

	syncPaymentMutex.Lock()
	defer syncPaymentMutex.Unlock()

	if sendAccountPtr.GetAmount().Available(summ) && sendAccountPtr.GetAmount().Sub(summ) {
		if receiveAccountPtr.GetAmount().Add(summ) {
			return true
		} else {
			sendAccountPtr.GetAmount().Add(summ)
		}
	}

	return false
}