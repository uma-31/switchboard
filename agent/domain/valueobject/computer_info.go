package valueobject

import (
	"regexp"
)

// コンピューターの ID のフォーマットが不正であることを示す例外。
type InvalidComputerIDFormatError struct {
	value string
}

func (e *InvalidComputerIDFormatError) Error() string {
	return "コンピューターの ID は半角英数字もしくは'-'か'_'のみ含めることができます: " + e.value
}

// コンピュータの ID。
type ComputerID struct {
	value string
}

// ComputerID のインスタンスを生成する。
func NewComputerID(value string) (*ComputerID, error) {
	if value == "" {
		return nil, &InvalidComputerIDFormatError{"(nil)"}
	}

	re := regexp.MustCompile(`^[\w\d_-]+$`)
	if !re.MatchString(value) {
		return nil, &InvalidComputerIDFormatError{value}
	}

	return &ComputerID{value}, nil
}

// コンピュータの ID を取得する。
func (id *ComputerID) Value() string {
	return id.value
}

// コンピュータの名前。
type ComputerName struct {
	value string
}

// ComputerName のインスタンスを生成する。
func NewComputerName(value string) (*ComputerName, error) {
	return &ComputerName{value}, nil
}

// コンピュータの名前を取得する。
func (name *ComputerName) Value() string {
	return name.value
}

// コンピュータの情報。
type ComputerInfo struct {
	id   *ComputerID
	name *ComputerName
}

// ComputerInfo のインスタンスを生成する。
func NewComputerInfo(id *ComputerID, name *ComputerName) *ComputerInfo {
	return &ComputerInfo{id, name}
}

// コンピュータの ID を取得する。
func (info *ComputerInfo) ID() ComputerID {
	return *info.id
}

// コンピュータの名前を取得する。
func (info *ComputerInfo) Name() ComputerName {
	return *info.name
}
