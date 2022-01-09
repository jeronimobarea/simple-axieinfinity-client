package utils

import (
	"math/big"
	"testing"

	"github.com/tj/assert"
)

func Test_BigIntPercentage_CorrectInput(t *testing.T) {
	balance := big.NewInt(100)
	percent := big.NewInt(25)
	res, err := BigIntPercentage(balance, percent)
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(25), res)
}

func Test_BigIntPercentage_ZeroPercent(t *testing.T) {
	balance := big.NewInt(100)
	percent := big.NewInt(0)
	res, err := BigIntPercentage(balance, percent)
	assert.NoError(t, err)
	assert.EqualValues(t, big.NewInt(0).BitLen(), res.BitLen())
}

func Test_BigIntPercentage_ZeroBalance(t *testing.T) {
	balance := big.NewInt(0)
	percent := big.NewInt(10)
	res, err := BigIntPercentage(balance, percent)
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(0), res)
}

func Test_BigIntPercentage_NilValue(t *testing.T) {
	_, err := BigIntPercentage(nil, nil)
	assert.Error(t, err)
}

func Test_SubtractBigInt_CorrectInput(t *testing.T) {
	x := big.NewInt(100)
	y := big.NewInt(10)
	res, err := SubtractBigInt(x, y)
	assert.NoError(t, err, nil)
	assert.Equal(t, big.NewInt(90), res)
}

func Test_SubtractBigInt_ZeroInput(t *testing.T) {
	x := big.NewInt(0)
	y := big.NewInt(0)
	res, err := SubtractBigInt(x, y)
	assert.NoError(t, err, nil)
	assert.Equal(t, big.NewInt(0), res)
}

func Test_SubtractBigInt_NilX(t *testing.T) {
	y := big.NewInt(20)
	_, err := SubtractBigInt(nil, y)
	assert.Error(t, err, nil)
}

func Test_SubtractBigInt_NilY(t *testing.T) {
	x := big.NewInt(20)
	_, err := SubtractBigInt(x, nil)
	assert.Error(t, err, nil)
}
