package axieinfinity

import (
	"fmt"

	axieAddress "github.com/jeronimobarea/simple-axieinfinity/address"
	"github.com/jeronimobarea/simple-axieinfinity/claims"
	axieCommon "github.com/jeronimobarea/simple-axieinfinity/common"
	"github.com/jeronimobarea/simple-axieinfinity/payments"
	axieMath "github.com/jeronimobarea/simple-axieinfinity/utils"
	simpleEthClient "github.com/jeronimobarea/simple-ethereum-client/client"
	simpleAddress "github.com/jeronimobarea/simple-ethereum/address"
	simpleCommon "github.com/jeronimobarea/simple-ethereum/common"
	"github.com/skip2/go-qrcode"
)

type apiImplementation struct {
	client simpleEthClient.Service
}

func NewApiImplementation() Api {
	return &apiImplementation{
		client: simpleEthClient.NewService(&simpleEthClient.Resources{
			API: simpleEthClient.NewApiImplementation(
				axieCommon.RoninProviderFree,
			),
		}),
	}
}

func (api *apiImplementation) MakePayment(
	manager axieCommon.RoninAddress,
	token string,
	scholar *ScholarPaymentData,
) (resp *PaymentResponse, err error) {
	teamPk, err := scholar.TeamAddress.ToECDSA()
	if err != nil {
		return
	}

	scholarAddress, err := simpleAddress.GetAddressFromPrivateKey(teamPk)
	if err != nil {
		return
	}

	balance, err := api.GetBalance(
		axieAddress.CommonToRoninAddress(scholarAddress), token,
	)
	if err != nil {
		return
	}

	if !simpleAddress.SafeBalanceIsValid(balance.Balance) {
		err = fmt.Errorf("balance is %s", balance)
		return
	}

	scholarQuantity, err := payments.CalculatePaymentQuantity(
		balance.Balance,
		scholar.ScholarPercent,
		scholar.ScholarPayout,
	)
	if err != nil {
		return
	}

	tokenContract, err := axieAddress.GetTokenAddress(token)
	if err != nil {
		return
	}

	scholarResponse, err := api.client.SendTransaction(
		scholarQuantity,
		teamPk,
		scholar.ScholarAddress.ToCommonAddress(),
		tokenContract,
	)
	if err != nil {
		return
	}
	resp.ScholarPayment = scholarResponse.Transaction

	managerQuantity, err := axieMath.SubtractBigInt(balance.Balance, scholarQuantity)
	if err != nil {
		return
	}

	managerResponse, err := api.client.SendTransaction(
		managerQuantity,
		teamPk,
		manager.ToCommonAddress(),
		tokenContract,
	)
	resp.ManagerPayment = managerResponse.Transaction
	return
}

func (api *apiImplementation) Claim(
	address simpleCommon.PrivateAddress,
) (result *ClaimResponse, err error) {
	teamPk, err := address.ToECDSA()
	if err != nil {
		return
	}

	commonAddress, err := simpleAddress.GetAddressFromPrivateKey(teamPk)
	if err != nil {
		return
	}

	claimable, err := claims.GetClaimableSlp(commonAddress)
	if err != nil {
		return
	}

	message, err := axieCommon.CreateRandomMessage()
	if err != nil {
		return
	}

	jwt, err := axieCommon.GetJWT(
		teamPk,
		message,
	)
	if err != nil {
		return
	}

	err = claims.ClaimSlp(
		nil,
		jwt,
		axieAddress.CommonToRoninAddress(commonAddress),
		claimable,
	)
	if err != nil {
		return
	}
	return
}

func (api *apiImplementation) GetQR(
	address simpleCommon.PrivateAddress,
) (result *QrResponse, err error) {
	message, err := axieCommon.CreateRandomMessage()
	if err != nil {
		return
	}

	teamPk, err := address.ToECDSA()
	if err != nil {
		return
	}

	jwt, err := axieCommon.GetJWT(
		teamPk,
		message,
	)
	if err != nil {
		return
	}

	qr, err := qrcode.Encode(
		jwt,
		qrcode.Medium,
		256,
	)
	if err != nil {
		return
	}

	result = &QrResponse{
		QrCode: qr,
	}
	return
}

func (api *apiImplementation) GetBalance(
	address axieCommon.RoninAddress,
	token string,
) (resp *BalanceResponse, err error) {
	tokenAddress, err := axieAddress.GetTokenAddress(token)
	if err != nil {
		return
	}

	balanceResp, err := api.client.CheckBalance(address.ToCommonAddress(), tokenAddress)
	if err != nil {
		return
	}

	resp = &BalanceResponse{
		Address: address,
		Token:   token,
		Balance: balanceResp.Balance,
	}
	return
}

func (api *apiImplementation) GetClaimableSLP(
	address axieCommon.RoninAddress,
) (resp *ClaimableResponse, err error) {
	claimable, err := claims.GetClaimableSlp(address.ToCommonAddress())
	if err != nil {
		return
	}

	resp = &ClaimableResponse{
		Quantity: claimable,
		Address:  address,
	}
	return
}

func (api *apiImplementation) GetClaimableBatchSLP(
	addresses []axieCommon.RoninAddress,
) (resp *ClaimableBatchResponse, err error) {
	for _, address := range addresses {
		resp.Total += 1
		var claimable *ClaimableResponse
		claimable, err = api.GetClaimableSLP(address)
		if err != nil {
			resp.Errors = append(resp.Errors, err)
			continue
		}
		resp.Processed = append(resp.Processed, claimable)
	}
	return
}
