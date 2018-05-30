package configs

import (
	"os"
	"strconv"
)

var ServicePort, _ = strconv.Atoi(os.Getenv("SERVICE_PORT"))
var SolrHost = os.Getenv("SOLR_HOST")
var SolrPort, _ = strconv.Atoi(os.Getenv("SOLR_PORT"))
var SolrCollectionHandshake = os.Getenv("SOLR_COLLECTION_HANDSHAKE")
var SolrCollectionUser = os.Getenv("SOLR_COLLECTION_USER")
