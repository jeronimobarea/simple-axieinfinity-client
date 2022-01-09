package axieinfinity

import (
	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"

	axieCommon "github.com/jeronimobarea/simple-axieinfinity-client/common"
)

type Api interface {
	Claim(address simpleCommon.PrivateAddress) (*ClaimResponse, error)
	GetClaimableSlp(address axieCommon.RoninAddress) (*ClaimableResponse, error)
	GetClaimableBatchSlp(addresses []axieCommon.RoninAddress) (*ClaimableBatchResponse, error)
	MakePayment(manager axieCommon.RoninAddress, token string, scholarPaymentData *ScholarPaymentData) (*PaymentResponse, error)
	GetQR(address simpleCommon.PrivateAddress) (*QrResponse, error)
	GetBalance(address axieCommon.RoninAddress, token string) (*BalanceResponse, error)
}
