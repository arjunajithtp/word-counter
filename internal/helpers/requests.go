package helpers

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func GetData(url string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte(""), fmt.Errorf("error while trying to create request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte(""), fmt.Errorf("error while trying to process request: %v", err)
	}

	defer res.Body.Close()
	if res.StatusCode > 300 {
		return []byte(""), fmt.Errorf("invalid response: error: %v : %v", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), fmt.Errorf("error while trying to get body from response: %v", err)
	}

	return body, nil
}
