package services

import (
    "fmt"
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "mime/multipart"
    "github.com/ninjadotorg/handshake-services/ipfs-service/config"
)

type IPFSService struct {}

func (s IPFSService) Upload(source *multipart.FileHeader) (string, error) { 
    conf := config.GetConfig() 
    endpoint := conf.GetString("ipfs_endpoint")

    file, fileErr := source.Open()
    if fileErr != nil {
        fmt.Println("Read file error: ", fileErr)
        return false, fileErr
    }

    defer file.Close()

    fb, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println("Read file error: ", err)
        return "", err
    }

    request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(fb))
  
    client := &http.Client{}
    response, err := client.Do(request)

    if err != nil {
        return "", err
    }

    b, _ := ioutil.ReadAll(response.Body)
    
    var data map[string]interface{}
    json.Unmarshal(b, &data)

    result, ok := data["status"]
    hash, _ := data["data"] 
    message, _ := data["message"]

    if ok && (float64(1) == result) {
        return hash, nil
    } else {
        return "", message
    }
}
