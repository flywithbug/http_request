package http_request

import (
	"fmt"
	"testing"
)

func TestApiRequest(t *testing.T) {
	req := NewRequest()
	url := ""
	req.SetCookies(map[string]string{})
	res, err := req.Get(url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	body, err := res.Body()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}
