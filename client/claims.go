package axieinfinity

import (
	axieCommon "github.com/jeronimobarea/simple-axieinfinity-client/common"
	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
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

func (svc *axiebuddyService) GetClaimableSlp(
	address axieCommon.RoninAddress,
) (*ClaimableResponse, error) {
	return svc.API.GetClaimableSlp(address)
}

func (svc *axiebuddyService) GetClaimableBatchSlp(
	addresses []axieCommon.RoninAddress,
) (*ClaimableBatchResponse, error) {
	return svc.API.GetClaimableBatchSlp(addresses)
}
