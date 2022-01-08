package axieinfinity

import (
	axieCommon "github.com/jeronimobarea/simple-axieinfinity/common"
	simpleCommon "github.com/jeronimobarea/simple-ethereum/common"
)

type Api interface {
	Claim(address simpleCommon.PrivateAddress) (*ClaimResponse, error)
	GetClaimableSLP(address axieCommon.RoninAddress) (*ClaimableResponse, error)
	GetClaimableBatchSLP(addresses []axieCommon.RoninAddress) (*ClaimableBatchResponse, error)
	MakePayment(manager axieCommon.RoninAddress, token string, scholarPaymentData *ScholarPaymentData) (*PaymentResponse, error)
	GetQR(address simpleCommon.PrivateAddress) (*QrResponse, error)
	GetBalance(address axieCommon.RoninAddress, token string) (*BalanceResponse, error)
}
