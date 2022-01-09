package axieinfinity

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	simpleClient "github.com/jeronimobarea/simple-ethereum-client/client"
	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
	"github.com/skip2/go-qrcode"

	axieCommon "github.com/jeronimobarea/simple-axieinfinity-client/common"
	"github.com/jeronimobarea/simple-axieinfinity-client/payments"
	"github.com/jeronimobarea/simple-axieinfinity-client/slp"
	"github.com/jeronimobarea/simple-axieinfinity-client/utils"
)

type apiImplementation struct {
	client    simpleClient.Service
	ethClient *ethclient.Client
}

func NewApiImplementation(client *ethclient.Client) Api {
	if client == nil {
		panic("error client  cannot be nil")
	}

	return &apiImplementation{
		client: simpleClient.NewService(&simpleClient.Resources{
			API: simpleClient.NewApiImplementation(client),
		}),
		ethClient: client,
	}
}

func SafeNewApiImplementation(client *ethclient.Client) (Api, error) {
	if client == nil {
		return nil, errors.New("error client cannot be nil")
	}

	return &apiImplementation{
		client: simpleClient.NewService(&simpleClient.Resources{
			API: simpleClient.NewApiImplementation(client),
		}),
		ethClient: client,
	}, nil
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

	scholarAddress, err := simpleCommon.GetAddressFromPrivateKey(teamPk)
	if err != nil {
		return
	}

	balance, err := api.GetBalance(
		axieCommon.CommonToRoninAddress(scholarAddress), token,
	)
	if err != nil {
		return
	}

	if !simpleCommon.SafeBalanceIsValid(balance.Balance) {
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

	tokenContract, err := axieCommon.GetTokenAddress(token)
	if err != nil {
		return
	}

	scholarResponse, err := api.client.SimpleSendTransaction(
		scholarQuantity,
		teamPk,
		scholar.ScholarAddress.ToCommonAddress(),
		tokenContract,
	)
	if err != nil {
		return
	}
	resp.ScholarPayment = scholarResponse.Transaction

	managerQuantity, err := utils.SubtractBigInt(balance.Balance, scholarQuantity)
	if err != nil {
		return
	}

	managerResponse, err := api.client.SimpleSendTransaction(
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

	commonAddress, err := simpleCommon.GetAddressFromPrivateKey(teamPk)
	if err != nil {
		return
	}

	roninAddress := axieCommon.CommonToRoninAddress(commonAddress)
	claimable, err := api.GetClaimableSlp(roninAddress)
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

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, axieCommon.RoninGraphqlUri, nil)
	if err != nil {
		return
	}

	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", jwt)},
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var data map[string]map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	signature, ok := data["blockchain_relate"]["signature"]
	if !ok {
		err = errors.New("error checking dict keys")
		return
	}

	nonce, err := api.ethClient.NonceAt(context.Background(), commonAddress, nil)
	if err != nil {
		return
	}

	slpInstance, err := slp.NewSlp(commonAddress, api.ethClient)
	if err != nil {
		return
	}

	checkpoint, err := slpInstance.Checkpoint(
		&bind.TransactOpts{
			Nonce: big.NewInt(int64(nonce)),
		},
		commonAddress,
		big.NewInt(int64(claimable.Quantity)),
		nil,
		[]byte(signature),
	)
	if err != nil {
		return
	}

	err = api.ethClient.SendTransaction(context.Background(), checkpoint)
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
	tokenAddress, err := axieCommon.GetTokenAddress(token)
	if err != nil {
		return
	}

	balanceResp, err := api.client.SimpleCheckBalance(address.ToCommonAddress(), tokenAddress)
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

func (api *apiImplementation) GetClaimableSlp(
	address axieCommon.RoninAddress,
) (resp *ClaimableResponse, err error) {
	uri := fmt.Sprintf(axieCommon.RoninSlpClaimUri, address)

	httpResp, err := http.Get(uri)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	var claimable int
	var ok bool
	if claimable, ok = data["claimable_total"].(int); !ok {
		err = errors.New("claimable_total not present in response")
		return
	}
	resp = &ClaimableResponse{
		Quantity: claimable,
		Address:  address,
	}
	return
}

func (api *apiImplementation) GetClaimableBatchSlp(
	addresses []axieCommon.RoninAddress,
) (resp *ClaimableBatchResponse, err error) {
	for _, address := range addresses {
		resp.Total += 1
		var claimable *ClaimableResponse
		claimable, err = api.GetClaimableSlp(address)
		if err != nil {
			resp.Errors = append(resp.Errors, err)
			continue
		}
		resp.Processed = append(resp.Processed, claimable)
	}
	return
}
