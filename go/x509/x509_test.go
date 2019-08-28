package x509

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"
)

func TestRsaDecrypt(t *testing.T) {
	got, _ := RsaEncrypt([]byte("hello"))
	b64 := base64.StdEncoding.EncodeToString(got)
	fmt.Println(b64)
}

func TestRsaEncrypt(t *testing.T) {
	type args struct {
		origData []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RsaEncrypt(tt.args.origData)
			if (err != nil) != tt.wantErr {
				t.Errorf("RsaEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RsaEncrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}