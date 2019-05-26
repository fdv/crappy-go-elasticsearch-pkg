package elasticsearch

import (
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cluster = Init("http://localhost", 9200)
var clustersettings = `{
	"persistent":{
		"cluster": {
			"routing": {
				"allocation": {
					"enable": "primary"
				}
			}
		}
	},
	"transient":{}
	}`

func TestInit(t *testing.T) {
	assert.Equal(t, cluster.URI, "http://localhost", "Cluster URI Should be http://localhost")
	assert.Equal(t, cluster.Port, uint16(9200), "Cluster port should be 9200")
	assert.Equal(t, cluster.URL, "http://localhost:9200", "Cluster URL should be http://lo alhost:9200")
}

func TestGeturi(t *testing.T) {
	assert.Equal(t, cluster.Geturi(), "http://localhost", "Cluster URI Should be http://localhost")
}

func TestGetport(t *testing.T) {
	assert.Equal(t, cluster.Getport(), uint16(9200), "Cluster port should be 9200")
}

func TestGeturl(t *testing.T) {
	cluster := Init("http://localhost", 9200)

	assert.Equal(t, cluster.Geturl(), "http://localhost:9200", "Cluster URL should be http://localhost:9200")

}

func TestGetsettings(t *testing.T) {
	resp := clustersettings
	expected := `{
	"persistent":{
		"cluster": {
			"routing": {
				"allocation": {
					"enable": "primary"
				}
			}
		}
	},
	"transient":{}
	}`

	ts := mockescluster(resp)
	defer ts.Close()

	result := cluster.Getsettings()
	assert.Equal(t, expected, result, "Error in cluster settings")
}

func mockescluster(resp string) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(resp))
	}
	l, err := net.Listen("tcp", "127.0.0.1:9200")
	if err != nil {
		log.Fatal(err)
	}

	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}))
	ts.Listener.Close()
	ts.Listener = l
	ts.Start()
	return ts
}
