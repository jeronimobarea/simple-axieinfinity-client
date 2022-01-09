package axieinfinity

import (
	axieCommon "github.com/jeronimobarea/simple-axieinfinity-client/common"
	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
	"github.com/stretchr/testify/mock"
)

type MockApi struct {
	mock.Mock
}

func (api *MockApi) Claim(address simpleCommon.PrivateAddress) (
	resp *ClaimResponse, err error,
) {
	args := api.Called(address)
	if args.Get(0) != nil {
		resp = args.Get(0).(*ClaimResponse)
	}
	return resp, args.Error(1)
}

func (api *MockApi) GetClaimableSlp(address axieCommon.RoninAddress) (
	resp *ClaimableResponse, err error,
) {
	args := api.Called(address)
	if args.Get(0) != nil {
		resp = args.Get(0).(*ClaimableResponse)
	}
	return resp, args.Error(1)
}

func (api *MockApi) GetClaimableBatchSlp(addresses []axieCommon.RoninAddress) (
	resp *ClaimableBatchResponse, err error,
) {
	args := api.Called(addresses)
	if args.Get(0) != nil {
		resp = args.Get(0).(*ClaimableBatchResponse)
	}
	return resp, args.Error(1)
}

func (api *MockApi) MakePayment(
	address axieCommon.RoninAddress,
	token string,
	scholarPaymentData *ScholarPaymentData,
) (
	resp *PaymentResponse,
	err error,
) {
	args := api.Called(address, token, scholarPaymentData)
	if args.Get(0) != nil {
		resp = args.Get(0).(*PaymentResponse)
	}
	return resp, args.Error(1)
}

func (api *MockApi) GetQR(address simpleCommon.PrivateAddress) (
	resp *QrResponse, err error,
) {
	args := api.Called(address)
	if args.Get(0) != nil {
		resp = args.Get(0).(*QrResponse)
	}
	return resp, args.Error(1)
}

func (api *MockApi) GetBalance(address axieCommon.RoninAddress, token string) (
	resp *BalanceResponse, err error,
) {
	args := api.Called(address, token)
	if args.Get(0) != nil {
		resp = args.Get(0).(*BalanceResponse)
	}
	return resp, args.Error(1)
}
