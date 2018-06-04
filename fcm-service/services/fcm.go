package services

import (
    "fmt"
    "errors"
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "github.com/ninjadotorg/handshake-services/fcm-service/config"
)

type FCMService struct {}

func (s FCMService) Send(jsonData map[string]interface{}) (bool, error) { 
    conf := config.GetConfig() 
    endpoint := conf.GetString("gcm_endpoint")
    serverKey := conf.GetString("gcm_server_key")

    jsonValue, _ := json.Marshal(jsonData) 

    request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonValue))
    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("Authorization", fmt.Sprintf("key=%s", serverKey))

    client := &http.Client{}
    response, err := client.Do(request)

    if err != nil {
        return false, err
    }

    b, _ := ioutil.ReadAll(response.Body)
    
    var data map[string]interface{}
    json.Unmarshal(b, &data)

    success, ok := data["success"]

    if ok && (float64(1) == success) {
        return true, nil
    } else {
        return false, errors.New(string(b[:]))
    }
}
