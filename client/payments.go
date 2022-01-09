package axieinfinity

import (
	"github.com/ethereum/go-ethereum/core/types"
	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
	"math/big"

	axieCommon "github.com/jeronimobarea/simple-axieinfinity-client/common"
)

type (
	ScholarPaymentData struct {
		TeamAddress    simpleCommon.PrivateAddress `json:"teamAddress"`
		ScholarAddress axieCommon.RoninAddress     `json:"scholarAddress"`
		ScholarPercent *big.Int                    `json:"scholarPercent"`
		ScholarPayout  []int64                     `json:"scholarPayout"`
	}

	PaymentData struct {
		ManagerAddress axieCommon.RoninAddress `json:"managerAddress"`
		Scholars       []*ScholarPaymentData   `json:"scholars"`
	}

	PaymentResponseResult struct {
		Error       error              `json:"error"`
		Transaction *types.Transaction `json:"transaction"`
	}

	PaymentResponse struct {
		Results map[axieCommon.RoninAddress]*PaymentResponseResult `json:"results"`
	}
)

func (svc *axiebuddyService) MakePayment(
	manager axieCommon.RoninAddress,
	token string,
	scholar *ScholarPaymentData,
) (response *PaymentResponse, err error) {
	return svc.API.MakePayment(manager, token, scholar)
}
