package controller_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uma-31/switchboard/agent/domain/valueobject"
	"github.com/uma-31/switchboard/agent/interface/http/controller"
)

func TestComputerInfoController_GetComputerInfo(t *testing.T) {
	t.Parallel()

	newComputerInfo := func(id string, name string) *valueobject.ComputerInfo {
		computerID, err := valueobject.NewComputerID(id)
		require.NoError(t, err, "テストデータの生成に失敗しました。")

		computerName, err := valueobject.NewComputerName(name)
		require.NoError(t, err, "テストデータの生成に失敗しました。")

		return valueobject.NewComputerInfo(computerID, computerName)
	}

	type fields struct {
		computerInfo *valueobject.ComputerInfo
	}

	tests := []struct {
		name    string
		fields  fields
		want    *controller.ComputerInfoDTO
		wantErr bool
	}{
		{
			name:    "Controller の初期化時に設定した ComputerInfo の情報を取得できる。",
			fields:  fields{newComputerInfo("desktop", "デスクトップ")},
			want:    &controller.ComputerInfoDTO{ID: "desktop", Name: "デスクトップ"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := controller.NewComputerInfoController(tt.fields.computerInfo)

			got, err := c.GetComputerInfo()
			if (err != nil) != tt.wantErr {
				assert.Errorf(t, err, "ComputerInfoController.GetComputerInfo() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				assert.Fail(t, "ComputerInfoController.GetComputerInfo() mismatch (-want +got):\n"+diff)
			}
		})
	}
}
