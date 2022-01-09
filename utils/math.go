package utils

import (
	"errors"
	"math/big"
)

func BigIntPercentage(x, y *big.Int) (*big.Int, error) {
	if x == nil || y == nil {
		return nil, errors.New("values cannot be nil")
	}

	percent := big.NewInt(100)
	var z big.Int
	z.Div(x, percent)
	z.Mul(&z, y)
	return &z, nil
}

func SubtractBigInt(x, y *big.Int) (*big.Int, error) {
	if x == nil || y == nil {
		return nil, errors.New("inputs can not be nil")
	}
	var z big.Int
	z.Sub(x, y)
	return &z, nil
}
