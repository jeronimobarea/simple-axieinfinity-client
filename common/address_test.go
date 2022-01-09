package common

import (
	"testing"

	"github.com/tj/assert"
)

/*
 * GetTokenAddress
 */
func Test_GetTokenAddress_CorrectInputTicker(t *testing.T) {
	contract, err := GetTokenAddress(SLP)
	assert.NoError(t, err)
	assert.Equal(t, SlpContract.ToCommonAddress(), contract)
}

func Test_GetTokenAddress_CorrectInputAddress(t *testing.T) {
	contract, err := GetTokenAddress(SlpContract.ToString())
	assert.NoError(t, err)
	assert.Equal(t, SlpContract.ToCommonAddress(), contract)
}

func Test_GetTokenAddress_EmptyInput(t *testing.T) {
	_, err := GetTokenAddress("")
	assert.Error(t, err)
}
