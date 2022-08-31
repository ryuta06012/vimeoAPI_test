package response

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadResponse(res *http.Response, video interface{}) (interface{}, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// jsonを構造体へデコード
	if err := json.Unmarshal(body, &video); err != nil {
		return nil, err
	}
	return &video, nil
}
