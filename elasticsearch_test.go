package elasticsearch

import (
	"testing"

	elasticsearch "github.com/fdv/crappy-go-elasticsearch-pkg"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	cluster := elasticsearch.Init("http://localhost", 9200)

	assert.Equal(t, cluster.URI, "http://localhost", "Cluster URI Should be http://localhost")
	assert.Equal(t, cluster.Port, uint16(9200), "Cluster port should be 9200")
}

func TestGeturi(t *testing.T) {
	cluster := elasticsearch.Init("http://localhost", 9200)

	assert.Equal(t, cluster.Geturi(), "http://localhost", "Cluster URI Should be http://localhost")
}

func TestGetport(t *testing.T) {
	cluster := elasticsearch.Init("http://localhost", 9200)

	assert.Equal(t, cluster.Getport(), uint16(9200), "Cluster port should be 9200")
}

func TestGeturl(t *testing.T) {
	cluster := elasticsearch.Init("http://localhost", 9200)

	assert.Equal(t, cluster.Geturl(), "http://localhost:9200", "Cluster URL should be http://localhost:9200")

}
