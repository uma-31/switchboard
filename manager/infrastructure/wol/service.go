package wol

import (
	"net"

	"github.com/uma-31/switchboard/manager/domain/entity"
)

const wolPacketSize = 102

const magicPacketPrefix = 0xff

// WoL のマジックパケット生成に失敗したことを示すエラー。
type BuildWoLPacketFailedError struct {
	cause error
}

func (e *BuildWoLPacketFailedError) Error() string {
	return "WoL パケットの生成に失敗しました: " + e.cause.Error()
}

// WoL パケットの送信に失敗したことを示すエラー。
type SendWoLPacketFailedError struct {
	cause error
}

func (e *SendWoLPacketFailedError) Error() string {
	return "WoL パケットの送信に失敗しました: " + e.cause.Error()
}

func buildMagicPacket(macAddress string) ([]byte, error) {
	magicPacket := make([]byte, 0, wolPacketSize)

	magicPacket = append(
		magicPacket,
		magicPacketPrefix,
		magicPacketPrefix,
		magicPacketPrefix,
		magicPacketPrefix,
	)

	mac, err := net.ParseMAC(macAddress)
	if err != nil {
		return nil, &BuildWoLPacketFailedError{cause: err}
	}

	for range 16 {
		magicPacket = append(magicPacket, mac...)
	}

	return magicPacket, nil
}

// WoL によるコンピュータの起動機能を提供するサービス。
type WakeComputerService struct{}

// WoLService のインスタンスを生成する。
func NewWakeComputerService() *WakeComputerService {
	return &WakeComputerService{}
}

// コンピュータを起動する。
func (s *WakeComputerService) Wake(computer *entity.ComputerEntity) error {
	magicPacket, err := buildMagicPacket(computer.MacAddress)
	if err != nil {
		return err
	}

	udpAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
	if err != nil {
		return &SendWoLPacketFailedError{err}
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return &SendWoLPacketFailedError{err}
	}

	if _, err := conn.Write(magicPacket); err != nil {
		return &SendWoLPacketFailedError{err}
	}

	return nil
}
