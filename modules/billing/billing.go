package billing

import (
	"sync"
	"fmt"
	"../../models"
)

//some comments here

var m = make(map[string] *sync.RWMutex)
var mapMutex = new(sync.RWMutex)

func genMutexKey(id1 int, id2 int) string {
	format := "%d-%d"

	if id1 <= id2 {
		return fmt.Sprintf(format, id1, id2)
	} else {
		return fmt.Sprintf(format, id2, id1)
	}
}

func lockMoneyTransaction(id1 int, id2 int) {
	key := genMutexKey(id1, id2)

	if _, ok := m[key]; !ok {
		mapMutex.Lock()
		m[key] = new(sync.RWMutex)
		mapMutex.Unlock()
	}

	m[key].Lock()
}

func unlockMoneyTransaction(id1 int, id2 int) {
	key := genMutexKey(id1, id2)
	m[key].Unlock()
}

func PayToAccount(sendAccountPtr *models.Account, receiveAccountPtr *models.Account, summ string) bool {
	if receiveAccountPtr.GetId() == sendAccountPtr.GetId() {
		return false
	}

	/*defer func() {
		recover()
	}()*/

	//transaction start
	lockMoneyTransaction(sendAccountPtr.GetId(), receiveAccountPtr.GetId())

	if !sendAccountPtr.GetAmount().Sub(summ) {
		return false
	}

	if !receiveAccountPtr.GetAmount().Add(summ) {
		sendAccountPtr.GetAmount().Add(summ)

		return false
	}

	unlockMoneyTransaction(sendAccountPtr.GetId(), receiveAccountPtr.GetId())
	//transaction finish
	return true
}