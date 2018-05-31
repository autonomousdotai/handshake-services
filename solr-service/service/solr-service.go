package service

import (
	"github.com/rtt/Go-Solr"
	"fmt"
	"log"
)

type SolrService struct {
	Host       string
	Port       int
	Core       string
	Connection *solr.Connection
}

func (solrService SolrService) Init(host string, port int, core string) SolrService {
	solrService.Host = host
	solrService.Port = port
	solrService.Core = core
	c, _ := solr.Init(solrService.Host, solrService.Port, solrService.Core)
	solrService.Connection = c
	return solrService
}

func (solrService SolrService) Select(q *solr.Query) (*solr.SelectResponse, error) {
	log.Println("solr query string")
	log.Println(q.String())
	res, err := solrService.Connection.Select(q)
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}

func (solrService SolrService) Update(document map[string]interface{}) (*solr.UpdateResponse, error) {
	res, err := solrService.Connection.Update(document, true)
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}
