package controller

import (
	"net"

	"github.com/uma-31/switchboard/agent/domain/valueobject"
)

// MAC アドレスの取得に失敗したことを示す例外。
type GetMacAddressFailedError struct {
	cause error
}

func (e *GetMacAddressFailedError) Error() string {
	if e.cause == nil {
		return "MAC アドレスの取得に失敗しました。"
	}

	return "MAC アドレスの取得に失敗しました: " + e.cause.Error()
}

// コンピュータの情報を表すDTO。
type ComputerInfoDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MacAddress string `json:"macAddress"`
}

// コンピュータの情報関連の機能を扱うためのコントローラー。
type ComputerInfoController struct {
	computerInfo *valueobject.ComputerInfo
}

// 新しいComputerInfoControllerを生成する。
func NewComputerInfoController(computerInfo *valueobject.ComputerInfo) *ComputerInfoController {
	return &ComputerInfoController{computerInfo}
}

// コンピュータの情報を取得する。
func (c *ComputerInfoController) GetComputerInfo() (*ComputerInfoDTO, error) {
	computerID := c.computerInfo.ID()
	computerName := c.computerInfo.Name()

	// NOTE:
	// 以下、仮実装のため MAC アドレスをここで取得しているが、
	// 本来は ComputerInfo に含めるか、UseCase で取得するべき。
	// また、この方法では任意のインターフェイスの MAC アドレスを取得できないため、
	// Config などで制御できるようにする必要がある。
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, &GetMacAddressFailedError{err}
	}

	var macAddress string

	for _, ifa := range ifas {
		address := ifa.HardwareAddr.String()

		if address != "" {
			macAddress = address

			break
		}
	}

	if macAddress == "" {
		return nil, &GetMacAddressFailedError{nil}
	}

	return &ComputerInfoDTO{
		ID:         computerID.Value(),
		Name:       computerName.Value(),
		MacAddress: macAddress,
	}, nil
}
