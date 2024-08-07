package controller

// コンピュータの情報を表す DTO。
type ComputerDTO struct {
	// コンピュータの ID。
	ID string `json:"id"`
	// コンピュータの名前。
	Name string `json:"name"`
	// コンピュータの MAC アドレス。
	MacAddress string `json:"macAddress"`
}
