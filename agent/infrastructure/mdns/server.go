package mdns

import (
	"github.com/hashicorp/mdns"
	"github.com/uma-31/switchboard/agent/domain/valueobject"
)

// mDNS サーバーの初期化に失敗したことを示すエラー。
type ServerInitializeFailedError struct {
	cause error
}

func (e *ServerInitializeFailedError) Error() string {
	return "mDNS サーバーの初期化に失敗しました: " + e.cause.Error()
}

// mDNS サーバーのシャットダウンに失敗したことを示すエラー。
type ServerShutdownFailedError struct {
	cause error
}

func (e *ServerShutdownFailedError) Error() string {
	return "mDNS サーバーのシャットダウンに失敗しました: " + e.cause.Error()
}

// mDNS サーバー。
type Server struct {
	mdnsService *mdns.Server
}

// Server のインスタンスを生成する。
func NewServer(computerID valueobject.ComputerID, port int) (*Server, error) {
	mdnsService, err := mdns.NewMDNSService(
		computerID.Value(),
		"_agent._switchboard._tcp",
		"",
		"",
		port,
		nil,
		nil,
	)
	if err != nil {
		return nil, &ServerInitializeFailedError{err}
	}

	mdnsServerConfig := &mdns.Config{
		Zone:              mdnsService,
		Iface:             nil,
		LogEmptyResponses: false,
	}

	mdnsServer, err := mdns.NewServer(mdnsServerConfig)
	if err != nil {
		return nil, &ServerInitializeFailedError{err}
	}

	return &Server{mdnsServer}, nil
}

// サーバーをシャットダウンする。
func (s *Server) Shutdown() error {
	err := s.mdnsService.Shutdown()
	if err != nil {
		return &ServerShutdownFailedError{err}
	}

	return nil
}
