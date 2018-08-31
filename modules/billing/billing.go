package billing

import (
	"sync"
	"../../models"
)

//some comments here

var transactionMutexMap = make(map[int] *sync.RWMutex)
var mapMutex = new(sync.RWMutex)

func getTransactionMutexOrCreate(id int) *sync.RWMutex {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	if !existsTransactionMutex(id) {
		transactionMutexMap[id] = new(sync.RWMutex)
	}

	return transactionMutexMap[id]
}

func existsTransactionMutex(id int) bool {
	_, ok := transactionMutexMap[id]
	return ok
}

//loop deadlock possible
//1->2 | 2->3 | 3->7 | 4->8
//1->2 | 2->3 | 3->1 | 4->8
//1->3 | 2->3 | 3->1 | 4->8
func lockMoneyTransaction(id1 int, id2 int) {
	getTransactionMutexOrCreate(id1).Lock()
	getTransactionMutexOrCreate(id2).Lock()
}

func unlockMoneyTransaction(id1 int, id2 int) {
	getTransactionMutexOrCreate(id1).Unlock()
	getTransactionMutexOrCreate(id2).Unlock()
}

func PayToAccount(sendAccountPtr *models.Account, receiveAccountPtr *models.Account, summ string) bool {
	if receiveAccountPtr.GetId() == sendAccountPtr.GetId() {
		return false
	}

	//transaction start
	lockMoneyTransaction(sendAccountPtr.GetId(), receiveAccountPtr.GetId())
	defer unlockMoneyTransaction(sendAccountPtr.GetId(), receiveAccountPtr.GetId())

	if !sendAccountPtr.GetAmount().Sub(summ) {
		return false
	}

	if !receiveAccountPtr.GetAmount().Add(summ) {
		sendAccountPtr.GetAmount().Add(summ)

		return false
	}
	//transaction finish
	return true
}

//sync payment method

var syncPaymentMutex = new(sync.Mutex)

//1->2 , 2->3 , 3->1 , 4->8
func PayToAccountSync(sendAccountPtr *models.Account, receiveAccountPtr *models.Account, summ string) bool {
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