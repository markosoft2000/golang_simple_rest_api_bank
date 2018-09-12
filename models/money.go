package models

import (
	//decimalPkg "github.com/ericlagergren/decimal"
	decimalPkg "github.com/shopspring/decimal"
)

type Money struct {
	Value decimalPkg.Decimal `json:"Value"`
}

func (m *Money) Init (amount string) error {
	if amount == "" {
		amount = "0.00"
	}

	value,err := decimalPkg.NewFromString(amount)

	if err != nil {
		return err
	}

	m.Value = value

	return nil
}

func (m *Money) Set(amount string) error {
	return m.Init(amount)
}

func (m *Money) String() string {
	return m.Value.String()
}

func (m *Money) Add(summ string) bool {
	operand, err := decimalPkg.NewFromString(summ)

	if err != nil {
		return false
	}

	m.Value = m.Value.Add(operand)
	//m.Value.Add(operand)

	return true
}

func (m *Money) Sub(summ string) bool {
	operand, err := decimalPkg.NewFromString(summ)

	if err != nil {
		return false
	}

	result := m.Value.Sub(operand)

	if result.IsNegative() {
		return false
	}

	m.Value = result

	return true
}

func (m *Money) Available(summ string) bool {
	operand, err := decimalPkg.NewFromString(summ)

	if err != nil {
		return false
	}

	return m.Value.GreaterThanOrEqual(operand)
}
