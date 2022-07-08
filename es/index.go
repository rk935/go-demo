package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)
	cfg := elasticsearch.Config{
		Addresses: []string{"http://127.0.0.1:9200"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Millisecond,
			DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
	log.Println(strings.Repeat("~", 37))

	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	params := `{
	 "settings": {
	   "index": {
	     "number_of_shards": "3",
	     "number_of_replicas": "2"
	   }
	 },
	 "mappings": {
	   "properties": {
	     "id": {
	       "type": "integer"
	     },
	     "name": {
	       "type": "text",
	       "analyzer": "ik_max_word",
		    "search_analyzer":"ik_smart"
	     },
	     "hobby": {
	       "type": "text"
	     }
	   }
	 }
	}`

	req := esapi.IndicesCreateRequest{
		Index:   "cph3",
		Body:    strings.NewReader(params),
		Timeout: time.Second * 3,
		Header:  nil,
	}

	response, err := req.Do(context.Background(), es)
	if err != nil {
		fmt.Println("create index err111: ", err.Error())
		return
	}

	if response.IsError() {
		fmt.Println("create index err222: ", response.String())
		return
	}
	defer response.Body.Close()
	var r map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	} else {
		// Print the response status and indexed document version.
		fmt.Printf("r: %+v", r)
	}
}
