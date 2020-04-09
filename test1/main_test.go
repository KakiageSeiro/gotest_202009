package test1

import (
	"testing"
)

func Test_retrunError(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		got     error
		v       int
	}{
		{"nilが返却される",
			false,
			nil,
			1,
		},
		{
			"Err1の内容が返却される",
			true,
			Err1,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := returnErr(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("retrunError() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != tt.got {
				t.Errorf("retrunError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
