package redis_exporter

import (
	"time"

	re "github.com/oliver006/redis_exporter/lib/exporter"

	"github.com/grafana/agent/pkg/integrations/config"
)

var (
	// DefaultConfig holds non-zero default options for the Config when it is
	// unmarshaled from YAML.
	DefaultConfig = Config{
		RedisAddr:         "redis://localhost:6379",
		Namespace:         "redis",
		ConfigCommand:     "CONFIG",
		ConnectionTimeout: (15 * time.Second),
		SetClientName:     true,
	}
)

// Config controls the redis_exporter integration. The exporter accepts more
// config properties than this, but these are the only fields with non-default
// values that we need to define right now.
type Config struct {
	CommonConfig config.Common `yaml:",inline"`
	Enabled      bool          `yaml:"enabled"`

	IncludeExporterMetrics bool          `yaml:"include_exporter_metrics"`
	RedisAddr              string        `yaml:"redis_addr"`
	Namespace              string        `yaml:"namespace"`
	ConfigCommand          string        `yaml:"config_command"`
	ConnectionTimeout      time.Duration `yaml:"connection_timeout"`
	SetClientName          bool          `yaml:"set_client_name"`
}

// GetExporterOptions returns relevant Config properties as a redis_exporter
// Options struct. The redis_exporter Options struct has no yaml tags, so
// we marshal the yaml into Config and then create the re.Options from that.
func (c Config) GetExporterOptions() re.Options {
	return re.Options{
		Namespace:          c.Namespace,
		ConfigCommandName:  c.ConfigCommand,
		SetClientName:      c.SetClientName,
		ConnectionTimeouts: c.ConnectionTimeout,
	}
}
