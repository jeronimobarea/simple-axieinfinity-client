package common

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
)

type RoninAddress string

func (a RoninAddress) ToString() string {
	return string(a)
}

func (a RoninAddress) Validate() bool {
	if a == "" || !strings.Contains(a.ToString(), "ronin:") {
		return false
	}
	return simpleCommon.ValidateAddress(a.ToCommonAddress())
}

func (a *RoninAddress) ToRawAddress() simpleCommon.RawAddress {
	rawAddress := strings.Replace(a.ToString(), "ronin:", "0x", 1)
	return simpleCommon.RawAddress(rawAddress)
}

func (a *RoninAddress) ToCommonAddress() common.Address {
	rawAddress := strings.Replace(a.ToString(), "ronin:", "0x", 1)
	return common.HexToAddress(rawAddress)
}
