package redis_exporter

import (
	re "github.com/oliver006/redis_exporter/lib/exporter"
	
	"github.com/grafana/agent/pkg/integrations/config"
)

var (
	// DefaultConfig holds non-zero default options for the Config when it is
	// unmarshaled from YAML.
	DefaultConfig = Config{
		RedisAddr:         "redis://localhost:6379",
		Namespace:         "redis",
		ListenAddress:     ":9121",
		MetricPath:        "/metrics",
		LogFormat:         "txt",
		ConfigCommand:     "CONFIG",
		ConnectionTimeout: "15s",
		SetClientName:     true,
	}
)

// Config controls the redis_exporter integration
type Config struct {
	CommmonConfig config.Common `yaml:",inline"`
	Enabled       bool          `yaml:"enabled"`

	IncludeExporterMetrics bool   `yaml:"include_exporter_metrics"`
	RedisAddr              string `yaml:"redis_addr"`
	RedisUser              string `yaml:"redis_user"`
	RedisPwd               string `yaml:"redis_pwd"`
	Namespace              string `yaml:"namespace"`
	CheckKeys              string `yaml:"check_keys"`
	CheckSingleKeys        string `yaml:"check_single_keys"`
	CheckStreams           string `yaml:"check_streams"`
	CheckSingleStreams     string `yaml:"check_single_streams"`
	ScriptPath             string `yaml:"script_path"`
	ListenAddress          string `yaml:"listen_address"`
	MetricPath             string `yaml:"metric_path"`
	LogFormat              string `yaml:"log_format"`
	ConfigCommand          string `yaml:"config_command"`
	ConnectionTimeout      string `yaml:"connection_timeout"`
	TLSClientKeyFile       string `yaml:"tls_client_key_file"`
	TLSClientCertFile      string `yaml:"tls_client_cert_file"`
	TLSCaCertFile          string `yaml:"tls_ca_cert_file"`
	TLSServerKeyFile       string `yaml:"tls_server_key_file"`
	TLSServerCertFile      string `yaml:"tls_server_cert_file"`
	IsDebug                bool   `yaml:"is_debug"`
	SetClientName          bool   `yaml:"set_client_name"`
	IsTile38               bool   `yaml:"is_tile38"`
	ExportClientList       bool   `yaml:"export_client_list"`
	ShowVersion            bool   `yaml:"show_version"`
	RedisMetricsOnly       bool   `yaml:"redis_metrics_only"`
	PingOnConnect          bool   `yaml:"ping_on_connect"`
	InclSystemMetrics      bool   `yaml:"incl_system_metrics"`
	SkipTLSVerification    bool   `yaml:"skip_tls_verification"`
}

// GetExporterOptions returns relevant Config properties as a redis_exporter
// Options struct. The redis_exporter Options struct has no yaml tags, so
// we marshal the yaml into Config and then create the re.Options from that.
func (c Config) GetExporterOptions() re.Options {
	return re.Options{
		Namespace:          c.Namespace,
		ConfigCommandName:  c.ConfigCommand,
		SetClientName:      c.SetClientName,
		MetricsPath:        c.MetricPath,
		ConnectionTimeouts: c.ConnectionTimeout
	}
}
