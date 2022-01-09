package payments

import (
	"math/big"
	"testing"

	"github.com/tj/assert"
)

func Test_CalculatePaymentQuantity_CorrectInput(t *testing.T) {
	balance := big.NewInt(100)
	percent := big.NewInt(10)
	rules := []int64{-30, 10}
	quantity, valid, err := CalculatePaymentQuantity(balance, percent, rules)
	assert.NoError(t, err, nil)
	assert.Equal(t, false, valid)
	assert.Equal(t, big.NewInt(0), quantity)
}

func Test_CalculatePaymentQuantity_NilBalance(t *testing.T) {
	percent := big.NewInt(10)
	rules := []int64{-30, 10}
	_, valid, err := CalculatePaymentQuantity(nil, percent, rules)
	assert.Error(t, err, nil)
	assert.Equal(t, false, valid)
}

func Test_CalculatePaymentQuantity_ZeroPercent(t *testing.T) {
	balance := big.NewInt(100)
	percent := big.NewInt(0)
	rules := []int64{10}
	quantity, valid, err := CalculatePaymentQuantity(balance, percent, rules)
	assert.NoError(t, err, nil)
	assert.Equal(t, true, valid)
	assert.Equal(t, big.NewInt(10), quantity)
}
