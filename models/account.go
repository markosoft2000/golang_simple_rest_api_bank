package models

type Account struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Amount Money  `json:"Amount"`
}

func (u *Account) Init(id int, name string) {

	validateName(name)

	u.SetId(id)
	u.SetName(name)
}

func validateId(id int)  {
	if id <= 0 {
		panic("Id must be greater then 0")
	}
}

func validateName(name string)  {
	if name == "" {
		panic("Id must be greater then 0")
	}
}

func (u *Account) GetId() int {
	return u.Id
}

func (u *Account) SetId(id int) {
	validateId(id)
	u.Id = id
}

func (u *Account) GetName() string {
	return u.Name
}

func (u *Account) SetName(name string) {
	validateName(name)
	u.Name = name
}

func (u *Account) GetAmount() *Money {
	return &u.Amount
}

func (u *Account) GetAmountAsString() string {
	return u.Amount.String()
}

func (u *Account) SetAmount(amount string) error {
	return u.Amount.Init(amount)
}