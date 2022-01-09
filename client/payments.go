package axieinfinity

import (
	"github.com/ethereum/go-ethereum/core/types"
	axieCommon "github.com/jeronimobarea/simple-axieinfinity-client/common"
	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
)

type (
	ScholarPaymentData struct {
		Name           string                      `json:"name"`
		TeamAddress    simpleCommon.PrivateAddress `json:"teamAddress"`
		ScholarAddress axieCommon.RoninAddress     `json:"scholarAddress"`
		ScholarPercent int                         `json:"scholarPercent"`
		ScholarPayout  []int                       `json:"scholarPayout"`
	}

	PaymentData struct {
		ManagerAddress axieCommon.RoninAddress `json:"managerAddress"`
		Scholars       []*ScholarPaymentData   `json:"scholars"`
	}

	PaymentResponse struct {
		Failed         int                 `json:"failed"`
		Pending        int                 `json:"pending"`
		Results        []map[string]string `json:"results"`
		ScholarPayment *types.Transaction  `json:"scholarPayment"`
		ManagerPayment *types.Transaction  `json:"managerPayment"`
	}
)

func (svc *axiebuddyService) MakePayment(
	manager axieCommon.RoninAddress,
	token string,
	scholar *ScholarPaymentData,
) (response *PaymentResponse, err error) {
	return svc.API.MakePayment(manager, token, scholar)
}
