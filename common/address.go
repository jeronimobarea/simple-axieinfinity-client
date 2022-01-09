package common

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
)

func GetTokenAddress(token string) (contract common.Address, err error) {
	if simpleCommon.ValidateAddress(token) {
		contract = common.HexToAddress(token)
		return
	}

	token = strings.ToLower(token)
	var rawContract simpleCommon.RawAddress
	switch token {
	case SLP:
		rawContract = SlpContract
	case AXS:
		rawContract = AxsContract
	case AXIE:
		rawContract = AxieContract
	case WETH:
		rawContract = WethContract
	case USDC:
		rawContract = UsdcContract
	case RON:
		rawContract = RonContract
	default:
		err = fmt.Errorf(
			"choose one of the valid tokens: [%s, %s, %s, %s, %s, %s] selected: %s",
			SLP, AXS, AXIE, WETH, USDC, RON, token,
		)
		return
	}
	contract = rawContract.ToCommonAddress()
	return
}

func CommonToRoninAddress(address common.Address) RoninAddress {
	tmp := strings.Replace(address.Hex(), "0x", "ronin:", 1)
	return RoninAddress(tmp)
}
