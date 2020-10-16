package redis_exporter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"

	re "github.com/oliver006/redis_exporter/lib/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/grafana/agent/pkg/integrations/config"
)

type Integration struct {
	c        Config
	exporter re.Exporter
}

// New creates a new redis_exporter integration.
func New(log log.Logger, c Config) (*Integration, error) {
	level.Debug(log).Log("msg", "initialising redis_exporer with config %w", c)

	exporter, err := re.NewRedisExporter(
		c.RedisAddr,
		c.GetExporterOptions(),
		re.BuildInfo{},
	)
	if err != nil {
		return nil, fmt.Error("failed to create redis exporter: %w", err)
	}

	return &Integration{
		c:        c,
		exporter: exporter,
	}, nil
}

// CommonConfig satisfies Integration.CommonConfig.
func (i *Integraiton) CommonConfig() config.Common { return i.c.CommonConfig }

// Name satisfies Integration.Name.
func (i *Integration) Name() string { return "redis_exporter" }

func (i *Integration) RegisterRoutes(r *mux.Router) error {
	handler, err := i.handler()
	if err != nil {
		return err
	}

	r.Handle(i.c.MetricPath, handler)
	return nil
}

func (i *Integration) handler() (http.Handler, error) {
	r := prometheus.NewRegistry()
	if err := r.Register(i.exporter); err != nil {
		return nil, fmt.Errorf("couldn't register redis_exporter: %w", err)
	}

	handler := promhttp.HandlerFor(
		r,
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
		},
	)

	// TODO handle instrumentation of metric handler

	return handler, nil
}

// ScrapeConfigs satisfies Integration.ScrapeConfigs.
func (i *Integration) ScrapeConfigs() []config.ScrapeConfig {
	return []config.ScrapeConfig{{
		JobName:     i.Name(),
		MetricsPath: "/metrics",
	}}
}

// Run satisfies Integration.Run.
func (i *Integration) Run(ctx context.Context) error {
	// We don't need to do anything here, so we can just wait for the context to
	// finish.
	<-ctx.Done()
	return ctx.Err()
}
