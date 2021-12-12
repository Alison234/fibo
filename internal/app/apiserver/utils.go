package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func parseRequest(r *http.Request) (Request, error) {
	result := Request{}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return result, fmt.Errorf("read request body, err=%v", err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, fmt.Errorf("unmarshal request body, err=%v", err)
	}

	return result, nil
}

func makeCacheKey(startIndex, endIndex int) string {
	return fmt.Sprintf("%v:%v", startIndex, endIndex)
}
