package configs

import (
	"os"
	"strconv"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

var ServicePort, _ = strconv.Atoi(getenv("SERVICE_PORT", "8088"))
var DB = getenv("DB", "")
var CdnDomain = os.Getenv("CDN_DOMAIN")
var CdnHttps, _ = strconv.ParseBool(os.Getenv("CDN_HTTPS"))
var DispatcherServiceUrl = getenv("DISPATCHER_SERVICE_URL", "localhost:8080")
var StorageServiceUrl = os.Getenv("STORAGE_SERVICE_URL")
var SolrServiceUrl = getenv("SOLR_SERVICE_URL", "localhost:6000")
var GcUrl = getenv("GC_URL", "")
var CommentHookServicesUrl = getenv("COMMENT_HOOOK_SERVICES_URL", "")
