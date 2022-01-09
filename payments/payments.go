package payments

import (
	"math/big"

	"github.com/jeronimobarea/simple-axieinfinity-client/utils"
)

func CalculatePaymentQuantity(
	balance *big.Int, percent int, rules []int,
) (quantity *big.Int, err error) {
	var rulesSum int
	for _, q := range rules {
		rulesSum += q
	}

	floatQuantity, err := utils.BigIntPercentage(balance, percent)
	if err != nil {
		return
	}

	floatQuantity += float32(rulesSum)
	if floatQuantity <= 0 {
		quantity = big.NewInt(0)
		return
	}

	quantity = big.NewInt(int64(floatQuantity))
	return
}
