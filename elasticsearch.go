package elasticsearch

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Cluster defines an Elasticsearch cluster
type Cluster struct {
	URI  string
	Port uint16
	URL  string
}

// Init initialises an Elasticsearch cluster
func Init(uri string, port uint16) Cluster {
	url := strings.Join([]string{uri, string(port)}, ":")
	return (Cluster{URI: uri, Port: port, URL: url})
}

// Geturi returns the cluster URI
func (c *Cluster) Geturi() string {
	return c.URI
}

// Getport returns the cluster port
func (c *Cluster) Getport() uint16 {
	return c.Port
}

// Geturl returns the cluster URL
func (c *Cluster) Geturl() string {
	return c.URL
}

func (c *Cluster) buildgetquery(querystring ...string) string {
	uri := strings.Join(querystring, "/")

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
