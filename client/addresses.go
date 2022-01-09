package axieinfinity

import (
	"errors"
	"math/big"

	axieCommon "github.com/jeronimobarea/simple-axieinfinity-client/common"
)

type BalanceResponse struct {
	Address axieCommon.RoninAddress `json:"address"`
	Token   string                  `json:"token"`
	Balance *big.Int                `json:"quantity"`
}

func (svc *axiebuddyService) GetBalance(
	address axieCommon.RoninAddress, token string,
) (response *BalanceResponse, err error) {
	if !address.Validate() {
		err = errors.New("invalid ethereum format address")
		return
	}
	return svc.API.GetBalance(address, token)
}
