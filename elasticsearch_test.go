package elasticsearch

import (
	"testing"

	elasticsearch "github.com/fdv/crappy-go-elasticsearch-pkg"
)

func TestInit(t *testing.T) {
	cluster := elasticsearch.Init("http://localhost", 9200)

	if cluster.URI != "http://localhost" {
		t.Error("Expected: http://localhost, got ", cluster.URI)
	}

	if cluster.Port != 9200 {
		t.Error("Expected: 9200, got ", cluster.Port)
	}
}

func TestGeturi(t *testing.T) {
	cluster := elasticsearch.Init("http://localhost", 9200)

	uri := cluster.Geturi()
	if uri != "http://localhost" {
		t.Error("Expected: http://localhost, got ", uri)
	}
}

func TestGetport(t *testing.T) {
	cluster := elasticsearch.Init("http://localhost", 9200)

	port := cluster.Getport()
	if port != 9200 {
		t.Error("Expected: 9200, got ", port)
	}
}
