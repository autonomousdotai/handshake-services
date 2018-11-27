package utils

import (
	"github.com/ninjadotorg/handshake-services/comment-service/configs"
)

func CdnUrlFor(fileUrl string) string {
	if fileUrl == "" {
		return ""
	}
	result := ""
	if configs.CdnHttps == true {
		result += "https://"
	} else {
		result += "http://"
	}
	result += configs.CdnDomain + "/" + fileUrl
	return result
}

func UrlFor(fileUrl string) string {
	if fileUrl == "" {
		return ""
	}
	result := configs.GcUrl + "/" + fileUrl
	return result
}
