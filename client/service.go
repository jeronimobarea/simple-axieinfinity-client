package axieinfinity

import (
	axieCommon "github.com/jeronimobarea/simple-axieinfinity-client/common"
	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
)

type (
	Resources struct {
		API Api
	}

	axiebuddyService struct {
		*Resources
	}

	Service interface {
		Claim(address simpleCommon.PrivateAddress) (*ClaimResponse, error)
		GetClaimableSlp(address axieCommon.RoninAddress) (*ClaimableResponse, error)
		GetClaimableBatchSlp(addresses []axieCommon.RoninAddress) (*ClaimableBatchResponse, error)
		MakePayment(manager axieCommon.RoninAddress, token string, scholar *ScholarPaymentData) (*PaymentResponse, error)
		GetQR(address simpleCommon.PrivateAddress) (*QrResponse, error)
		GetBalance(address axieCommon.RoninAddress, token string) (*BalanceResponse, error)
	}
)

func NewService(res *Resources) Service {
	return &axiebuddyService{res}
}
