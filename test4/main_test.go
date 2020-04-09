package test4

import "testing"
import "fmt"

func Test_deferTest(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		want error
	}{
		{
			fmt.Sprintf("%sの内容が出力される事", ErrFire),
			true,
			ErrFire,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := deferTest()
			if (err != nil) != tt.wantErr {
				t.Errorf("deferTest() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != ErrFire {
				t.Errorf("%sを期待しているが、%s", ErrFire, err)
			}
		})
	}
}
