package tests

import (
	"../models"
	"../modules/billing"
	"testing"
)

func TestGetAmount(t *testing.T) {
	account1 := models.Account{}
	account1.Init(1, "Mark")
	account1.SetAmount("2454354358735793579823794723875982738472387423759872385792374283748723947238749237.33")

	res := account1.GetAmount().String()
	testRes := "2454354358735793579823794723875982738472387423759872385792374283748723947238749237.33"

	if res != testRes {
		t.Error(
			"set amount failed. expects " + testRes + " got ", res)
	}
}

func TestAddMoney(t *testing.T) {
	account1 := models.Account{}
	account1.Init(1, "Mark")
	account1.SetAmount("2454354358735793579823794723875982738472387423759872385792374283748723947238749237.33")

	account1.GetAmount().Add("22.02")

	res := account1.GetAmount().String()
	testRes := "2454354358735793579823794723875982738472387423759872385792374283748723947238749259.35"

	if res != testRes {
		t.Error(
			"add summ for amount failed. expects " + testRes + " got ", res)
	}
}

func TestSubMoney(t *testing.T) {
	account1 := models.Account{}
	account1.Init(1, "Mark")
	account1.SetAmount("2454354358735793579823794723875982738472387423759872385792374283748723947238749237.33")

	account1.GetAmount().Sub("3.07")

	res := account1.GetAmount().String()
	testRes := "2454354358735793579823794723875982738472387423759872385792374283748723947238749234.26"

	if res != testRes {
		t.Error(
			"sub summ for amount failed. expects " + testRes + " got ", res)
	}
}

func TestPay(t *testing.T) {
	account1 := models.Account{}
	account1.Init(1, "Mark")
	account1.SetAmount("2454354358735793579823794723875982738472387423759872385792374283748723947238749237.33")

	account2 := models.Account{}
	account2.Init(2, "Archi")
	account2.SetAmount("1000")

	billing.PayToAccount(&account1, &account2, "1.01")

	res := account1.GetAmount().String()
	testRes := "2454354358735793579823794723875982738472387423759872385792374283748723947238749236.32"

	if res != testRes {
		t.Error(
			"Account1: payment failed. amount after operation expects " + testRes + " got ", res)
	}

	res = account2.GetAmount().String()
	testRes = "1001.01"

	if res != testRes {
		t.Error(
			"Account2: payment failed. amount after operation expects " + testRes + " got ", res)
	}
}