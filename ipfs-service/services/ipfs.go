package services

import (
    "io"
    "fmt"
    "errors"
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
        return "", fileErr
    }
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, partErr := writer.CreateFormFile("path", source.Filename)

    if partErr != nil {
        return "", partErr
    }

    _, copyErr := io.Copy(part, file)

    if copyErr != nil {
        return "", copyErr
    }

    writerErr := writer.Close()
    
    if writerErr != nil {
        return "", writerErr
    }

    endpoint = fmt.Sprintf("%s/api/v0/add", endpoint)

    request, _ := http.NewRequest("POST", endpoint, body)
    request.Header.Add("Content-Type", writer.FormDataContentType())

    client := &http.Client{}
    response, err := client.Do(request)

    if err != nil {
        return "", err
    }

    b, _ := ioutil.ReadAll(response.Body)
   
    fmt.Println(string(b))

    var data map[string]interface{}
    json.Unmarshal(b, &data)

    fmt.Println(data)

    hash, ok := data["Hash"] 

    if ok {
        return hash.(string), nil
    } else {
        return "", errors.New("Upload fail.")
    }
}

func (s IPFSService) View(hash string) ([]byte, error) { 
    conf := config.GetConfig() 
    endpoint := conf.GetString("ipfs_endpoint")

    endpoint = fmt.Sprintf("%s/api/v0/cat?arg=%s", endpoint, hash)

    request, _ := http.NewRequest("GET", endpoint, nil)

    client := &http.Client{}
    response, err := client.Do(request)

    if err != nil {
        return nil, err
    }

    defer response.Body.Close()
    if response.StatusCode != http.StatusOK {
        return nil, errors.New(response.Status)
    }

    var data bytes.Buffer 
    _, readErr := io.Copy(&data, response.Body)

    if readErr != nil {
        return nil, readErr
    }

    return data.Bytes(), nil
}
