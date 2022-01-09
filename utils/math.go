package utils

import (
	"errors"
	"math/big"
)

func BigIntPercentage(balance *big.Int, percent int) (float32, error) {
	if balance == nil {
		return 0, errors.New("error balance can not be nil")
	}
	return Percentage(float32(balance.Int64()), percent), nil
}

func Percentage(balance float32, percentage int) float32 {
	return balance / 100 * float32(percentage)
}

func SubtractBigInt(x, y *big.Int) (*big.Int, error) {
	if x == nil || y == nil {
		return nil, errors.New("inputs can not be nil")
	}
	res := x.Int64() - y.Int64()
	return big.NewInt(res), nil
}
