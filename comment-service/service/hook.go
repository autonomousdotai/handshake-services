package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ninjadotorg/handshake-services/comment-service/configs"
)

type HookService struct{}

// CommentCountHooks :
func (s HookService) CommentCountHooks(objectId string, commentNumber int) {
	jsonData := make(map[string]interface{})
	jsonData["objectId"] = objectId
	jsonData["commentNumber"] = commentNumber
	jsonValue, _ := json.Marshal(jsonData)

	b := bytes.NewBuffer(jsonValue)

	services := configs.DispatcherServiceUrl
	// Send all number of comment's event to services
	for _, value := range services {
		fmt.Printf("Start call to hook services: %s \n", value)
		endpoint := value
		request, _ := http.NewRequest("POST", string(endpoint), b)
		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		b, _ := ioutil.ReadAll(response.Body)

		var data map[string]interface{}
		json.Unmarshal(b, &data)
		fmt.Println("====== Result ======")
		fmt.Println(data)
		result, ok := data["status"]
		message, _ := data["message"]

		if !ok || !(float64(1) == result) {
			fmt.Println(errors.New(message.(string)))
			return
		}
	}
	return
}
