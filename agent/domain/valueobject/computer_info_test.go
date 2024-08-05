package valueobject_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uma-31/switchboard/agent/domain/valueobject"
)

func TestNewComputerID(t *testing.T) {
	t.Parallel()

	type args struct {
		value string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "値が ID として適切な場合、ComputerID のインスタンスは正しく生成できる。",
			args:    args{"DesktopPC-1_0"},
			wantErr: false,
		},
		{
			name:    "値が空文字の場合、ComputerID のインスタンスは生成できない。",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "値が半角英数字もしくは'-'か'_'以外の文字を含む場合、ComputerID のインスタンスは生成できない。",
			args:    args{"日本語は指定できないよ"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := valueobject.NewComputerID(tt.args.value)

			if tt.wantErr {
				assert.Nil(t, got)
				assert.Errorf(t, err, "NewComputerID() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Equal(t, got.Value(), tt.args.value)
				assert.NoError(t, err, "NewComputerID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
