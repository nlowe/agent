package redis_exporter

import (
	"fmt"
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

// TestRedisExporterDefaults runs a basic integration test for redis_exporter,
// doing the following:
//
// 1. Creating an integration with default exporter config
// 2. Scrape the integration
// 3. Parse the result to ensure there's at least one metric
//
// This ensures that the default integration config results in an exporter
// being created and the handler is set up properly. We do not check the
// contents of the scrape, just that it was parsable by Prometheus.
//
// Note that the scrape results in an error from the exporter because it
// cannot connect to a redis. This is irrelevant for the purposes of this
// test.
func TestRedisExporterDefaults(t *testing.T) {
	cfg := DefaultConfig

	logger := log.NewNopLogger()
	integration, err := New(logger, cfg)
	require.NoError(t, err, "failed to setup redis_exporter")

	r := mux.NewRouter()
	err = integration.RegisterRoutes(r)
	require.NoError(t, err)

	// Invoke /metrics and parse the response
	srv := httptest.NewServer(r)
	defer srv.Close()

	res, err := http.Get(srv.URL + "/metrics")
	require.NoError(t, err)

	body, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)

	fmt.Printf("%s", body)

	p := textparse.NewPromParser(body)
	for {
		_, err := p.Next()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
	}
}

// Test some config has an effect on the exporter.