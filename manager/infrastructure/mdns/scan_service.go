package mdns

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/hashicorp/mdns"
	"github.com/uma-31/switchboard/manager/domain/entity"
)

const numEntries = 10

// コンピュータ情報の取得に失敗したことを示すエラー。
type FetchComputerInfoFailedError struct {
	cause error
}

func (e *FetchComputerInfoFailedError) Error() string {
	return "コンピュータの情報取得に失敗しました: " + e.cause.Error()
}

// mDNS クエリが失敗したことを示すエラー。
type QueryFailedError struct {
	cause error
}

func (e *QueryFailedError) Error() string {
	return "mDNS クエリが失敗しました: " + e.cause.Error()
}

type foundComputer struct {
	IP   string
	Port int
}

type computerInfo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MacAddress string `json:"macAddress"`
}

// NOTE: 仮実装なので後で別のパッケージに切り出す。
func fetchComputerInfo(computer foundComputer) (*entity.ComputerEntity, error) {
	apiPath := fmt.Sprintf(
		"http://%s/info",
		net.JoinHostPort(computer.IP, strconv.Itoa(computer.Port)),
	)

	res, err := http.Get(apiPath)
	if err != nil {
		return nil, &FetchComputerInfoFailedError{err}
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &FetchComputerInfoFailedError{err}
	}

	var info computerInfo

	if err := json.Unmarshal(body, &info); err != nil {
		return nil, &FetchComputerInfoFailedError{err}
	}

	return entity.NewComputerEntity(
		info.ID,
		info.Name,
		info.MacAddress,
	), nil
}

// mDNS によってコンピュータを検索するサービス。
type ScanComputerService struct{}

// ScanComputerService のインスタンスを生成する。
func NewScanComputerService() *ScanComputerService {
	return &ScanComputerService{}
}

// ネットワーク上のコンピュータを検索する。
func (s *ScanComputerService) ScanComputers() ([]*entity.ComputerEntity, error) {
	var foundComputers []foundComputer

	var mutex sync.Mutex

	entriesCh := make(chan *mdns.ServiceEntry, numEntries)

	go func() {
		for entry := range entriesCh {
			// NOTE: これも並列化したい
			mutex.Lock()
			foundComputers = append(foundComputers, foundComputer{
				IP:   entry.AddrV4.String(),
				Port: entry.Port,
			})
			mutex.Unlock()
		}
	}()

	params := mdns.DefaultParams("_agent._switchboard._tcp")

	params.Entries = entriesCh
	params.DisableIPv6 = true

	err := mdns.Query(params)
	if err != nil {
		return nil, &QueryFailedError{err}
	}

	close(entriesCh)

	computers := make([]*entity.ComputerEntity, 0, len(foundComputers))

	for _, computer := range foundComputers {
		computer, err := fetchComputerInfo(computer)
		if err != nil {
			continue
		}

		computers = append(computers, computer)
	}

	return computers, nil
}
