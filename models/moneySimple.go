package models

import (
	"fmt"
	"strconv"
	"math"
)

type MoneySimple struct {
	value int
	precision int
}

func (m *MoneySimple) Init (amount string, precision int)  {
	m.value = convertStringToInt(amount, precision)
	m.precision = precision
}

func (m *MoneySimple) Set(amount string, precision int)  {
	m.Init(amount, precision)
}

func (m *MoneySimple) GetValueAsString(separator string) string {
	divider := int(math.Pow10(m.precision))
	return fmt.Sprintf("%d%s%d", m.value / divider, separator, m.value % divider)
}

func (m *MoneySimple) GetValue() float64 {
	return float64(m.value) / float64(m.precision)
}

func (m *MoneySimple) Add(summ string) bool {
	amount := convertStringToInt(summ, m.precision)
	result := m.value + amount

	if (result > m.value) == (amount > 0) {
		m.value = result

		return true
	}

	return false
}

func (m *MoneySimple) Sub(summ string) bool {
	amount := convertStringToInt(summ, m.precision)
	result := m.value - amount

	if (result < m.value) == (amount > 0) {
		m.value = result

		return true
	}

	return false
}

func convertStringToInt(s string, precision int) int {
	temp,_ := strconv.ParseFloat(s, 64)

	return int(temp * math.Pow10(precision))
}
