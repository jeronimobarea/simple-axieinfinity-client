package common

import (
	"github.com/jeronimobarea/simple-ethereum-client/common"

	"github.com/jeronimobarea/simple-axieinfinity-client/utils"
)

var (
	AxieContract = common.RawAddress(
		utils.GetEnvFallback("AXIE_CONTRACT", "0x32950db2a7164ae833121501c797d79e7b79d74c"),
	)
	AxsContract = common.RawAddress(
		utils.GetEnvFallback("AXS_CONTRACT", "0x97a9107c1793bc407d6f527b77e7fff4d812bece"),
	)
	SlpContract = common.RawAddress(
		utils.GetEnvFallback("SLP_CONTRACT", "0xa8754b9fa15fc18bb59458815510e40a12cd2014"),
	)
	WethContract = common.RawAddress(
		utils.GetEnvFallback("WETH_CONTRACT", "0xc99a6a985ed2cac1ef41640596c5a5f9f4e19ef5"),
	)
	RonContract = common.RawAddress(
		utils.GetEnvFallback("RON_CONTRACT", ""),
	)
	UsdcContract = common.RawAddress(
		utils.GetEnvFallback("USDC_CONTRACT", "0x0b7007c13325c48911f73a2dad5fa5dcbf808adc"),
	)

	RoninProviderFree = utils.GetEnvFallback("RONIN_PROVIDER_FREE", "https://proxy.roninchain.com/free-gas-rpc")
	RoninProvider     = utils.GetEnvFallback("RONIN_PROVIDER", "https://api.roninchain.com/rpc")

	RoninSlpClaimUri = utils.GetEnvFallback("RONIN_SLP_CLAIM_URI", "https://game-api.skymavis.com/game-api/clients/%s/items/1")
	RoninGraphqlUri  = utils.GetEnvFallback("RONIN_GRAPHQL_URI", "https://graphql-gateway.axieinfinity.com/graphql")
)

const (
	AXIE string = "axie"
	AXS  string = "axs"
	SLP  string = "slp"
	WETH string = "weth"
	RON  string = "ron"
	USDC string = "usdc"
)
