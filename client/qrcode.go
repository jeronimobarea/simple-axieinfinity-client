package axieinfinity

import (
	simpleCommon "github.com/jeronimobarea/simple-ethereum/common"
)

type QrResponse struct {
	QrCode []byte `json:"qrCode"`
}

func (svc *axiebuddyService) GetQR(
	team simpleCommon.PrivateAddress,
) (response *QrResponse, err error) {
	return svc.API.GetQR(team)
}
