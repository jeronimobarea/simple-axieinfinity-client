package payments

import (
	"math/big"
	"testing"

	"github.com/tj/assert"
)

func Test_CalculatePaymentQuantity_CorrectInput(t *testing.T) {
	balance := big.NewInt(100)
	percent := 10
	rules := []int{-30, 10}
	quantity, err := CalculatePaymentQuantity(balance, percent, rules)
	assert.NoError(t, err, nil)
	assert.Equal(t, big.NewInt(0), quantity)
}

func Test_CalculatePaymentQuantity_NilBalance(t *testing.T) {
	percent := 10
	rules := []int{-30, 10}
	_, err := CalculatePaymentQuantity(nil, percent, rules)
	assert.Error(t, err, nil)
}

func Test_CalculatePaymentQuantity_ZeroPercent(t *testing.T) {
	balance := big.NewInt(100)
	percent := 0
	rules := []int{10}
	quantity, err := CalculatePaymentQuantity(balance, percent, rules)
	assert.NoError(t, err, nil)
	assert.Equal(t, big.NewInt(10), quantity)
}
