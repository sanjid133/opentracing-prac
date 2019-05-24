package http

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func Do(req *http.Request) ([]byte, error)  {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Status code %v, Body %v", resp.StatusCode, resp.Body)
	}
	return body, nil
}
