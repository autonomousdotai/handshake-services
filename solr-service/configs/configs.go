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

var ServicePort, _ = strconv.Atoi(getenv("SERVICE_PORT", "8081"))
var SolrHost = os.Getenv("SOLR_HOST")
var SolrPort, _ = strconv.Atoi(os.Getenv("SOLR_PORT"))
var SolrCollectionHandshake = os.Getenv("SOLR_COLLECTION_HANDSHAKE")
var SolrCollectionUser = os.Getenv("SOLR_COLLECTION_USER")
