package elasticsearch

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Cluster defines an Elasticsearch cluster
type Cluster struct {
	URI  string
	Port uint16
}

// Init initialises an Elasticsearch cluster
func Init(uri string, port uint16) Cluster {
	return (Cluster{URI: uri, Port: port})
}

// Geturi returns the cluster URI
func (c *Cluster) Geturi() string {
	return c.URI
}

// Getport returns the cluster port
func (c *Cluster) Getport() uint16 {
	return c.Port
}

func (c *Cluster) buildgetquery(querystring string) string {
	uri := fmt.Sprintf("%s:%d/%s", c.URI, c.Port, querystring)

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return (string(body))
}
