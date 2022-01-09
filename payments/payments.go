package payments

import (
	"math/big"

	"github.com/jeronimobarea/simple-ethereum-client/common"

	"github.com/jeronimobarea/simple-axieinfinity-client/utils"
)

func CalculatePaymentQuantity(
	balance, percent *big.Int, rules []int64,
) (quantity *big.Int, valid bool, err error) {
	var rulesSum int64
	for _, q := range rules {
		rulesSum += q
	}

	quantity, err = utils.BigIntPercentage(balance, percent)
	if err != nil {
		valid = false
		return
	}

	quantity.Add(quantity, big.NewInt(rulesSum))

	valid = common.SafeBalanceIsValid(quantity)
	if !valid {
		quantity = big.NewInt(0)
		return
	}
	return
}
