package request

import (
	"fmt"
	"testing"
)

func TestGetWithWbi(t *testing.T) {
	res, err := GetUserDetail(648113003)
	fmt.Println(string(res))
	if err != nil {
		t.Error(err)
	}
	t.Log(string(res))
}
