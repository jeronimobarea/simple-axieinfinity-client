package axieinfinity

import (
	axieCommon "github.com/jeronimobarea/simple-axieinfinity/common"
	simpleCommon "github.com/jeronimobarea/simple-ethereum/common"
)

type (
	ClaimResponse struct{}

	ClaimableResponse struct {
		Quantity int                     `json:"quantity"`
		Address  axieCommon.RoninAddress `json:"address"`
	}

	ClaimableBatchResponse struct {
		Processed []*ClaimableResponse `json:"processed"`
		Errors    []error              `json:"errors"`
		Total     int                  `json:"total"`
	}
)

func (svc *axiebuddyService) Claim(
	address simpleCommon.PrivateAddress,
) (*ClaimResponse, error) {
	return svc.API.Claim(address)
}

func (svc *axiebuddyService) GetClaimableSLP(
	address axieCommon.RoninAddress,
) (*ClaimableResponse, error) {
	return svc.API.GetClaimableSLP(address)
}

func (svc *axiebuddyService) GetClaimableBatchSLP(
	addresses []axieCommon.RoninAddress,
) (*ClaimableBatchResponse, error) {
	return svc.API.GetClaimableBatchSLP(addresses)
}
