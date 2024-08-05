package config_test

import (
	"os"
	"path"
	"testing"

	"github.com/denisbrodbeck/machineid"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uma-31/switchboard/agent/domain/valueobject"
	"github.com/uma-31/switchboard/agent/infrastructure/config"
	"github.com/uma-31/switchboard/agent/infrastructure/http/gin"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	cwd, err := os.Getwd()
	require.NoErrorf(t, err, "ワーキングディレクトリのパスの取得に失敗しました。")

	machineID, err := machineid.ID()
	require.NoError(t, err, "マシン ID の取得に失敗しました。")

	getTestFilePath := func(fileName string) config.FilePath {
		return config.FilePath(path.Join(cwd, "testdata", fileName))
	}

	newComputerInfo := func(id string, name string) *valueobject.ComputerInfo {
		computerID, err := valueobject.NewComputerID(id)
		require.NoError(t, err, "テストデータの生成に失敗しました。")

		computerName, err := valueobject.NewComputerName(name)
		require.NoError(t, err, "テストデータの生成に失敗しました。")

		return valueobject.NewComputerInfo(computerID, computerName)
	}

	type args struct {
		configFilePath config.FilePath
	}

	tests := []struct {
		name    string
		args    args
		want    *config.Config
		wantErr bool
	}{
		{
			name: "全ての項目が正しく設定されている場合、正しい設定情報を取得できる。",
			args: args{getTestFilePath("complete.yaml")},
			want: &config.Config{
				ComputerInfo: newComputerInfo("desktop", "デスクトップ"),
				Port:         gin.NewServerPort(11331),
			},
			wantErr: false,
		},
		{
			name: "必須でない項目が null である場合、自動生成された値を含む設定情報を取得できる。",
			args: args{getTestFilePath("minimum_with_null.yaml")},
			want: &config.Config{
				ComputerInfo: newComputerInfo(machineID, machineID),
				Port:         gin.NewServerPort(11331),
			},
			wantErr: false,
		},
		{
			name: "必須でない項目が指定されていない場合、自動生成された値を含む設定情報を取得できる。",
			args: args{getTestFilePath("minimum_without_null.yaml")},
			want: &config.Config{
				ComputerInfo: newComputerInfo(machineID, machineID),
				Port:         gin.NewServerPort(11331),
			},
			wantErr: false,
		},
		{
			name:    "ポート番号の指定が不正な場合、エラーが発生する。",
			args:    args{getTestFilePath("invalid_port.yaml")},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "コンピュータ ID の指定が不正な場合、エラーが発生する。",
			args:    args{getTestFilePath("invalid_computer_id.yaml")},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "設定ファイルが存在しない場合、エラーが発生する。",
			args:    args{getTestFilePath("not_found.yaml")},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := config.Load(tt.args.configFilePath)
			if (err != nil) != tt.wantErr {
				assert.Errorf(t, err, "Load() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			cmpOption := cmp.AllowUnexported(
				//exhaustruct:ignore
				config.Config{},
				valueobject.ComputerInfo{},
				valueobject.ComputerID{},
				valueobject.ComputerName{},
				gin.ServerPort{},
			)

			if diff := cmp.Diff(got, tt.want, cmpOption); diff != "" {
				assert.Fail(t, "Load() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
