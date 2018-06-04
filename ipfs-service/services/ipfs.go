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

type UploadService struct {}

func (s UploadService) Upload(source *multipart.FileHeader) (string, error) { 
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
        return false, err
    }

    request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(fb))
  
    client := &http.Client{}
    response, err := client.Do(request)

    if err != nil {
        return false, err
    }

    b, _ := ioutil.ReadAll(response.Body)
    
    var data map[string]interface{}
    json.Unmarshal(b, &data)

    result, ok := data["status"]
    
    if ok {
        return (float64(1) == result), nil
    } else {
        return false, nil
    }
}

func (u UploadService) View(hash string) (bytes[], error) {
    
}
