package redis_exporter

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/prometheus/pkg/textparse"
	"github.com/stretchr/testify/require"
)

func TestRedisExporter(t *testing.T) {
	cfg = DefaultConfig

	logger := log.NewNopLogger()
	integration, err := New(logger, cfg)
	require.NoError(t, err, "failed to setup redis_exporter")

	r := mux.NewRouter()
	err = integration.RegisterRoutes(r)
	require.NoError(t, err)

	// Invoke /metrics and parse the response
	srv := httptest.NewServer(r)
	defer srv.Close()

	res, err := http.Get(srv.URL + cfg.MetricPath)
	require.NoError(t, err)

	body, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)

	p := textparse.NewPromParser(body)
	for {
		_, err := p.Next()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
	}
}

// TODO think if you've covered reasonable tests
