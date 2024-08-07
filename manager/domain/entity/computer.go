package entity

// 操作対象のコンピュータを表すエンティティ。
//
// NOTE: それぞれのフィールドの型を Value Object に変更したい。
type ComputerEntity struct {
	// コンピュータの ID。
	ID string

	// コンピュータの名前。
	Name string

	// コンピュータの MAC アドレス。
	MacAddress string
}

// ComputerEntity のインスタンスを生成する。
func NewComputerEntity(id, name, macAddress string) *ComputerEntity {
	return &ComputerEntity{id, name, macAddress}
}
