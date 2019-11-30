package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
 curl -X POST -H 'Content-type: application/json' --data '{"text":"Hello, World!"}' https://hooks.slack.com/services/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
*/

func SendHttpMessage(url string, m string) (string, error) {

	msg, _ := json.Marshal(map[string]string{
		"text": m,
	})
	resp, err := http.Post(url,
		"application/json",
		bytes.NewBuffer(msg))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("body %s\n", string(body))

	return string(body), err
	// ...
}
