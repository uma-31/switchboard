package controller

import (
	"github.com/uma-31/switchboard/agent/domain/valueobject"
)

// コンピュータの情報を表すDTO。
type ComputerInfoDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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

	return &ComputerInfoDTO{
		ID:   computerID.Value(),
		Name: computerName.Value(),
	}, nil
}
